package getRequests

import (
	"encoding/json"
	"fmt"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
	"log"
	"net/http"
)

type Member struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	Description string `json:"description"`
	StartupID   int    `json:"startup_id"`
}
type Achievement struct {
	ID        int    `json:"id"`
	Name      string `json:"achievement"`
	StartupID int    `json:"startupID"`
}
type WholeStartup struct {
	ID                int           `json:"startup_id"`
	Name              string        `json:"name"`
	Login             string        `json:"login"`
	Password          string        `json:"password"`
	Email             string        `json:"email"`
	Description       string        `json:"description"`
	Logo              string        `json:"logo"`
	LowestInvestment  int           `json:"lowestInvestment"`
	HighestInvestment int           `json:"highestInvestment"`
	Region            string        `json:"region"`
	WebSite           string        `json:"website"`
	TeamSize          int           `json:"team_size"`
	Industry          string        `json:"industry"`
	Team              []Member      `json:"team"`
	Favourites        []Favourite   `json:"favourites"`
	Achievements      []Achievement `json:"achievements"`
}

func GetStartup(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Fprintf(w, "wrong method type")
		return
	}
	//other.AccessSetter(w)
	params := r.URL.Query()
	id := params.Get("id")
	query := server.DBConn.QueryRow("SELECT * FROM startups WHERE id=?", id)
	var startup WholeStartup
	err := query.Scan(&startup.ID, &startup.Name, &startup.Login, &startup.Password,
		&startup.Email, &startup.Description, &startup.Logo, &startup.LowestInvestment, &startup.HighestInvestment,
		&startup.Region, &startup.WebSite, &startup.TeamSize, &startup.Industry)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			fmt.Fprintf(w, "{}")
			return
		}
		log.Fatal(err)
		return
	}
	members, err := server.DBConn.Query("SELECT * FROM team WHERE startup_id=" + string(id))
	for members.Next() {
		var member Member
		err = members.Scan(&member.ID, &member.Name, &member.Role, &member.Description, &member.StartupID)
		if err != nil {
			log.Fatal(err)
			return
		}
		startup.Team = append(startup.Team, member)
	}
	members.Close()
	favourites, err := server.DBConn.Query("SELECT * FROM favourite_investors WHERE startup_id=" + string(id))
	if err != nil {
		log.Fatal(err)
		return
	}
	for favourites.Next() {
		var favourite Favourite
		err = favourites.Scan(&favourite.ID, &favourite.InvestorID, &favourite.StartupID)
		if err != nil {
			log.Fatal(err)
			return
		}
		startup.Favourites = append(startup.Favourites, favourite)
	}
	favourites.Close()
	achievements, err := server.DBConn.Query("SELECT * FROM achievements WHERE startup_id=" + string(id))
	defer achievements.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	for achievements.Next() {
		var achievement Achievement
		err = achievements.Scan(&achievement.ID, &achievement.Name, &achievement.StartupID)
		if err != nil {
			log.Fatal(err)
			return
		}
		startup.Achievements = append(startup.Achievements, achievement)
	}
	jsonBytes, err := json.Marshal(startup)
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}
	fmt.Fprintf(w, string(jsonBytes))
}

package getRequests

import (
	"encoding/json"
	"fmt"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
	"log"
	"net/http"
)

type Startup struct {
	Name              string `json:"name"`
	Login             string `json:"login"`
	Password          string `json:"password"`
	Email             string `json:"email"`
	Description       string `json:"description"`
	Logo              string `json:"logo"`
	LowestInvestment  int    `json:"lowestInvestment"`
	HighestInvestment int    `json:"highestInvestment"`
	Region            string `json:"region"`
	WebSite           string `json:"website"`
	TeamSize          string `json:"team_size"`
}

func GetStartups(w http.ResponseWriter, r http.Request) {
	if r.Method != "GET" {
		fmt.Fprintf(w, "wrong method type")
		return
	}
	params := r.URL.Query()
	region := params.Get("region")
	category := params.Get("category")
	lowestTeam := params.Get("lowestTeam")
	highestTeam := params.Get("highestTeam")
	lowestInvestment := params.Get("lowestInvestment")
	highestInvestment := params.Get("highestInvestment")
	request := "SELECT * FROM startups WHERE and"
	if region != "" {
		request += " region=" + region + " and"
	}
	if category != "" {
		request += " category=" + category + " and"
	}
	if lowestTeam != "" && highestTeam != "" {
		request += " team_size>=" + lowestTeam + " and team_size<=" + highestTeam + " and"
	}
	if lowestInvestment != "" && highestInvestment != "" {
		request += " lowest_investment>=" + lowestInvestment + " and highest_investment<=" + highestInvestment + " and"
	}
	var startups []Startup
	rows, err := server.DBConn.Query(request)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	for rows.Next() {
		var startup Startup
		err = rows.Scan(&startup.Name, &startup.Login, &startup.Password,
			&startup.Email, &startup.Description, &startup.Logo, &startup.LowestInvestment, &startup.HighestInvestment,
			&startup.Region, &startup.WebSite, &startup.TeamSize)
		if err != nil {
			log.Fatal(err)
			return
		}
		startups = append(startups, startup)
	}
	jsonBytes, err := json.Marshal(rows)
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}
	fmt.Fprintf(w, string(jsonBytes))
}

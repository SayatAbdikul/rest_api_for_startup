package getRequests

import (
	"encoding/json"
	"fmt"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
	"log"
	"net/http"
)

type Startup struct {
	ID                string `json:"startup_id"`
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
	TeamSize          int    `json:"team_size"`
}

func GetStartups(w http.ResponseWriter, r *http.Request) {
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
	ok := 0
	request := "SELECT * FROM startups WHERE"
	if region != "" {
		request += " region='" + region + "' and"
		ok = 1
	}
	if category != "" {
		request += " category='" + category + "' and"
		ok = 1
	}
	if lowestTeam != "" && highestTeam != "" {
		request += " team_size>=" + lowestTeam + " and team_size<=" + highestTeam + " and"
		ok = 1
	}
	if lowestInvestment != "" && highestInvestment != "" {
		request += " lowest_investment>=" + lowestInvestment + " and highest_investment<=" + highestInvestment + " and"
		ok = 1
	}
	if ok == 1 {
		newStr := request[:len(request)-3]
		request = newStr
	} else {
		newStr := request[:len(request)-6]
		request = newStr
	}
	fmt.Println(request)
	var startups []Startup
	rows, err := server.DBConn.Query(request)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	for rows.Next() {
		var startup Startup
		err = rows.Scan(&startup.ID, &startup.Name, &startup.Login, &startup.Password,
			&startup.Email, &startup.Description, &startup.Logo, &startup.LowestInvestment, &startup.HighestInvestment,
			&startup.Region, &startup.WebSite, &startup.TeamSize)
		if err != nil {
			log.Fatal(err)
			return
		}
		startups = append(startups, startup)
	}
	jsonBytes, err := json.Marshal(startups)
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}
	fmt.Fprintf(w, string(jsonBytes))
}

package patchRequests

import (
	"encoding/json"
	"fmt"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
	"log"
	"net/http"
)

type Startup struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Login             string `json:"login"`
	Password          string `json:"password"`
	Email             string `json:"email"`
	Logo              string `json:"logo"`
	LowestInvestment  int    `json:"lowestInvestment"`
	HighestInvestment int    `json:"highestInvestment"`
	Region            string `json:"region"`
	WebSite           string `json:"website"`
}

func PatchStartup(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PATCH" {
		fmt.Fprintf(w, "the method is not patch type")
		return
	}
	var startup Startup
	err := json.NewDecoder(r.Body).Decode(&startup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	query, err := server.DBConn.Prepare("UPDATE startups SET name=?, login=?, password=?, email=?, " +
		"logo=?, lowest_investment=?, highest_investment=?, region=?, website=? WHERE id=?")
	defer query.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = query.Exec(&startup.Name, &startup.Login, &startup.Password, &startup.Email,
		&startup.Logo, &startup.LowestInvestment, &startup.HighestInvestment, &startup.Region, &startup.WebSite, &startup.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Fprintf(w, "the data was successfully updated")
}

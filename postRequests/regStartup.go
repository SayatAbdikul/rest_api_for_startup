package postRequests

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/SayatAbdikul/rest_api_for_startup/other"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
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
	Industry          string `json:"industry"`
}

func RegStartup(w http.ResponseWriter, r *http.Request) {
	other.AccessSetter(w)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if r.Method != "POST" {
		fmt.Fprintf(w, "error: the request is not a POST type")
		return
	}
	var query Startup
	err := json.NewDecoder(r.Body).Decode(&query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	other.Connect()
	defer server.DBConn.Close()
	other.MuteStartup.Lock()
	defer other.MuteStartup.Unlock()
	stmt, err := server.DBConn.Prepare("INSERT INTO startups (name, login, password, email, " +
		"description, logo, lowest_investment, highest_investment, region, website, industry) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	defer stmt.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = stmt.Exec(query.Name, query.Login, query.Password, query.Email, query.Description, query.Logo,
		query.LowestInvestment, query.HighestInvestment, query.Region, query.WebSite, query.Industry)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Fprintf(w, "data entered successfully")
}

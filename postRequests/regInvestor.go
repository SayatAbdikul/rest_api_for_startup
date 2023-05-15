package postRequests

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/SayatAbdikul/rest_api_for_startup/other"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
)

type Investor struct {
	Name        string `json:"name"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
	Region      string `json:"region"`
	WebSite     string `json:"website"`
	Investment  int    `json:"investment"`
	Industry    string `json:"industry"`
}

func RegInvestor(w http.ResponseWriter, r *http.Request) {
	other.AccessSetter(w)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if r.Method != "POST" {
		fmt.Fprintf(w, "error: the request is not a POST type")
		return
	}
	var query Investor
	err := json.NewDecoder(r.Body).Decode(&query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	other.Connect()
	defer server.DBConn.Close()
	other.MuteInvestor.Lock()
	defer other.MuteInvestor.Unlock()
	stmt, err := server.DBConn.Prepare("INSERT INTO investors (name, login, password, email, " +
		"description, picture, region, website, investment, industry) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	defer stmt.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = stmt.Exec(query.Name, query.Login, query.Password, query.Email, query.Description,
		query.Picture, query.Region, query.WebSite, query.Investment, query.Industry)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Fprintf(w, "data entered successfully")
}

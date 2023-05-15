package patchRequests

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/SayatAbdikul/rest_api_for_startup/other"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
)

type Investor struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Picture    string `json:"logo"`
	Region     string `json:"region"`
	WebSite    string `json:"website"`
	Investment int    `json:"investment"`
	Industry   string `json:"industry"`
}

func PatchInvestor(w http.ResponseWriter, r *http.Request) {
	other.AccessSetter(w)
	if r.Method != "PATCH" {
		fmt.Fprintf(w, "the method is not patch type")
		return
	}
	var investor Investor
	err := json.NewDecoder(r.Body).Decode(&investor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	other.Connect()
	defer server.DBConn.Close()
	other.MuteInvestor.Lock()
	defer other.MuteInvestor.Unlock()
	query, err := server.DBConn.Prepare("UPDATE investors SET name=?, login=?, password=?, email=?, " +
		"picture=?, region=?, website=?, investment=?, industry=? WHERE id=?")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer query.Close()
	_, err = query.Exec(&investor.Name, &investor.Login, &investor.Password, &investor.Email,
		&investor.Picture, &investor.Region, &investor.WebSite, &investor.Investment, &investor.Industry, &investor.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Fprintf(w, "the data was successfully updated")
}

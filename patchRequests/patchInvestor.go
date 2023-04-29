package patchRequests

import (
	"encoding/json"
	"fmt"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
	"log"
	"net/http"
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
	if r.Method != "PATCH" {
		fmt.Fprintf(w, "the method is not patch type")
		return
	}
	//other.AccessSetter(w)
	var investor Investor
	err := json.NewDecoder(r.Body).Decode(&investor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	query, err := server.DBConn.Prepare("UPDATE investors SET name=?, login=?, password=?, email=?, " +
		"picture=?, region=?, website=?, investment=?, industry=? WHERE id=?")
	defer query.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = query.Exec(&investor.Name, &investor.Login, &investor.Password, &investor.Email,
		&investor.Picture, &investor.Region, &investor.WebSite, &investor.Investment, &investor.Industry, &investor.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Fprintf(w, "the data was successfully updated")
}

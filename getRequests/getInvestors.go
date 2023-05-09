package getRequests

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/SayatAbdikul/rest_api_for_startup/other"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
)

type Investor struct {
	ID          string `json:"investor_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
	Region      string `json:"region"`
	WebSite     string `json:"website"`
	Investment  int    `json:"investment"`
	Industry    string `json:"industry"`
}

func GetInvestors(w http.ResponseWriter, r *http.Request) {
	other.AccessSetter(w)
	if r.Method != "GET" {
		fmt.Fprintf(w, "wrong method type")
		return
	}
	//other.AccessSetter(w)
	params := r.URL.Query()
	region := params.Get("region")
	lowestInvestment := params.Get("lowestInvestment")
	highestInvestment := params.Get("highestInvestment")
	sort := params.Get("sort")
	ok := 0
	request := "SELECT id, name, email, description, picture, region, website, investment, industry FROM investors WHERE"
	if region != "" {
		request += " region='" + region + "' and"
		ok = 1
	}
	if lowestInvestment != "" && highestInvestment != "" {
		request += " investment>=" + lowestInvestment + " and investment<=" + highestInvestment + " and"
		ok = 1
	}
	if ok == 1 {
		newStr := request[:len(request)-3]
		request = newStr
	} else {
		newStr := request[:len(request)-6]
		request = newStr
	}
	if sort == "ascending" {
		request += " ORDER BY name ASC"
	}
	if sort == "descending" {
		request += " ORDER BY name DESC"
	}
	fmt.Println(request)
	var investors []Investor
	rows, err := server.DBConn.Query(request)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	for rows.Next() {
		var investor Investor
		err = rows.Scan(&investor.ID, &investor.Name,
			&investor.Email, &investor.Description, &investor.Picture,
			&investor.Region, &investor.WebSite, &investor.Investment, &investor.Industry)
		if err != nil {
			log.Fatal(err)
			return
		}
		investors = append(investors, investor)
	}
	jsonBytes, err := json.Marshal(investors)
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}
	fmt.Fprintf(w, string(jsonBytes))
}

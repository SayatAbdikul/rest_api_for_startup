package getRequests

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/SayatAbdikul/rest_api_for_startup/other"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
)

type Favourite struct {
	ID         int `json:"id"`
	StartupID  int `json:"startup_id"`
	InvestorID int `json:"investorID"`
}
type Case struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Investment  int    `json:"investment"`
	InvestorID  int    `json:"investorID"`
}
type WholeInvestor struct {
	ID          int         `json:"investor_id"`
	Name        string      `json:"name"`
	Login       string      `json:"login"`
	Password    string      `json:"password"`
	Email       string      `json:"email"`
	Picture     string      `json:"picture"`
	Region      string      `json:"region"`
	Description string      `json:"description"`
	WebSite     string      `json:"website"`
	Investment  int         `json:"investment"`
	Industry    string      `json:"industry"`
	Favourites  []Favourite `json:"favourites"`
	Cases       []Case      `json:"cases"`
}

func GetInvestor(w http.ResponseWriter, r *http.Request) {
	other.AccessSetter(w)
	if r.Method != "GET" {
		fmt.Fprintf(w, "wrong method type")
		return
	}
	other.Connect()
	defer server.DBConn.Close()
	params := r.URL.Query()
	id := params.Get("id")
	query := server.DBConn.QueryRow("SELECT * FROM investors WHERE id=?", id)
	var investor WholeInvestor
	err := query.Scan(&investor.ID, &investor.Name, &investor.Login, &investor.Password,
		&investor.Email, &investor.Picture, &investor.Region,
		&investor.Description, &investor.WebSite, &investor.Investment, &investor.Industry)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			fmt.Fprintf(w, "{}")
			return
		}
		log.Fatal(err)
		return
	}
	favourites, err := server.DBConn.Query("SELECT * FROM favourite_startups WHERE investor_id=" + string(id))
	if err != nil {
		log.Fatal(err)
		return
	}
	for favourites.Next() {
		var favourite Favourite
		err = favourites.Scan(&favourite.ID, &favourite.StartupID, &favourite.InvestorID)
		if err != nil {
			log.Fatal(err)
			return
		}
		investor.Favourites = append(investor.Favourites, favourite)
	}
	favourites.Close()
	cases, err := server.DBConn.Query("SELECT * FROM cases WHERE investor_id=" + string(id))
	defer cases.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	for cases.Next() {
		var data Case
		err = cases.Scan(&data.ID, &data.Title, &data.Description, &data.Investment, &data.InvestorID)
		if err != nil {
			log.Fatal(err)
			return
		}
		investor.Cases = append(investor.Cases, data)
	}
	jsonBytes, err := json.Marshal(investor)
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}
	fmt.Fprintf(w, string(jsonBytes))
}

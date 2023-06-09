package postRequests

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/SayatAbdikul/rest_api_for_startup/other"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
)

type FavStartup struct {
	ID         int `json:"id"`
	InvestorID int `json:"investorID"`
}

func RegFavouriteStartup(w http.ResponseWriter, r *http.Request) {
	other.AccessSetter(w)
	if r.Method != "POST" {
		fmt.Fprintf(w, "the method of the request is not a post type")
		return
	}
	//other.AccessSetter(w)
	var data FavStartup
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	other.Connect()
	defer server.DBConn.Close()
	other.MuteFavStartup.Lock()
	defer other.MuteFavStartup.Unlock()
	query, err := server.DBConn.Prepare("INSERT INTO favourite_startups (startup_id, investor_id) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = query.Exec(data.ID, data.InvestorID)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Fprintf(w, "all data was entered successfully")
}

package postRequests

import (
	"encoding/json"
	"fmt"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
	"log"
	"net/http"
)

type FavInvestor struct {
	ID        int `json:"id"`
	StartupID int `json:"startupID"`
}

func RegFavouriteInvestor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "the method of the request is not a post type")
		return
	}
	var data FavInvestor
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	query, err := server.DBConn.Prepare("INSERT INTO favourite_investors (investor_id, startup_id) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = query.Exec(data.ID, data.StartupID)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Fprintf(w, "all data was entered successfully")
}

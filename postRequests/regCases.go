package postRequests

import (
	"encoding/json"
	"fmt"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
	"log"
	"net/http"
)

type Case struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Investment  int    `json:"investment"`
	InvestorID  int    `json:"investor_id"`
}

func RegCases(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	if r.Method != "POST" {
		fmt.Fprintf(w, "error: the request is not a POST type")
		return
	}
	var data []Case
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for _, val := range data {
		stmt, err := server.DBConn.Prepare("INSERT INTO cases (title, description, investment, investor_id) VALUES (?, ?, ?, ?)")
		if err != nil {
			log.Fatal(err)
			return
		}
		_, err = stmt.Exec(val.Title, val.Description, val.Investment, val.InvestorID)
		if err != nil {
			log.Fatal(err)
			return
		}
		stmt.Close()
	}
	fmt.Fprintf(w, "all records were saved")
}

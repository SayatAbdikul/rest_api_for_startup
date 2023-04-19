package postRequests

import (
	"encoding/json"
	"fmt"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
	"log"
	"net/http"
)

type Member struct {
	Name        string `json:"name"`
	Role        string `json:"role"`
	Description string `json:"description"`
	StartupID   int    `json:"startup_id"`
}

func RegTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	if r.Method != "POST" {
		fmt.Fprintf(w, "error: the request is not a POST type")
		return
	}
	var team []Member
	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for _, val := range team {
		stmt, err := server.DBConn.Prepare("INSERT INTO team (name, role, description, startup_id) VALUES (?, ?, ?, ?)")
		if err != nil {
			log.Fatal(err)
			return
		}
		_, err = stmt.Exec(val.Name, val.Role, val.Description, val.StartupID)
		if err != nil {
			log.Fatal(err)
			return
		}
		stmt.Close()
	}
	fmt.Fprintf(w, "all records were saved")
}

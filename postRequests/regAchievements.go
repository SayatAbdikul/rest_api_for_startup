package postRequests

import (
	"encoding/json"
	"fmt"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
	"log"
	"net/http"
)

type Achievement struct {
	Name      string `json:"achievement"`
	StartupID int    `json:"startupID"`
}

func RegAchievements(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	if r.Method != "POST" {
		fmt.Fprintf(w, "error: the request is not a POST type")
		return
	}
	var achievements []Achievement
	err := json.NewDecoder(r.Body).Decode(&achievements)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for _, val := range achievements {
		stmt, err := server.DBConn.Prepare("INSERT INTO achievements (achievement, startup_id) VALUES (?, ?)")
		if err != nil {
			log.Fatal(err)
			return
		}
		_, err = stmt.Exec(val.Name, val.StartupID)
		if err != nil {
			log.Fatal(err)
			return
		}
		stmt.Close()
	}
	fmt.Fprintf(w, "all records were saved")
}

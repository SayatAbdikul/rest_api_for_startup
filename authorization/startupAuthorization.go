package authorization

import (
	"encoding/json"
	"fmt"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
	"net/http"
)

type AuthStartup struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
type Startup struct {
	ID          int    `json:"id"`
	Login       string `json:"login"`
	ErrorStatus bool   `json:"errorStatus"`
}

func StartupAuthorization(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "the request type is not a post")
		return
	}
	var authData AuthStartup
	err := json.NewDecoder(r.Body).Decode(&authData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	var data Startup
	query, err := server.DBConn.Prepare("SELECT id, login FROM startups WHERE login=? and password=?")
	rows, err := query.Query(authData.Login, authData.Password)
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&data.ID, &data.Login)
	}
	if data.Login == "" {
		data.ErrorStatus = true
	}
	jsonBytes, err := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

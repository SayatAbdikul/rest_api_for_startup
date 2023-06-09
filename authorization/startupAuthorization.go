package authorization

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SayatAbdikul/rest_api_for_startup/other"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
)

func StartupAuthorization(w http.ResponseWriter, r *http.Request) {
	other.AccessSetter(w)
	if r.Method != "POST" {
		fmt.Fprintf(w, "the request type is not a post")
		return
	}
	//other.AccessSetter(w)
	var authData AuthData
	err := json.NewDecoder(r.Body).Decode(&authData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	var data OutData
	other.Connect()
	defer server.DBConn.Close()
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

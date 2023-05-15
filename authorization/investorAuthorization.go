package authorization

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SayatAbdikul/rest_api_for_startup/other"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
)

type AuthData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
type OutData struct {
	ID          int    `json:"id"`
	Login       string `json:"login"`
	ErrorStatus bool   `json:"errorStatus"`
}

func InvestorAuthorization(w http.ResponseWriter, r *http.Request) {
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
	query, _ := server.DBConn.Prepare("SELECT id, login FROM investors WHERE login=? and password=?")
	rows, _ := query.Query(authData.Login, authData.Password)
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&data.ID, &data.Login)
	}
	if data.Login == "" {
		data.ErrorStatus = true
	}
	jsonBytes, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

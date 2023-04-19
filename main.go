package main

import (
	"database/sql"
	post "github.com/SayatAbdikul/rest_api_for_startup/postRequests"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

var DB = sql.DB{}

func main() {
	var err error
	server.DBConn, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/infomatrix_project")
	defer server.DBConn.Close()
	if err != nil {
		panic(err.Error())
	}
	http.HandleFunc("/regStartup", post.RegStartup)
	http.HandleFunc("/regTeam", post.RegTeam)
	http.ListenAndServe(":9090", nil)
}

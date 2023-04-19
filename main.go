package main

import (
	"database/sql"
	post "github.com/SayatAbdikul/rest_api_for_startup/postRequests"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

var DB = sql.DB{}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1)/infomatrix_project")
	if err != nil {
		panic(err.Error())
	}
	http.HandleFunc("/regStartup", post.RegStartup())
	db.Close()
	http.ListenAndServe(":9090", nil)
}

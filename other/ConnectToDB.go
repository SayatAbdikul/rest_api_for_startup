package other

import (
	"database/sql"
	"log"

	"github.com/SayatAbdikul/rest_api_for_startup/server"
)

func Connect() {
	var err error
	server.DBConn, err = sql.Open("mysql", "root:root@tcp(localhost:8889)/infomatrix_project")
	if err != nil {
		log.Fatal(err)
	}
}

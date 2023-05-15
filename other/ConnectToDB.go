package other

import (
	"database/sql"
	"log"

	"github.com/SayatAbdikul/rest_api_for_startup/server"
)

func Connect() {
	var err error
	server.DBConn, err = sql.Open("mysql", "admin_root:BMKX55Rnt3MECAHB@tcp(31.172.67.121:3306)/infomatrix_project")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
}

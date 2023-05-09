package other

import (
	"database/sql"

	"github.com/SayatAbdikul/rest_api_for_startup/server"
)

func Connect() error {
	var err error
	server.DBConn, err = sql.Open("mysql", "admin_root:BMKX55Rnt3MECAHB@tcp(31.172.67.121:3306)/infomatrix_project")
	return err
}

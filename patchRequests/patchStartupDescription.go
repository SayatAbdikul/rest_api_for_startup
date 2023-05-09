package patchRequests

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/SayatAbdikul/rest_api_for_startup/other"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
)

type StartupDescription struct {
	StartupID   int    `json:"startup_id"`
	Description string `json:"description"`
}

func PatchStartupDescription(w http.ResponseWriter, r *http.Request) {
	other.AccessSetter(w)
	if r.Method != "PATCH" {
		fmt.Fprintf(w, "the method is not patch type")
		return
	}
	//other.AccessSetter(w)
	var newContent StartupDescription
	err := json.NewDecoder(r.Body).Decode(&newContent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	query, err := server.DBConn.Prepare("UPDATE startups SET description=? WHERE id=?")
	defer query.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = query.Exec(newContent.Description, newContent.StartupID)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Fprintf(w, "the patch request was completed successfully")
}

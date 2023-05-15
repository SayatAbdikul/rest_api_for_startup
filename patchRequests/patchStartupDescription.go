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
	var newContent StartupDescription
	err := json.NewDecoder(r.Body).Decode(&newContent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	other.Connect()
	defer server.DBConn.Close()
	other.MuteStartup.Lock()
	defer other.MuteStartup.Unlock()
	query, err := server.DBConn.Prepare("UPDATE startups SET description=? WHERE id=?")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer query.Close()
	_, err = query.Exec(newContent.Description, newContent.StartupID)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Fprintf(w, "the patch request was completed successfully")
}

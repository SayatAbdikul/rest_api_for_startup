package patchRequests

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/SayatAbdikul/rest_api_for_startup/other"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
)

type Achievement struct {
	ID          int    `json:"id"`
	Achievement string `json:"achievement"`
}

func PatchStartupAchievement(w http.ResponseWriter, r *http.Request) {
	other.AccessSetter(w)
	if r.Method != "PATCH" {
		fmt.Fprintf(w, "the method is not patch type")
		return
	}
	var newContent Achievement
	err := json.NewDecoder(r.Body).Decode(&newContent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	other.Connect()
	defer server.DBConn.Close()
	other.MuteAchievement.Lock()
	defer other.MuteAchievement.Unlock()
	query, err := server.DBConn.Prepare("UPDATE achievements SET achievement=? WHERE id=?")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer query.Close()
	_, err = query.Exec(newContent.Achievement, newContent.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Fprintf(w, "the patch request was completed successfully")
}

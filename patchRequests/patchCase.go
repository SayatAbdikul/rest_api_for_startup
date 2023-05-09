package patchRequests

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/SayatAbdikul/rest_api_for_startup/other"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
)

type Case struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Investment  int    `json:"investment"`
}

func PatchCase(w http.ResponseWriter, r *http.Request) {
	other.AccessSetter(w)
	if r.Method != "PATCH" {
		fmt.Fprintf(w, "the method is not patch type")
		return
	}
	//other.AccessSetter(w)
	var newContent Case
	err := json.NewDecoder(r.Body).Decode(&newContent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	query, err := server.DBConn.Prepare("UPDATE cases SET title=?, description=?, investment=? WHERE id=?")
	defer query.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = query.Exec(newContent.Title, newContent.Description, newContent.Investment, newContent.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Fprintf(w, "the patch request was completed successfully")
}

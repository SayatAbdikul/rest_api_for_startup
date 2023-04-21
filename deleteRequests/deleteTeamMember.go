package deleteRequests

import (
	"encoding/json"
	"fmt"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
	"log"
	"net/http"
)

func DeleteTeamMember(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		fmt.Fprintf(w, "the method of request is not delete")
	}
	var element Element
	err := json.NewDecoder(r.Body).Decode(&element)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = server.DBConn.Exec("DELETE FROM team WHERE id=?", element.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Fprintf(w, "the delete request completed successfully")
}

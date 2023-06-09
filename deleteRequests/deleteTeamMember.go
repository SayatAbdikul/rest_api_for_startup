package deleteRequests

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/SayatAbdikul/rest_api_for_startup/other"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
)

func DeleteTeamMember(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		fmt.Fprintf(w, "the method of request is not delete")
	}
	//other.AccessSetter(w)
	var element Element
	err := json.NewDecoder(r.Body).Decode(&element)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	other.MuteMember.Lock()
	defer other.MuteMember.Unlock()
	other.Connect()
	defer server.DBConn.Close()
	_, err = server.DBConn.Exec("DELETE FROM team WHERE id=?", element.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Fprintf(w, "the delete request completed successfully")
}

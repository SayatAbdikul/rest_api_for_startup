package deleteRequests

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/SayatAbdikul/rest_api_for_startup/other"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
)

func DeleteFavStartup(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		fmt.Fprintf(w, "the request type is not a delete")
		return
	}
	//other.AccessSetter(w)
	var element Element
	err := json.NewDecoder(r.Body).Decode(&element)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	other.MuteFavStartup.Lock()
	defer other.MuteFavStartup.Unlock()
	other.Connect()
	defer server.DBConn.Close()
	_, err = server.DBConn.Exec("DELETE FROM favourite_startups WHERE id=?", element.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Fprintf(w, "the delete request completed successfully")
}

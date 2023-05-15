package deleteRequests

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/SayatAbdikul/rest_api_for_startup/other"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
)

func DeleteInvestor(w http.ResponseWriter, r *http.Request) {
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
	other.MuteInvestor.Lock()
	defer other.MuteInvestor.Unlock()
	other.Connect()
	defer server.DBConn.Close()
	_, err = server.DBConn.Exec("DELETE FROM investors WHERE id=?", element.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	other.MuteFavStartup.Lock()
	defer other.MuteFavStartup.Unlock()
	_, err = server.DBConn.Exec("DELETE FROM favourite_startups WHERE investor_id=?", element.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	other.MuteCase.Lock()
	defer other.MuteCase.Unlock()
	_, err = server.DBConn.Exec("DELETE FROM cases WHERE investor_id=?", element.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Fprintf(w, "the delete request completed successfully")
}

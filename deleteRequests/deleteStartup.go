package deleteRequests

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/SayatAbdikul/rest_api_for_startup/other"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
)

func DeleteStartup(w http.ResponseWriter, r *http.Request) {
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
	other.MuteStartup.Lock()
	defer other.MuteStartup.Unlock()
	other.Connect()
	defer server.DBConn.Close()
	_, err = server.DBConn.Exec("DELETE FROM startups WHERE id=?", element.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	other.MuteFavInvestor.Lock()
	defer other.MuteFavInvestor.Unlock()
	_, err = server.DBConn.Exec("DELETE FROM favourite_investors WHERE startup_id=?", element.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	other.MuteMember.Lock()
	defer other.MuteMember.Unlock()
	_, err = server.DBConn.Exec("DELETE FROM team WHERE startup_id=?", element.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	other.MuteAchievement.Lock()
	defer other.MuteAchievement.Unlock()
	_, err = server.DBConn.Exec("DELETE FROM achievements WHERE startup_id=?", element.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Fprintf(w, "the delete request completed successfully")
}

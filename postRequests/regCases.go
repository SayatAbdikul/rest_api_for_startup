package postRequests

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/SayatAbdikul/rest_api_for_startup/other"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
)

type Case struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Investment  int    `json:"investment"`
	InvestorID  int    `json:"investor_id"`
}

func RegCases(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	other.AccessSetter(w)
	if r.Method != "POST" {
		fmt.Fprintf(w, "error: the request is not a POST type")
		return
	}
	var data []Case
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for _, val := range data {
		wg.Add(1)
		go func(val Case) {
			other.MuteCase.Lock()
			defer other.MuteCase.Unlock()
			wg.Done()
			other.Connect()
			defer server.DBConn.Close()
			stmt, err := server.DBConn.Prepare("INSERT INTO cases (title, description, investment, investor_id) VALUES (?, ?, ?, ?)")
			if err != nil {
				log.Fatal(err)
				return
			}
			_, err = stmt.Exec(val.Title, val.Description, val.Investment, val.InvestorID)
			if err != nil {
				log.Fatal(err)
				return
			}
			stmt.Close()
		}(val)
	}
	wg.Wait()
	fmt.Fprintf(w, "all records were saved")
}

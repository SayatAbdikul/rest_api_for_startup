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

type Member struct {
	Name        string `json:"name"`
	Role        string `json:"role"`
	Description string `json:"description"`
	StartupID   int    `json:"startup_id"`
}

func RegTeam(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	other.AccessSetter(w)
	if r.Method != "POST" {
		fmt.Fprintf(w, "error: the request is not a POST type")
		return
	}
	var team []Member
	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if len(team) > 0 {
		stmt, err := server.DBConn.Prepare("UPDATE startups SET team_size=team_size+? WHERE id=?")
		if err != nil {
			log.Fatal(err)
			return
		}
		_, err = stmt.Exec(len(team), team[0].StartupID)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	for _, val := range team {
		wg.Add(1)
		go func(val Member) {
			other.MuteMember.Lock()
			defer other.MuteMember.Unlock()
			wg.Done()
			other.Connect()
			defer server.DBConn.Close()
			stmt, err := server.DBConn.Prepare("INSERT INTO team (name, role, description, startup_id) VALUES (?, ?, ?, ?)")
			if err != nil {
				log.Fatal(err)
				return
			}
			_, err = stmt.Exec(val.Name, val.Role, val.Description, val.StartupID)
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

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

type Achievement struct {
	Name      string `json:"achievement"`
	StartupID int    `json:"startupID"`
}

func RegAchievements(w http.ResponseWriter, r *http.Request) {
	other.AccessSetter(w)
	var wg sync.WaitGroup
	myMutex := &other.MuteAchievement
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error: the request is not a POST type")
		return
	}

	var achievements []Achievement
	err := json.NewDecoder(r.Body).Decode(&achievements)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, val := range achievements {
		go func(val Achievement) {
			myMutex.Lock()
			defer myMutex.Unlock()
			defer wg.Done()
			other.Connect()
			defer server.DBConn.Close()
			stmt, err := server.DBConn.Prepare("INSERT INTO achievements (achievement, startup_id) VALUES (?, ?)")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Fatal(err)
				return
			}
			_, err = stmt.Exec(val.Name, val.StartupID)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Fatal(err)
				return
			}
			stmt.Close()
		}(val)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Fprintf(w, "all records were saved")
}

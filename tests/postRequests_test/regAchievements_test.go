package postRequests_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SayatAbdikul/rest_api_for_startup/other"
	"github.com/SayatAbdikul/rest_api_for_startup/postRequests"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
)

type Achievement struct {
	Name      string `json:"achievement"`
	StartupID int    `json:"startupID"`
}

func TestRegAchievements(t *testing.T) {
	other.Connect()
	defer server.DBConn.Close()
	var load [1]Achievement
	load[0] = Achievement{"win in hackathon", 33}
	loadBytes, _ := json.Marshal(load)
	request, err := http.NewRequest("POST", "/reg_achievements", bytes.NewBuffer(loadBytes))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(postRequests.RegAchievements)
	handler.ServeHTTP(rr, request)
	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		return
	}

	expected := `all records were saved`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
		return
	}
}

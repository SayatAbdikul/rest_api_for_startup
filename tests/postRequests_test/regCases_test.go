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

type Case struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Investment  int    `json:"investment"`
	InvestorID  int    `json:"investor_id"`
}

func TestRegCases(t *testing.T) {
	other.Connect()
	defer server.DBConn.Close()
	var load [1]Case
	load[0] = Case{"My Startup Case 1", "This is a description of my startup case.", 1000, 23}
	loadBytes, _ := json.Marshal(load)
	request, err := http.NewRequest("POST", "/reg_cases", bytes.NewBuffer(loadBytes))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(postRequests.RegCases)
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

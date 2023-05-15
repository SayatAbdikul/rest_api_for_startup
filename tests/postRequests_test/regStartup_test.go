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
	_ "github.com/go-sql-driver/mysql"
)

type Startup struct {
	Name              string `json:"name"`
	Login             string `json:"login"`
	Password          string `json:"password"`
	Email             string `json:"email"`
	Description       string `json:"description"`
	Logo              string `json:"logo"`
	LowestInvestment  int    `json:"lowestInvestment"`
	HighestInvestment int    `json:"highestInvestment"`
	Region            string `json:"region"`
	WebSite           string `json:"website"`
	Industry          string `json:"industry"`
}

func TestRegStartup(t *testing.T) {
	other.Connect()
	defer server.DBConn.Close()
	load := Startup{"название на кириллице", "test_startup", "test_password", "test@example.com", "This is a test startup.",
		"https://example.com/test_startup_logo.jpg", 1000, 10000, "Kazakhstan", "https://teststartup.com", "IT"}
	loadBytes, _ := json.Marshal(load)
	request, err := http.NewRequest("POST", "/reg_startup", bytes.NewBuffer(loadBytes))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(postRequests.RegStartup)
	handler.ServeHTTP(rr, request)
	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		return
	}

	expected := `data entered successfully`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
		return
	}
}

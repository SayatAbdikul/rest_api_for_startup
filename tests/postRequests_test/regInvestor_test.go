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

type Investor struct {
	Name        string `json:"name"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
	Region      string `json:"region"`
	WebSite     string `json:"website"`
	Investment  int    `json:"investment"`
	Industry    string `json:"industry"`
}

func TestRegInvestor(t *testing.T) {
	other.Connect()
	defer server.DBConn.Close()
	load := Investor{"Петр Петров", "test_investor", "test_password", "test@example.com", "This is a test инвестор.",
		"https://example.com/test_picture.jpg", "Kazakhstan", "https://testinvestor.com", 1000, "IT"}
	loadBytes, _ := json.Marshal(load)
	request, err := http.NewRequest("POST", "/reg_investor", bytes.NewBuffer(loadBytes))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(postRequests.RegInvestor)
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

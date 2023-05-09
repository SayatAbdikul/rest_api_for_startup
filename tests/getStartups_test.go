package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SayatAbdikul/rest_api_for_startup/getRequests"
	"github.com/SayatAbdikul/rest_api_for_startup/other"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
	_ "github.com/go-sql-driver/mysql"
)

type GetStartup struct {
	ID                string `json:"startup_id"`
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
	TeamSize          int    `json:"team_size"`
	Industry          string `json:"industry"`
}

func TestGetStartups(t *testing.T) {
	err := other.Connect()
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	defer server.DBConn.Close()
	request, err := http.NewRequest("GET", "/get_startups?region=&category=&lowestTeam=&highestTeam=&lowestInvestment=&highestInvestment=&sort=", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getRequests.GetStartups)
	handler.ServeHTTP(rr, request)
	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		return
	}
}

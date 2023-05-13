package getRequests_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SayatAbdikul/rest_api_for_startup/getRequests"
	"github.com/SayatAbdikul/rest_api_for_startup/other"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
	_ "github.com/go-sql-driver/mysql"
)

func TestGetInvestors(t *testing.T) {
	err := other.Connect()
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	defer server.DBConn.Close()
	req, err := http.NewRequest("GET", "/api/get_investors", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getRequests.GetInvestors)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

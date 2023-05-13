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

func TestGetInvestor(t *testing.T) {
	err := other.Connect()
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	defer server.DBConn.Close()
	req, err := http.NewRequest("GET", "/get_investor?id=4", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getRequests.GetInvestor)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	body := rr.Body.String()
	expected := `{"investor_id":4,"name":"иван иванов","login":"johndoe","password":"password123","email":"johndoe@example.com","picture":"https://example.com/logo.png","region":"US","description":"I am an old investor.","website":"https://example.com","investment":20000,"industry":"EdTech","favourites":null,"cases":[{"id":2,"title":"something","description":"something new","investment":10000,"investorID":4},{"id":3,"title":"something","description":"something","investment":10000,"investorID":4}]}`
	if body != expected {
		t.Errorf("handler returned unexpected body: got %v, wanted %v",
			body, expected)
	}
}

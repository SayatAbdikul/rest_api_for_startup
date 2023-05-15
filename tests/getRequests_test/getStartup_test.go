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

func TestGetStartup(t *testing.T) {
	other.Connect()
	defer server.DBConn.Close()
	req, err := http.NewRequest("GET", "/get_startup?id=9", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getRequests.GetStartup)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	body := rr.Body.String()
	expected := `{"startup_id":9,"name":"название на кириллице","login":"test_startup","password":"test_password","email":"test@example.com","description":"This is a test startup.","logo":"https://example.com/test_startup_logo.jpg","lowestInvestment":1000,"highestInvestment":10000,"region":"Kazakhstan","website":"https://teststartup.com","team_size":2,"industry":"IT","team":[{"id":6,"name":"Sayat","role":"programmer","description":"someone","startup_id":9},{"id":7,"name":"Sayaat","role":"programmer","description":"someone","startup_id":9}],"favourites":null,"achievements":[{"id":7,"achievement":"Win in Infomatrix","startupID":9}]}`
	if body != expected {
		t.Errorf("handler returned unexpected body: got %v, wanted %v",
			body, expected)
	}
}

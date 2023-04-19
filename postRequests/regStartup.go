package postRequests

import (
	"fmt"
	"net/http"
)

func RegStartup(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "error: the request is not a POST type")
		return
	}

}

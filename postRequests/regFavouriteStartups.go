package postRequests

import (
	"fmt"
	"net/http"
)

func RegFavouriteStartup(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "the method of the request is not a post type")
		return
	}
	
}

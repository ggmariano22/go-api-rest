package handlers

import (
	"net/http"
	"fmt"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page with Go and Docker")
}

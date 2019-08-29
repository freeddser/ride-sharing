package handlers

import (
	"fmt"
	"net/http"
)

func GetHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1 align='center'>Welcome to ride sharing system</h1>")
}

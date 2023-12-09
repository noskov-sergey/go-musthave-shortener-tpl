package server

import (
	"net/http"
)

func RunServer() {
	http.HandleFunc("/", RouteRedirect)
	http.ListenAndServe(":8080", nil)
}

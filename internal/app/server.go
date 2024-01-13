package server

import (
	"net/http"
)

func RunServer() {
	http.ListenAndServe(":8080", LinkRouter())
}

package server

import (
	"fmt"
	"go-musthave-shortener-tpl/internal/app/config"
	"net/http"
)

func RunServer(params *config.NetAddress) error {
	fmt.Println("Running server on")
	return http.ListenAndServe(params.String(), LinkRouter())
}

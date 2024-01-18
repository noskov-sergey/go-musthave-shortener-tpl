package main

import (
	"go-musthave-shortener-tpl/internal/app"
)

func main() {
	parseFlags(params)
	if err := server.RunServer(params); err != nil {
		panic(err)
	}
}

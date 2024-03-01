package main

import (
	serv "go-musthave-shortener-tpl/internal/app"
	storage "go-musthave-shortener-tpl/internal/app/backup"
	"go-musthave-shortener-tpl/internal/app/config"
	"log"
)

func main() {
	parseFlags(params, config.Fileparams)
	Consumer, err := storage.NewReader(config.Fileparams.String())
	if err != nil {
		log.Fatal(err)
	}
	defer Consumer.Close()
	err = Consumer.ReadFile()
	if err != nil {
		log.Fatal(err, "error from main with ReadFile")
	}

	if errServ := serv.RunServer(params); errServ != nil {
		panic(errServ)
	}
}

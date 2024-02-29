package main

import (
	"fmt"
	"go-musthave-shortener-tpl/internal/app"
	"go-musthave-shortener-tpl/internal/app/backup"
	"log"
)

func main() {
	parseFlags(params, fileparams)

	Consumer, err := backup.NewConsumer(fileparams.String())
	if err != nil {
		log.Fatal(err)
	}
	defer Consumer.Close()
	readEvent, err := Consumer.ReadEvent()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(readEvent)

	if err1 := server.RunServer(params); err != nil {
		panic(err1)
	}
}

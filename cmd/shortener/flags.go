package main

import (
	"flag"
	"go-musthave-shortener-tpl/internal/app/config"
)

var params = config.NewNetAddress()

func parseFlags(p *config.NetAddress) {
	_ = flag.Value(p)
	flag.Var(p, "a", "address and port to run server")
	flag.StringVar(&config.BaseUrla, "b", "http://localhost:8000/", "base url for short link server")
	flag.Parse()
}

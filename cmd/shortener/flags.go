package main

import (
	"flag"
	"go-musthave-shortener-tpl/internal/app/config"
	"regexp"
)

var params = config.NewNetAddress()

func parseFlags(p *config.NetAddress) {
	_ = flag.Value(p)
	flag.Var(p, "a", "address and port to run server")
	flag.Func("b", "base url for short link server", func(flagValue string) error {
		// разбиваем значение флага на слайс строк через запятую
		// и заливаем в переменную
		re := regexp.MustCompile(`([a-z]*)://([a-z]*):([0-9]*)`)
		config.BaseUrl = re.FindString(flagValue) + "/"
		return nil
	})
	flag.Parse()
}

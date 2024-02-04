package main

import (
	"flag"
	"fmt"
	"go-musthave-shortener-tpl/internal/app/config"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var params = config.NewNetAddress()

func parseFlags(p *config.NetAddress) {
	_ = flag.Value(p)
	flag.Var(p, "a", "address and port to run server")
	flag.Func("b", "base url for short link server", func(flagValue string) error {
		// разбиваем значение флага на слайс строк через запятую
		// и заливаем в переменную
		re := regexp.MustCompile(`([a-z]*)://([a-z]*):([0-9]*)`)
		config.BaseURL = re.FindString(flagValue) + "/"
		return nil
	})
	flag.Parse()

	if envRunAddr := os.Getenv("SERVER_ADDRESS"); envRunAddr != "" {
		hp := strings.Split(envRunAddr, ":")
		if len(hp) != 2 {
			fmt.Errorf("need address in a form host:port")
		}
		port, err := strconv.Atoi(hp[1])
		if err != nil {
			panic(err)
		}
		params.Host = hp[0]
		params.Port = port
	}

	if envBaseAddr := os.Getenv("BASE_URL"); envBaseAddr != "" {
		re := regexp.MustCompile(`([a-z]*)://([a-z]*):([0-9]*)`)
		config.BaseURL = re.FindString(envBaseAddr) + "/"
	}
}

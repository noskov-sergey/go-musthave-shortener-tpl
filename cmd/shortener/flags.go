package main

import (
	"errors"
	"flag"
	"go-musthave-shortener-tpl/internal/app/config"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var params = config.NewNetAddress()

func parseFlags(p *config.NetAddress, s *config.Backup, dbConf *config.DataBase) {
	_ = flag.Value(p)
	flag.Var(p, "a", "address and port to run server")
	flag.Func("b", "base url for short link server", func(flagValue string) error {
		re := regexp.MustCompile(`([a-z]*)://([a-z]*):([0-9]*)`)
		config.BaseURL = re.FindString(flagValue) + "/"
		return nil
	})
	flag.Var(s, "f", "base file address for json base")
	flag.Var(dbConf, "d", "base address for data base")
	flag.Parse()

	if envRunAddr := os.Getenv("SERVER_ADDRESS"); envRunAddr != "" {
		hp := strings.Split(envRunAddr, ":")
		if len(hp) != 2 {
			panic(errors.New("need address in a form host:port"))
		}
		port, err := strconv.Atoi(hp[1])
		if err != nil {
			panic(err)
		}
		p.Host = hp[0]
		p.Port = port
	}

	if envBaseAddr := os.Getenv("BASE_URL"); envBaseAddr != "" {
		re := regexp.MustCompile(`([a-z]*)://([a-z]*):([0-9]*)`)
		config.BaseURL = re.FindString(envBaseAddr) + "/"
	}

	if envFileStorage := os.Getenv("FILE_STORAGE_PATH"); envFileStorage != "" {
		s.Set(envFileStorage)
	}

	if envBataBase := os.Getenv("DATABASE_DSN"); envBataBase != "" {
		dbConf.Set(envBataBase)
	}
}

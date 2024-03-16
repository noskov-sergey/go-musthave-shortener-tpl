package main

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	serv "go-musthave-shortener-tpl/internal/app"
	storage "go-musthave-shortener-tpl/internal/app/backup"
	"go-musthave-shortener-tpl/internal/app/config"
	"log"
)

func main() {
	parseFlags(params, config.Fileparams, config.DBConf)
	if !config.DBConf.Active {
		Consumer, err := storage.NewReader(config.Fileparams.String())
		if err != nil {
			log.Fatal(err)
		}
		defer Consumer.Close()
		err = Consumer.ReadFile()
		if err != nil {
			log.Fatal(err, "error from main with ReadFile")
		}
	}

	if config.DBConf.Active {
		db, err := sql.Open("pgx", config.DBConf.Config)
		if err != nil {
			panic(err)
		}
		defer db.Close()

		config.DBConf.Base = db

		err = db.Ping()
		if err != nil {
			log.Fatal(err, "mysql connection failed!")
		}

		err = config.DBConf.CreateNewTable()
		if err != nil {
			log.Fatal(err, "Postresql can't create table")
		}

	}

	if errServ := serv.RunServer(params); errServ != nil {
		panic(errServ)
	}
}

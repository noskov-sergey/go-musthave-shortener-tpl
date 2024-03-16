package config

import (
	"database/sql"
	"errors"
	"strconv"
	"strings"
)

var BaseURL = "http://localhost:8080/"
var Fileparams = NewFileParams()
var DBConf = NewDataBase()

type Backup struct {
	BaseFile string
	W        int
}

func NewFileParams() *Backup {
	return &Backup{
		BaseFile: "./tmp/short-url-db.json",
		W:        0,
	}
}

func (b *Backup) String() string {
	return b.BaseFile
}

func (b *Backup) Set(src string) error {
	b.BaseFile = src
	b.W = 1
	return nil
}

type NetAddress struct {
	Host string
	Port int
}

func NewNetAddress() *NetAddress {
	return &NetAddress{
		Host: "localhost",
		Port: 8080,
	}
}

func (n *NetAddress) String() string {
	return n.Host + ":" + strconv.Itoa(n.Port)
}

func (n *NetAddress) Set(src string) error {
	hp := strings.Split(src, ":")
	if len(hp) != 2 {
		return errors.New("need address in a form host:port")
	}
	port, err := strconv.Atoi(hp[1])
	if err != nil {
		return err
	}
	n.Host = hp[0]
	n.Port = port
	return nil
}

type DataBase struct {
	Config string
	Base   *sql.DB
	Active bool
}

func NewDataBase() *DataBase {
	var db *sql.DB
	return &DataBase{
		Config: "",
		Base:   db,
		Active: false,
	}
}

func (d *DataBase) String() string {
	return d.Config
}

func (d *DataBase) Set(src string) error {
	d.Config = src
	d.Active = true
	return nil
}

func (d *DataBase) CreateNewTable() error {
	query := `
			create table if not exists shorten(
				id SERIAL PRIMARY KEY,
				shorten_uri text,
				original_uri text
			);
		`

	_, err := d.Base.Exec(query)
	return err
}

func (d *DataBase) WriteShorten(key, uri string) error {
	_, err := d.Base.Exec("INSERT INTO shorten(shorten_uri, original_uri) VALUES($1, $2)",
		key,
		uri,
	)
	return err
}

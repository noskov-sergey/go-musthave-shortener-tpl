package config

import (
	"database/sql"
	"errors"
	"github.com/jackc/pgerrcode"
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
			CREATE TABLE if NOT EXISTS shorten(
				id SERIAL PRIMARY KEY,
				shorten_uri text,
				original_uri text UNIQUE
			);
		`

	_, err := d.Base.Exec(query)
	return err
}

func (d *DataBase) WriteShorten(key, uri string) error {
	queryW := `
			INSERT INTO shorten(shorten_uri, original_uri)
			    VALUES($1, $2)
			    ON CONFLICT (original_uri) DO NOTHING;
		`
	e, err := d.Base.Exec(queryW, key, uri)
	val, _ := e.RowsAffected()
	if val == 0 {
		err = errors.New(pgerrcode.UniqueViolation)
	}
	return err
}

func (d *DataBase) ReadOriginal(key string) (string, error) {
	var res string

	row := d.Base.QueryRow("SELECT original_uri FROM shorten WHERE shorten_uri = $1",
		key,
	)
	err := row.Scan(&res)

	return res, err
}

func (d *DataBase) ReadShorten(uri string) (string, error) {
	var res string

	row := d.Base.QueryRow("SELECT shorten_uri FROM shorten WHERE original_uri = $1",
		uri,
	)
	err := row.Scan(&res)

	return res, err
}

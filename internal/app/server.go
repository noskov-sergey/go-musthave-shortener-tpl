package server

import (
	"errors"
	"math/rand"
	"net/http"
	"strings"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Storage struct {
	links map[string]string
}

func (c *Storage) Add(url string) string {
	short := make([]rune, 8)
	for i := range short {
		short[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	key := string(short)
	c.links[key] = url
	return key
}

func (c *Storage) Get(key string) (string, error) {
	url, ok := c.links[key]
	if ok == false {
		return "", errors.New("Key not exist")
	}
	return url, nil
}

var storage = Storage{
	links: map[string]string{},
}

func CreateRedirect(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	url := req.FormValue("url")
	key := "http://localhost:8080/" + storage.Add(url)
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(key))
}

func Redirect(res http.ResponseWriter, req *http.Request) {
	key := strings.TrimPrefix(req.URL.Path, "/")
	url, err := storage.Get(key)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	http.Redirect(res, req, url, http.StatusMovedPermanently)
}

func RouteRedirect(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		Redirect(res, req)
	} else {
		CreateRedirect(res, req)
	}
}

func RunServer() {
	http.HandleFunc("/", RouteRedirect)
	http.ListenAndServe(":8080", nil)
}

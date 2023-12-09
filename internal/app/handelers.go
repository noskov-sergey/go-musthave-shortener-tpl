package server

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func CreateRedirect(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}
	url := string(body)
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
	res.Header().Set("Location", string(url))
	res.WriteHeader(http.StatusTemporaryRedirect)
}

func RouteRedirect(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		Redirect(res, req)
	} else {
		CreateRedirect(res, req)
	}
}

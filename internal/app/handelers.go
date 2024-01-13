package server

import (
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"net/http"
	"strconv"
)

func CreateRedirect(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}
	req.Body.Close()
	url := string(body)
	if url == "" {
		log.Fatalln("should be not empty data")
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	shortkey, err := storage.Add(url)
	if err != nil {
		log.Fatalln(err)
	}
	key := "http://localhost:8080/" + shortkey
	res.Header().Add("Content-Type", "text/plain")
	res.Header().Add("Content-Length", strconv.Itoa(len(key)))
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(key))
}

func Redirect(res http.ResponseWriter, req *http.Request) {
	key := chi.URLParam(req, "shortlink")
	url, err := storage.Get(key)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	res.Header().Set("Location", url)
	res.WriteHeader(http.StatusTemporaryRedirect)
}

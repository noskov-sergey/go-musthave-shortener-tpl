package storage

import (
	"bytes"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go-musthave-shortener-tpl/internal/app/backup"
	"go-musthave-shortener-tpl/internal/app/config"
	"go-musthave-shortener-tpl/internal/app/models"
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
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	shortkey, err := storage.RealStorage.Add(url)
	if err != nil {
		log.Fatalln(err)
	}
	key := config.BaseURL + shortkey

	if config.DBConf.Active {
		err = config.DBConf.WriteShorten(shortkey, url)
		if err != nil {
			log.Fatalln(err)
		}
	}

	res.Header().Add("Content-Type", "text/plain")
	res.Header().Add("Content-Length", strconv.Itoa(len(key)))
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(key))
}

func Redirect(res http.ResponseWriter, req *http.Request) {
	key := chi.URLParam(req, "shortlink")
	url, err := storage.RealStorage.Get(key)
	if config.DBConf.Active {
		url, err = config.DBConf.ReadOriginal(key)
		if err != nil {
			log.Fatalln(err)
		}
	}
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	res.Header().Set("Location", url)
	res.WriteHeader(http.StatusTemporaryRedirect)
}

func APIShorten(res http.ResponseWriter, req *http.Request) {
	var buf bytes.Buffer
	var requestAPI models.RequestShorten
	var responsAPI models.ResponseShorten

	_, err := buf.ReadFrom(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(buf.Bytes(), &requestAPI); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	shortKey, err := storage.RealStorage.Add(requestAPI.URI)
	if err != nil {
		log.Fatalln(err)
	}

	if config.DBConf.Active {
		err = config.DBConf.WriteShorten(shortKey, requestAPI.URI)
		if err != nil {
			log.Fatalln(err)
		}
	}

	responsAPI.Result = config.BaseURL + shortKey

	resp, err := json.Marshal(responsAPI)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	res.Write(resp)
}

func PingAPI(res http.ResponseWriter, req *http.Request) {
	err := config.DBConf.Base.Ping()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}

package models

type ShortenAPI struct {
	URI    string `json:"-"`
	Result string `json:"result"`
}

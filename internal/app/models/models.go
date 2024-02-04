package models

type RequestShorten struct {
	Url string `json:"url"`
}

type ResponseShorten struct {
	Result string `json:"result"`
}

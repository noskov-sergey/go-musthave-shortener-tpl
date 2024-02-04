package models

type RequestShorten struct {
	URI string `json:"url"`
}

type ResponseShorten struct {
	Result string `json:"result"`
}

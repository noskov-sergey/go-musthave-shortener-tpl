package models

type RequestShorten struct {
	Uri string `json:"url"`
}

type ResponseShorten struct {
	Result string `json:"result"`
}

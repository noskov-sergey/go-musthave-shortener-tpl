package models

type RequestShorten struct {
	URI string `json:"url"`
}

type ResponseShorten struct {
	Result string `json:"result"`
}

type BackupModel struct {
	URI         string `json:"short_url"`
	OriginalUri string `json:"original_url"`
}

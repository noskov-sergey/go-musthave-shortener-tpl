package models

type RequestShorten struct {
	URI string `json:"url"`
}

type ResponseShorten struct {
	Result string `json:"result"`
}

type BackupModel struct {
	URI         string `json:"short_url"`
	OriginalURI string `json:"original_url"`
}

type RequestBath struct {
	CorrID      string `json:"correlation_id"`
	OriginalURI string `json:"original_url"`
}

type ResponseBath struct {
	CorrID   string `json:"correlation_id"`
	ShortURI string `json:"short_url"`
}

type BatchMapper struct {
	CorrID      string
	OriginalURI string
	ShortURI    string
}

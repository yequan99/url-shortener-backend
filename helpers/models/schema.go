package models

type UserAuth struct {
	Username  string `json:"Username"`
	HashedPwd string `json:"HashedPwd"`
}

type UserURL struct {
	Username string `json:"Username"`
	ShortURL string `json:"ShortURL"`
	LongURL  string `json:"LongURL"`
}

type UrlTable struct {
	LongURL  string `json:"LongURL"`
	ShortURL string `json:"ShortURL"`
	ShortID  string `json:"ShortID"`
}

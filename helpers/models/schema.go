package models

type UserAuth struct {
	Username  string `json:"Username"`
	HashedPwd string `json:"HashedPwd"`
}

type UrlCode struct {
	Username string `json:"Username"`
	ShortURL string `json:"ShortURL"`
	LongURL  string `json:"LongURL"`
}

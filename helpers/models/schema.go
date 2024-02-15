package models

type UserAuth struct {
	Username  string `json:"Username"`
	HashedPwd string `json:"HashedPwd"`
}

type UserURL struct {
	ID       uint64 `json:"ID"`
	Username string `json:"Username"`
	ShortURL string `json:"ShortURL"`
	LongURL  string `json:"LongURL"`
}

type UserUrlID struct {
	ID     uint64 `json:"ID"`
	LastID uint64 `json:"LastID"`
}

type UrlTable struct {
	LongURL  string `json:"LongURL"`
	ShortURL string `json:"ShortURL"`
	ShortID  string `json:"ShortID"`
}

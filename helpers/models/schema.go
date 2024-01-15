package models

type UserAuth struct {
	UserID    string `json:"UserID`
	Username  string `json:"Username`
	Salt      string `json:"Salt"`
	HashedPwd string `json:"HashedPwd"`
}

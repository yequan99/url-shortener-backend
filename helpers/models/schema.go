package models

type UserAuth struct {
	Username  string `json:"Username`
	Salt      string `json:"Salt"`
	HashedPwd string `json:"HashedPwd"`
}

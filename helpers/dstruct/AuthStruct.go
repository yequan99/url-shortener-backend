package dstruct

type UserLoginCredentials struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type ErrorMessage struct {
	Error string `json:"error"`
}

package dstruct

// ========================
// Auth Microservice
// ========================
type UserLoginCredentials struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type ErrorMessage struct {
	Error string `json:"error"`
}

// ========================
// Webservice Microservice
// ========================

type GenerateShortURL struct {
	LongURL string `json:"Longurl"`
}

type ReturnShortURL struct {
	ShortURL string `json:"ShortURL"`
}

type ReturnUrlArray struct {
	ID       uint64 `json:"ID"`
	ShortURL string `json:"ShortURL"`
	LongURL  string `json:"LongURL"`
}

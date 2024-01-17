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
	Username string `json:"Username"`
	LongURL  string `json:"Longurl"`
}

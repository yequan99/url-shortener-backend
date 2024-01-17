module server

go 1.19

require helpers v1.2.3

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/go-chi/chi v1.5.5 // indirect
	github.com/go-chi/chi/v5 v5.0.11 // indirect
	github.com/rs/cors v1.10.1 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
)

replace helpers => ../helpers

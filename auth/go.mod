module auth

go 1.19

require helpers v1.2.3

require (
	github.com/aws/aws-sdk-go v1.49.21 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/go-chi/chi/v5 v5.0.11 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/rs/cors v1.10.1 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	golang.org/x/crypto v0.18.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
)

replace helpers => ../helpers

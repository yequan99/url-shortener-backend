module webservice

go 1.19

require helpers v1.2.3

require (
	github.com/aws/aws-sdk-go v1.49.21 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/go-chi/chi v1.5.5 // indirect
	github.com/go-chi/chi/v5 v5.0.11 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/rs/cors v1.10.1 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/teris-io/shortid v0.0.0-20220617161101-71ec9f2aa569 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
)

replace helpers => ../helpers

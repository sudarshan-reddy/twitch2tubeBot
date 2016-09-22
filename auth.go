package main

import (
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)

func auth(clientEmail, key, scope1 string) (*http.Client, error) {
	conf := &jwt.Config{
		Email:      clientEmail,
		PrivateKey: []byte(key),
		Scopes:     []string{scope1},
		TokenURL:   google.JWTTokenURL,
		Subject:    "Sudarshan@everywhere.com"}

	client := conf.Client(oauth2.NoContext)
	return client, nil
}

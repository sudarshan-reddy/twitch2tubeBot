package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

func auth(CID, CSecret, scope1, auth, token string) (*http.Client, error) {
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     CID,
		ClientSecret: CSecret,
		Scopes:       []string{scope1},
		Endpoint: oauth2.Endpoint{
			AuthURL:  auth,
			TokenURL: token,
		},
	}

	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}

	tok, err := conf.Exchange(ctx, code)

	return conf.Client(ctx, tok), err
}

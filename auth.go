package main

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
)

func auth(CID, CSecret, scope1 string) {
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     CID,
		ClientSecret: CSecret,
		Scopes:       []string{scope1},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://youtube.com/o/oauth2/auth",
			TokenURL: "https://youtube.com/o/oauth2/token",
		},
	}
	fmt.Println("test")
}

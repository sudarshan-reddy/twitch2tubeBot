package main

import (
	"fmt"

	youtube "google.golang.org/api/youtube/v3"
)

func main() {
	json, _ := getSecrets()
	client, err := auth(json.Web.ClientID,
		json.Web.ClientSecret,
		json.Web.AuthURI,
		json.Web.TokenURI,
		youtube.YoutubeUploadScope)
	if err != nil {
		fmt.Println(err)
	}
	if client != nil {
		fmt.Println("client success")
	}
}

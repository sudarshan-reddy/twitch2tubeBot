package main

import (
	"testing"

	youtube "google.golang.org/api/youtube/v3"
)

func Test_Auth(t *testing.T) {
	json, _ := getSecrets()
	if i, e := auth(json.Web.ClientID,
		json.Web.ClientSecret,
		json.Web.AuthURI,
		json.Web.TokenURI,
		youtube.YoutubeUploadScope); i != nil || e == nil {
		t.Log("test passed")
	} else {
		t.Error("test failed")
	}
}

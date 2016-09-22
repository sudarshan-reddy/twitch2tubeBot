package main

import (
	"testing"

	youtube "google.golang.org/api/youtube/v3"
)

func Test_Auth(t *testing.T) {
	json, _ := getSecrets()
	if i, e := auth(json.ClientEmail, json.PrivateKey,
		youtube.YoutubeUploadScope); i != nil || e == nil {
		t.Log("test passed")
	} else {
		t.Error("test failed")
	}
}

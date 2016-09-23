//no pun intended
package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
)

type secrets struct {
	Typ         string `json:"type"`
	ProjectID   string `json:"project_id"`
	PrivateID   string `json:"private_key_id"`
	PrivateKey  string `json:"private_key"`
	ClientEmail string `json:"client_email"`
	ClientID    string `json:"client_id"`
	AuthURI     string `json:"auth_uri"`
	TokenURI    string `json:"token_uri"`
	X509        string `json:"auth_provider_x509_cert_url"`
	cX509       string `json:"client_x509_cert_url"`
}

func setFileName(videoName string) *string {
	//flag.String("filename", "", "Name of video file to upload")
	return flag.String("filename", videoName, videoName)
}

func setTitle(testTitle, videoTitle string) *string {
	//flag.String("title", "Test Title", "Video title")
	return flag.String("title", testTitle, videoTitle)
}

func setdesc(testDesc, videoDesc string) *string {
	//flag.String("description", "Test Description", "Video description")
	return flag.String("description", testDesc, videoDesc)
}

func setCategory() *string {
	//flag.String("category", "22", "Video category")
	return flag.String("category", "22", "Video category")
}

func setKeywords(listOfKeywords string) *string {
	//flag.String("keywords", "", "Comma separated list of video keywords")
	return flag.String("keywords", "", listOfKeywords)
}

func setPrivacy() *string {
	//flag.String("privacy", "unlisted", "Video privacy status")
	return flag.String("privacy", "unlisted", "Video privacy status")
}

func getSecrets() (*secrets, error) {
	var clientSecret secrets
	file, err := os.Open("twitch2tube-ff26c1d2607c.json")
	data, err := ioutil.ReadAll(file)

	json.Unmarshal(data, &clientSecret)

	return &clientSecret, err

}

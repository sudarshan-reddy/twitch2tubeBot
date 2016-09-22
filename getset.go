package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
)

type secrets struct {
	Web wb
}

type wb struct {
	ClientID     string `json:"client_id"`
	ProjectID    string `json:"project_id"`
	AuthURI      string `json:"auth_uri"`
	TokenURI     string `json:"token_uri"`
	X509         string `json:"auth_provider_x509_cert_url"`
	ClientSecret string `json:"client_secret"`
}

func setFileName(videoName string) *string {
	//flag.String("filename", "", "Name of video file to upload")
	return flag.String("filename", "", videoName)
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
	file, err := os.Open("clientsecret.json")
	data, err := ioutil.ReadAll(file)

	json.Unmarshal(data, &clientSecret)

	return &clientSecret, err

}

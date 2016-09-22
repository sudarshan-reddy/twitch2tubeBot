package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getFileName(downloadLink string) string {
	hash := md5.New()
	io.WriteString(hash, downloadLink)

	var hashStr = hex.EncodeToString(hash.Sum(nil)) + ".mp4" //need to change this

	fmt.Printf("%#v", hashStr)
	return hashStr
}

func downloadVids(downloadLink string) (string, error) {
	resp, err := http.Get(downloadLink)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var file = getFileName(downloadLink)

	if _, err := os.Stat(file); os.IsNotExist(err) {
		out, err := os.Create(file)
		if err != nil {
			fmt.Println(err)
		}
		defer out.Close()

		io.Copy(out, resp.Body)

		fmt.Println("saved file ", file)
		return file, nil
	} else {
		return "", errors.New("file already exists")
	}
}

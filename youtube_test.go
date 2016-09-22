package main

import (
	"os"
	"testing"
)

func Test_DownloadVids(t *testing.T) {
	if i, err := downloadVids("https://clips-media-assets.twitch.tv/23230359872-offset-2418-854x480.mp4"); i != "" || err == nil {
		os.Remove(i)
		t.Log("test passed")
	} else {
		t.Error("Download failed")
	}
}

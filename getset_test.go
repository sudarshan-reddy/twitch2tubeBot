package main

import (
	"testing"
)

func Test_GetSet(t *testing.T) {
	if i, e := getSecrets(); i != nil || e == nil {
		if i.PrivateKey != "" {
			t.Log("test passed")

		}
	} else {
		t.Error("Failed")
	}
}

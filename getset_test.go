package main

import (
	"testing"
)

func Test_GetSet(t *testing.T) {
	if i, e := getSecrets(); i != nil || e == nil {
		t.Log("test passed", i.Web.ClientID)
	} else {
		t.Error("Failed")
	}
}

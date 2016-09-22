package main

import (
	"testing"
)

func Test_GetSet(t *testing.T) {
	if i, e := getSecrets(); i != nil || e == nil {
		if i.Web.X509 != "" && i.Web.AuthURI != "" && i.Web.TokenURI != "" &&
			i.Web.ClientID != "" && i.Web.ProjectID != "" &&
			i.Web.ClientSecret != "" {
			t.Log("test passed")

		}
	} else {
		t.Error("Failed")
	}
}

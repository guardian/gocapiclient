package gocapiclient_test

import (
	"github.com/guardian/gocapiclient"
	"testing"
)

func TestNewCapiClient(t *testing.T) {
	t.Log("Creating a new GuardianContentClient")

	targetUrl := "https://content.guardianapis.com/"
	apiKey := "my-api-key"

	client := gocapiclient.NewGuardianContentClient(targetUrl, apiKey)

	apiKeyIsSet := (client.ApiKey == apiKey)
	targetUrlIsSet := (client.TargetUrl == targetUrl)

	expectedValuesSet := apiKeyIsSet && targetUrlIsSet

	if !expectedValuesSet {
		t.Error("Expected values not set on GuardianContentClient")
	}
}

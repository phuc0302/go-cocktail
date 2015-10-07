package gcm

import (
	"fmt"
	"testing"
)

func TestCreateClient(t *testing.T) {
	client := CreateClient("api_key", GATEWAY)

	if client.ApiKey != "api_key" {
		t.Errorf("Expect 'api_key' but found %s", client.ApiKey)
	}
	if client.HttpClient == nil {
		t.Errorf("Expect HttpClient not nil but found nil")
	}
}

func TestSendMessage(t *testing.T) {
	message := CreateMessage([]string{"1", "2"})
	message.DryRun = true

	client := CreateClient("api_key", GATEWAY)

	responses, err := client.SendMessage(message)
	fmt.Sprintf("%s %s", responses, err)
}

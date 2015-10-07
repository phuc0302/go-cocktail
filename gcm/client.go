package gcm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	ApiKey     string
	Gateway    string
	HttpClient *http.Client
}

// MARK: Struct's constructors
func CreateClient(apiKey string, gateway string) *Client {
	return &Client{
		ApiKey:     apiKey,
		Gateway:    gateway,
		HttpClient: http.DefaultClient,
	}
}

// MARK: Struct's public functions
func (m *Client) SendMessage(templateMessage *Message) ([]*Response, error) {
	/* Condition validation */
	if templateMessage == nil {
		return nil, errors.New("Message cannot be nil.")
	}

	messages := templateMessage.Encode()
	responses := make([]*Response, len(messages))

	idx := 0
	for _, message := range messages {
		response, err := m.send(message)
		if err == nil {
			responses[idx] = response
			idx++
		}
	}

	responses = responses[:idx]
	return responses, nil
}

// MARK: Struct's private functions
func (m *Client) send(message *Message) (*Response, error) {
	// Encode Json
	data, _ := json.Marshal(message)

	// Prepare request
	request, _ := http.NewRequest("POST", m.Gateway, bytes.NewBuffer(data))
	request.Header.Add("Authorization", fmt.Sprintf("key=%s", m.ApiKey))
	request.Header.Add("Content-Type", "application/json")

	// Send request
	httpResponse, err := m.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	// Analyze response status
	if httpResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d Error: %s", httpResponse.StatusCode, httpResponse.Status)
	}

	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	} else {
		result := Response{}
		err = json.Unmarshal(body, &result)
		return &result, err
	}
}

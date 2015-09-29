package gcm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Messenger struct {
	ApiKey     string
	HttpClient *http.Client
}

// MARK: Struct's constructors
func DefaultMessenger(apiKey string) *Messenger {
	return &Messenger{
		ApiKey:     apiKey,
		HttpClient: new(http.Client),
	}
}

// MARK: Struct's public functions
func (m *Messenger) SendMessage(message *Message, toDevices []string) ([]*Response, error) {
	/* Condition validation */
	if message == nil || len(toDevices) == 0 {
		return nil, errors.New("Message or devices cannot be nil.")
	}
	maxDevices := 1000

	length := len(toDevices)
	r := length % maxDevices

	// Calculate step
	c := (length - r) / maxDevices
	if r > 0 {
		c++
	}

	idx := 0
	responses := make([]*Response, c)
	for i := 0; i < c; i++ {
		strIdx := i * maxDevices
		endIdx := strIdx + maxDevices

		if endIdx > length {
			endIdx = length
		}
		devices := toDevices[strIdx:endIdx]

		response, err := m.send(message, devices)
		if err == nil {
			responses[idx] = response
			idx++
		}
	}

	responses = responses[:idx]
	return responses, nil
}

// MARK: Struct's private functions
func (m *Messenger) send(message *Message, toDevices []string) (*Response, error) {
	// Encode Json
	data, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	// Prepare request
	request, _ := http.NewRequest("POST", GATEWAY, bytes.NewBuffer(data))
	request.Header.Add("Authorization", fmt.Sprintf("key=%s", m.ApiKey))
	request.Header.Add("Content-Type", "application/json")

	// Send request
	response, err := m.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Analyze response
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d Error: %s", response.StatusCode, response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	} else {
		result := Response{}
		err = json.Unmarshal(body, &result)
		return &result, err
	}

}

package apns

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	iOS7_MaxPayload = 255
	iOS8_MaxPayload = 2048

	ItemId_DeviceToken = 1
	ItemId_Payload     = 2
	ItemId_Identifier  = 3
	ItemId_Expired     = 4
	ItemId_Priority    = 5
	Length_DeviceToken = 32
	Length_Identifier  = 4
	Length_Expired     = 4
	Length_Priority    = 1
)

var APNsResponses = map[uint8]string{
	0:   "NO_ERRORS",
	1:   "PROCESSING_ERROR",
	2:   "MISSING_DEVICE_TOKEN",
	3:   "MISSING_TOPIC",
	4:   "MISSING_PAYLOAD",
	5:   "INVALID_TOKEN_SIZE",
	6:   "INVALID_TOPIC_SIZE",
	7:   "INVALID_PAYLOAD_SIZE",
	8:   "INVALID_TOKEN",
	10:  "SHUTDOWN",
	255: "UNKNOWN",
}

type APNsResponse struct {
	Success       bool
	AppleResponse string
	Error         error
}

type Alert struct {
	Body         string   `json:"body,omitempty"`
	LaunchImage  string   `json:"launch-image,omitempty"`
	ActionLocKey string   `json:"action-loc-key,omitempty"`
	LocKey       string   `json:"loc-key,omitempty"`
	LocArgs      []string `json:"loc-args,omitempty"`
}

type Payload struct {
	Alert            interface{} `json:"alert,omitempty"`
	Badge            uint        `json:"badge,omitempty"`
	Sound            string      `json:"sound,omitempty"`
	Category         string      `json:"category,omitempty"`
	ContentAvailable int         `json:"content-available,omitempty"`
}

type APNs struct {
	Id          int32
	Expired     uint32
	DeviceToken string
	OsVersion   string

	priority uint8
	payload  map[string]interface{}
}

// MARK: Struct's constructors
func CreateApns(deviceToken string, oVversion string) *APNs {
	apns := APNs{
		Id:          rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(9999),
		DeviceToken: deviceToken,
		OsVersion:   oVversion,
		priority:    10,
		payload:     make(map[string]interface{}),
	}
	return &apns
}

// MARK: Struct's public functions
func (a *APNs) Encode() ([]byte, error) {
	token, err := base64.StdEncoding.DecodeString(a.DeviceToken)

	/* Condition validation: Validate decoded device's token */
	if err != nil || len(token) != Length_DeviceToken {
		return nil, errors.New("Device's token is not base64 format or invalid token's length.")
	}

	// Encode payload
	payload, err := json.Marshal(a.payload)
	if err != nil {
		return nil, err
	}

	/* Condition validation: Validate payload's token */
	strings := strings.Split(a.OsVersion, ".")
	v, err := strconv.Atoi(strings[0])

	if err == nil && v >= 8 && len(payload) > iOS8_MaxPayload {
		return nil, errors.New(fmt.Sprintf("Payload is larger than: %i bytes.", iOS8_MaxPayload))
	} else if len(payload) > iOS7_MaxPayload {
		return nil, errors.New(fmt.Sprintf("Payload is larger than: %i bytes.", iOS7_MaxPayload))
	}

	// Encode message
	buffer := bytes.NewBuffer([]byte{})
	binary.Write(buffer, binary.BigEndian, uint8(ItemId_DeviceToken))
	binary.Write(buffer, binary.BigEndian, uint16(Length_DeviceToken))
	binary.Write(buffer, binary.BigEndian, token)

	binary.Write(buffer, binary.BigEndian, uint8(ItemId_Payload))
	binary.Write(buffer, binary.BigEndian, uint16(len(payload)))
	binary.Write(buffer, binary.BigEndian, payload)

	binary.Write(buffer, binary.BigEndian, uint8(ItemId_Identifier))
	binary.Write(buffer, binary.BigEndian, uint16(Length_Identifier))
	binary.Write(buffer, binary.BigEndian, a.Id)

	binary.Write(buffer, binary.BigEndian, uint8(ItemId_Expired))
	binary.Write(buffer, binary.BigEndian, uint16(Length_Expired))
	binary.Write(buffer, binary.BigEndian, a.Expired)

	binary.Write(buffer, binary.BigEndian, uint8(ItemId_Priority))
	binary.Write(buffer, binary.BigEndian, uint16(Length_Priority))
	binary.Write(buffer, binary.BigEndian, a.priority)

	return buffer.Bytes(), nil
}
func (a *APNs) EncodeJson() (string, error) {
	bytes, err := json.Marshal(a.payload)
	return string(bytes), err
}

func (a *APNs) SetField(key string, value interface{}) {
	/* Condition validation */
	if len(key) == 0 || key == "aps" {
		return
	}
	a.payload[key] = value
}

func (a *APNs) SetPayload(p *Payload) {
	a.payload["aps"] = p
}

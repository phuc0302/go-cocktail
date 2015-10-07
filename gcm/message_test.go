package gcm

import (
	"fmt"
	"testing"
)

func TestCreateMessage(t *testing.T) {
	data := []struct {
		registrationIds []string
		result          *Message
	}{
		{nil, nil},
		{make([]string, 0), nil},
	}

	for _, test := range data {
		if m := CreateMessage(test.registrationIds); m != test.result {
			t.Errorf("Expect nil but found %s", m)
		}
	}
}

func TestEncode1(t *testing.T) {
	registrationIds := []string{"1"}
	originalMessage := CreateMessage(registrationIds)
	originalMessage.RestrictedPackageName = "test"
	originalMessage.Data["key"] = "value"

	messages := originalMessage.Encode()
	if messages == nil {
		t.Error("Expect not nil but found nil")
	} else if len(messages) != 1 {
		t.Errorf("Expect only 1 message after encoded but found %d", len(messages))
	}

	if len(messages) >= 1 {
		if messages[0].RestrictedPackageName != originalMessage.RestrictedPackageName {
			t.Errorf("Expect %s but found %s", originalMessage.RestrictedPackageName, messages[0].RestrictedPackageName)
		}

		if messages[0].Data["key"] != originalMessage.Data["key"] {
			t.Errorf("Expect %s but found %s", originalMessage.Data["key"], messages[0].Data["key"])
		}
	}
}
func TestEncode1000(t *testing.T) {
	registrationIds := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		registrationIds[i] = fmt.Sprintf("%d", i)
	}
	originalMessage := CreateMessage(registrationIds)

	messages := originalMessage.Encode()
	if messages == nil {
		t.Error("Expect not nil but found nil")
	} else if len(messages) != 1 {
		t.Errorf("Expect only 1 message after encoded but found %d", len(messages))
	}
}
func TestEncode1001(t *testing.T) {
	registrationIds := make([]string, 1001)
	for i := 0; i < 1001; i++ {
		registrationIds[i] = fmt.Sprintf("%d", i)
	}
	originalMessage := CreateMessage(registrationIds)

	messages := originalMessage.Encode()
	if messages == nil {
		t.Error("Expect not nil but found nil")
	} else if len(messages) != 2 {
		t.Errorf("Expect only 1 message after encoded but found %d", len(messages))
	} else if len(originalMessage.RegistrationIds) != 1001 {
		t.Error("Original registration Ids list had been edited")
	}
}

func TestSetField(t *testing.T) {
	data := []struct {
		key   string
		value interface{}
	}{
		{"", nil},
		{"key", nil},
		{"key", "value"},
		{"number", 100},
	}

	registrationIds := []string{"1"}
	message := CreateMessage(registrationIds)

	for _, test := range data {
		message.SetField(test.key, test.value)

		if message.Data[test.key] != test.value {
			t.Errorf("Expect %s but found %s", message.Data[test.key])
		}
	}
}

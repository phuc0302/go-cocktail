package gcm

type Message struct {
	RegistrationIds []string `json:"registration_ids"`
	CollapseKey     string   `json:"collapse_key,omitempty"`
	Priority        string   `json:"priority,omitempty"`
	DryRun          bool     `json:"dry_run,omitempty"`
	DelayWhileIdle  bool     `json:"delay_while_idle,omitempty"`

	Data map[string]interface{} `json:"data,omitempty"`

	RestrictedPackageName string `json:"restricted_package_name,omitempty"`
}

// MARK: Struct's constructors
func CreateMessage(registrationIds []string) *Message {
	/* Condition validation */
	if len(registrationIds) == 0 {
		return nil
	}

	return &Message{
		RegistrationIds: registrationIds,
		Priority:        "high",

		Data: make(map[string]interface{}),
	}
}

/**
 * Enforce the message. Split original message into multiple messages if required.
 */
func (m *Message) Encode() []*Message {
	length := len(m.RegistrationIds)
	maxIds := 1000

	if length <= maxIds {
		return []*Message{m}
	} else {
		// Calculate step
		remain := length % maxIds
		counter := (length - remain) / maxIds
		if remain > 0 {
			counter++
		}

		// Create message collection
		messages := make([]*Message, counter)
		for i := 0; i < counter; i++ {
			strIdx := i * maxIds
			endIdx := strIdx + maxIds

			/* Condition validation: Validate upper bound */
			if endIdx > length {
				endIdx = length
			}
			ids := m.RegistrationIds[strIdx:endIdx]

			message := *m
			message.RegistrationIds = ids

			messages[i] = &message
		}
		return messages
	}
}

/**
 * Add custom key-value pair to the message's payload.
 */
func (m *Message) SetField(key string, value interface{}) {
	/* Condition validation */
	if len(key) == 0 || value == nil {
		return
	}
	m.Data[key] = value
}

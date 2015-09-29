package gcm

type Message struct {
	RegistrationIds       []string `json:"registration_ids"`
	CollapseKey           string   `json:"collapse_key,omitempty"`            // Identify a group of messages (e.g., with collapse_key: "Updates Available") that can be collapsed, so that only the last message gets sent.
	Priority              string   `json:"priority,omitempty"`                // Valid values are "normal" and "high". When a message is sent with high priority, it is sent immediately, and the app can wake a sleeping device.
	RestrictedPackageName string   `json:"restricted_package_name,omitempty"` // Specify the package name of the application where the registration tokens must match in order to receive the message.

	DryRun         bool `json:"dry_run,omitempty"`          // If set to true, allows developers to test a request without actually sending a message.
	DelayWhileIdle bool `json:"delay_while_idle,omitempty"` // If set to true, it indicates that the message should not be sent until the device becomes active.

	Data map[string]interface{} `json:"data,omitempty"` // Specify the custom key-value pairs of the message's payload.
}

//type IndividualMessage struct {
//	Message
//}

//type MulticastMessage struct {
//	Message
//}

func NewMessage(data map[string]interface{}, regIDs ...string) *Message {
	return &Message{RegistrationIds: regIDs, Data: data}
}

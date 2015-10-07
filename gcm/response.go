package gcm

const (
	DEVICE_MESSAGE_RATE_EXCEEDED = "DeviceMessageRateExceeded"
	INTERNAL_SERVER_ERROR        = "InternalServerError"
	INVALID_DATA_KEY             = "InvalidDataKey"
	INVALID_PACKAGE_NAME         = "InvalidPackageName"
	INVALID_REGISTRATION_TOKEN   = "InvalidRegistration"
	INVALID_TIME_TO_LIVE         = "InvalidTtl"
	MESSAGE_TOO_BIG              = "MessageTooBig"
	MISMATCHED_SENDER            = "MismatchSenderId"
	MISSING_REGISTRATION_TOKEN   = "MissingRegistration"
	TOPICS_MESSAGE_RATE_EXCEEDED = "TopicsMessageRateExceeded"
	TIMEOUT                      = "Unavailable"
	UNREGISTERED_DEVICE          = "NotRegistered"
)

type Response struct {
	MulticastId  int64 `json:"multicast_id"`
	Success      int   `json:"success"`
	Failure      int   `json:"failure"`
	CanonicalIds int   `json:"canonical_ids"`

	Results []Result `json:"results"`
}

////////////////////////////////////////////////////////////////////////////////
//-- Result ------------------------------------------------------------------//
type Result struct {
	MessageId string `json:"message_id"`
	Error     string `json:"error"`

	RegistrationId    string `json:"-"`
	NewRegistrationId string `json:"registration_id"`
}

/**
 * Validate if should not send a new message for a period of time.
 */
func (r *Result) ShouldPostpone() bool {
	if r.Error == TIMEOUT || r.Error == INTERNAL_SERVER_ERROR || r.Error == DEVICE_MESSAGE_RATE_EXCEEDED || r.Error == TOPICS_MESSAGE_RATE_EXCEEDED {
		return true
	} else {
		return false
	}
}

/**
 * Validate error to decide if registrationId should be deleted or not.
 */
func (r *Result) ShouldRemove() bool {
	if r.Error == UNREGISTERED_DEVICE {
		return true
	} else {
		return false
	}
}

/**
 * Validate if registrationId should be update or not.
 */
func (r *Result) ShouldUpdate() bool {
	if len(r.MessageId) > 0 && len(r.NewRegistrationId) > 0 {
		return true
	} else {
		return false
	}
}

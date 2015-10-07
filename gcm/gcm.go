package gcm

const (
	GATEWAY = "https://android.googleapis.com/gcm/send"
)

const (
	BACKOFF_INITIAL_DELAY = 1000
	MAX_BACKOFF_DELAY     = 1024000
)

//https://developers.google.com/cloud-messaging/http
//https://developers.google.com/cloud-messaging/http-server-ref#downstream-http-messages-json

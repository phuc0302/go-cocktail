package cocktail

import (
	"encoding/json"
	"net/http"
	"time"
)

type RequestLog struct {
	Uri           string `json:"uri"`
	UserAgent     string `json:"user_agent"`
	HttpReferer   string `json:"http_referer"`
	RemoteAddress string `json:"remote_address"`
	ContentType   string `json:"content_type"`
	RequestBody   string `json:"request_body"`
}

type RecoveryLog struct {
	Request RequestLog `json:"request"`
	Date    string     `json:"date"`
	Message string     `json:"message"`
	Trace   []string   `json:"trace"`
}

/**
 * Create default recovery log with time stamp
 */
func CreateRecoveryLog(request *http.Request) (*RecoveryLog, time.Time) {
	end := time.Now().UTC()
	log := RecoveryLog{}

	log.Date = end.Format(time.RFC822)
	log.Request.Uri = request.RequestURI
	log.Request.UserAgent = request.UserAgent()
	log.Request.HttpReferer = request.Referer()
	log.Request.RemoteAddress = request.RemoteAddr
	log.Request.ContentType = request.Header.Get("Content-Type")

	// // Read request body
	// bytes := make([]byte, request.ContentLength)
	// _, err := request.Body.Read(bytes)

	// if err == nil {
	// 	log.Request.RequestBody = base64.StdEncoding.EncodeToString(bytes)
	// }
	return &log, end
}

func (r *RecoveryLog) String() string {
	jsonBytes, _ := json.MarshalIndent(r, "", "  ")
	return string(jsonBytes)
}

package common

import (
	"encoding/base64"
	"net/http"
	"time"
)

type LogRequest struct {
	Uri           string `json:"uri,omitempty"`
	UserAgent     string `json:"user_agent,omitempty"`
	HttpReferer   string `json:"http_referer,omitempty"`
	RemoteAddress string `json:"remote_address,omitempty"`
	ContentType   string `json:"content_type,omitempty"`
	RequestBody   string `json:"request_body,omitempty"`
}

type LogRecovery struct {
	Request LogRequest `json:"request,omitempty"`
	Date    string     `json:"date,omitempty"`
	Message string     `json:"message,omitempty"`
	Trace   []string   `json:"trace,omitempty"`
}

/** Create default recovery log with time stamp. */
func CreateLog(request *http.Request) (*LogRecovery, time.Time) {
	end := time.Now().UTC()
	log := LogRecovery{}

	log.Date = end.Format(time.RFC822)
	log.Request.Uri = request.RequestURI
	log.Request.UserAgent = request.UserAgent()
	log.Request.HttpReferer = request.Referer()
	log.Request.RemoteAddress = request.RemoteAddr
	log.Request.ContentType = request.Header.Get("Content-Type")

	// Read request body
	bytes := make([]byte, request.ContentLength)
	_, err := request.Body.Read(bytes)

	if err == nil {
		log.Request.RequestBody = base64.StdEncoding.EncodeToString(bytes)
	}
	return &log, end
}

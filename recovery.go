package cocktail

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	dunno     = []byte("???")
	centerDot = []byte("·")
	dot       = []byte(".")
	slash     = []byte("/")
)

/**
 * Recovery returns a middleware that recovers from any panics and writes a 500
 * if there was one.
 */
func Recovery(request *http.Request, response http.ResponseWriter, logger *log.Logger) {
	if err := recover(); err != nil {
		// log, time := CreateRecoveryLog(request)
		log, _ := createLog(request)
		log.Message = fmt.Sprintf("%s", err)
		log.Trace = getStack(3)

		// Write error to file

		// Return error
		httpError := InternalServerError()
		WriteError(response, httpError)
	}
}

type requestLog struct {
	Uri           string `json:"uri"`
	UserAgent     string `json:"user_agent"`
	HttpReferer   string `json:"http_referer"`
	RemoteAddress string `json:"remote_address"`
	ContentType   string `json:"content_type"`
	RequestBody   string `json:"request_body"`
}

type recoveryLog struct {
	Request requestLog `json:"request"`
	Date    string     `json:"date"`
	Message string     `json:"message"`
	Trace   []string   `json:"trace"`
}

/** Create default recovery log with time stamp. */
func createLog(request *http.Request) (*recoveryLog, time.Time) {
	end := time.Now().UTC()
	log := recoveryLog{}

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

package cocktail

import (
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
)

/**
 * IHandler represent a routing handler function.
 *
 *  Injector will inject these parameters dynamically as function's inputs
 *    + http.Header           (Optional)
 *    + *http.Request         (Optional)
 *    + http.ResponseWriter   (Optional)
 *
 *    + url.Values            (Optional)
 *    + cocktail.FileParams   (Optional)
 *    + cocktail.PathParams   (Optional)
 *
 *
 *  Function should only return one or two parameter(s)
 *    + cocktail.HttpStatus   (Optional)
 *    + struct or string      (Optional)
 *    + template              (Optional)  (html/template)
 */
type Context struct {
	Queries    url.Values
	PathParams map[string]string
	FileParams map[string][]*multipart.FileHeader

	request  *http.Request
	response http.ResponseWriter
}

// MARK: Struct's public functions
func (c *Context) RenderError(status *Status) {
	c.response.Header().Set("Content-Type", "application/problem+json")
	c.response.WriteHeader(status.Status)

	cause, _ := json.Marshal(status)
	c.response.Write(cause)
}
func (c *Context) RenderJson() {
}
func (c *Context) RenderHtml() {
}

func (c *Context) Recovery(logger *log.Logger) {
	if err := recover(); err != nil {
		// log, time := CreateRecoveryLog(request)
		log, _ := createLog(c.request)
		log.Message = fmt.Sprintf("%s", err)
		log.Trace = getStack(3)

		// Write error to file

		// Return error
		httpError := InternalServerError()
		c.RenderError(httpError)
	}
}

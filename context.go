package cocktail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
	"text/template"

	"github.com/phuc0302/go-cocktail/common"
)

var (
	dunno     = []byte("???")
	centerDot = []byte("·")
	dot       = []byte(".")
	slash     = []byte("/")
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
func (c *Context) RenderError(status *common.Status) {
	c.response.Header().Set("Content-Type", "application/problem+json")
	c.response.WriteHeader(status.Status)

	cause, _ := json.Marshal(status)
	c.response.Write(cause)
}
func (c *Context) RenderJson(status *common.Status, model interface{}) {
	c.response.Header().Set("Content-Type", "application/json")
	c.response.WriteHeader(status.Status)

	data, _ := json.Marshal(model)
	c.response.Write(data)
}
func (c *Context) RenderHtml(filePath string, model interface{}) {
	tmpl, error := template.ParseFiles(filePath)
	if error != nil {
		c.RenderError(common.Status404())
	} else {
		tmpl.Execute(c.response, model)
	}
}

func (c *Context) Recovery(logger *log.Logger) {
	if err := recover(); err != nil {
		log, _ := common.CreateLog(c.request)
		log.Message = fmt.Sprintf("%s", err)
		log.Trace = c.callStack(3)

		// Write error to file

		// Return error
		httpError := common.Status500()
		httpError.Detail = log
		c.RenderError(httpError)
	}
}

// MARK: Struct's private functions
func (c *Context) callStack(skip int) []string {
	// FIX FIX FIX: What if we have more than 1 go path???
	srcPath := fmt.Sprintf("%s/src", os.Getenv("GOPATH"))
	traces := make([]string, 5)

	for i, j := skip, 0; ; i++ {
		// Condition validation: Stop if there is nothing else
		pc, file, line, ok := runtime.Caller(i)
		if !ok || j >= 5 {
			break
		}
		fmt.Println(file, line)

		// Condition validation: Skip go root
		if !strings.HasPrefix(file, srcPath) {
			continue
		}

		// Trim prefix
		file = file[len(srcPath):]

		// Print this much at least. If we can't find the source, it won't show.
		traces[j] = fmt.Sprintf("%s: %s (%d)", file, c.callFunction(pc), line)
		j++
	}
	return traces
}
func (c *Context) callFunction(pc uintptr) string {
	fn := runtime.FuncForPC(pc)

	// Condition validation: return don't know if function is not available
	if fn == nil {
		return string(dunno)
	}

	// Convert function name to byte array for modification
	name := []byte(fn.Name())

	// Eliminate the path prefix
	if lastslash := bytes.LastIndex(name, slash); lastslash >= 0 {
		name = name[lastslash+1:]
	}

	// Eliminate period prefix
	if period := bytes.Index(name, dot); period >= 0 {
		name = name[period+1:]
	}

	// Convert center dot to dot
	name = bytes.Replace(name, centerDot, dot, -1)
	return string(name)
}

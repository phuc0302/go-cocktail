package cocktail

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/phuc0302/go-cocktail-di"
)

type Route struct {
	Pattern string

	regex    *regexp.Regexp
	handlers map[string]interface{}
}

// MARK: Struct's constructors
func createRoute(pattern string) *Route {
	regex := regexp.MustCompile(`:[^/#?()\.\\]+`)

	// Convert param to regular expression
	regexPattern := regex.ReplaceAllStringFunc(pattern, func(m string) string {
		return fmt.Sprintf(`(?P<%s>[^/#?]+)`, m[1:])
	})
	regexPattern += "/?"

	route := Route{pattern, regexp.MustCompile(regexPattern), make(map[string]interface{}, 7)}
	return &route
}

// MARK: Struct's public functions
func (r *Route) AddHandler(method string, handler interface{}) {
	if reflect.TypeOf(handler).Kind() != reflect.Func {
		panic("Request handler must be a function type.")
	}

	method = strings.ToUpper(method)
	r.handlers[method] = handler
}

func (r *Route) Match(method string, urlPath string) (bool, map[string]string) {
	// Condition validation: Match request url
	matches := r.regex.FindStringSubmatch(urlPath)
	if len(matches) == 0 || matches[0] != urlPath {
		return false, nil
	}

	// Condition validation: Match request method
	handler := r.handlers[method]
	if handler == nil {
		return false, nil
	}

	// Extract path params
	params := make(map[string]string)
	for i, name := range r.regex.SubexpNames() {
		if len(name) > 0 {
			params[name] = matches[i]
		}
	}
	return true, params
}

func (r *Route) InvokeHandler(c *Context) {
	switch c.request.Method {
	case POST, PATCH:
		contentType := strings.ToLower(c.request.Header.Get("CONTENT-TYPE"))

		if strings.Contains(contentType, "application/x-www-form-urlencoded") {
			params := ParseForm(c.request)
			if len(params) > 0 {
				c.Queries = params
			}
		} else if strings.Contains(contentType, "multipart/form-data") {
			params, fileParams := ParseMultipartForm(c.request)

			if len(fileParams) > 0 {
				c.FileParams = fileParams
			}
			if len(params) > 0 {
				c.Queries = params
			}
		} else {
			// Do nothing, let the runtime decide
		}
		break

	default:
		params := c.request.URL.Query()
		if len(params) > 0 {
			c.Queries = params
		}
		break
	}

	injector := di.Injector()
	handler := r.handlers[c.request.Method]

	// Call handler
	injector.Map(c)
	_, err := injector.Invoke(handler)

	// Condition validation: Validate error
	if err != nil {
		panic(err)
	}
}

// MARK: Struct's private functions

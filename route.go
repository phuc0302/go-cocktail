package cocktail

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strings"

	"github.com/phuc0302/go-cocktail-di"
)

type Route struct {
	pattern string
	regex   *regexp.Regexp

	handlers map[string]HandlerFunc
}

// MARK: Struct's constructors
func createRoute(pattern string) *Route {
	regex := regexp.MustCompile(`:[^/#?()\.\\]+`)

	// Convert param to regular expression
	regexPattern := regex.ReplaceAllStringFunc(pattern, func(m string) string {
		return fmt.Sprintf(`(?P<%s>[^/#?]+)`, m[1:])
	})
	regexPattern += "/?"

	route := Route{pattern, regexp.MustCompile(regexPattern), make(map[string]HandlerFunc, len(HTTP_METHODS))}
	return &route
}

// MARK: IRoute interface's members
func (r *Route) Pattern() string {
	return r.pattern
}

func (r *Route) AddHandler(method string, handler HandlerFunc) {
	if reflect.TypeOf(handler).Kind() != reflect.Func {
		panic("Request handler must be a function type.")
	}

	method = strings.ToUpper(method)
	r.handlers[method] = handler
}

func (r *Route) Match(method string, urlPath string) (bool, PathParams) {
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
	params := make(PathParams)
	for i, name := range r.regex.SubexpNames() {
		if len(name) > 0 {
			params[name] = matches[i]
		}
	}
	return true, params
}

func (r *Route) InvokeHandler(request *http.Request, response http.ResponseWriter, pathParams PathParams) {
	injector := r.prepareInjector(request, pathParams)
	handler := r.handlers[request.Method]

	// Call handler
	injector.Map(request)
	injector.Map(response)
	injector.Map(request.Header)
	values, err := injector.Invoke(handler)

	// Condition validation: Validate error
	if err != nil {
		panic(err)
	}

	// if the handler returned something, write it to the http response
	if len(values) == 1 {
		var responseVal reflect.Value

		if len(values) > 1 && values[0].Kind() == reflect.Int {
			response.WriteHeader(int(values[0].Int()))
			responseVal = values[1]
		} else if len(values) > 0 {
			responseVal = values[0]
		}

		if canDeref(responseVal) {
			responseVal = responseVal.Elem()
		}

		if isByteSlice(responseVal) {
			response.Write(responseVal.Bytes())
		} else {
			response.Write([]byte(responseVal.String()))
		}
	}
	// else {
	// 	panic("Invalid return value, should be interface.")
	// }
}

// MARK: Struct's private functions
func (r *Route) prepareInjector(request *http.Request, pathParams PathParams) di.IInjector {
	injector := di.Injector()

	// Inject path params
	if len(pathParams) > 0 {
		injector.Map(pathParams)
	}

	// Inject params
	switch request.Method {
	case POST, PATCH:
		contentType := strings.ToLower(request.Header.Get("CONTENT-TYPE"))

		if strings.Contains(contentType, "application/x-www-form-urlencoded") {
			params := ExtractInputForm(request)
			if len(params) > 0 {
				injector.Map(params)
			}
		} else if strings.Contains(contentType, "multipart/form-data") {
			params, fileParams := ExtractMultipartForm(request)
			if len(fileParams) > 0 {
				injector.Map(fileParams)
			}
			if len(params) > 0 {
				injector.Map(params)
			}
		} else {
			// Do nothing, let the runtime decide
		}
		break

	default:
		params := request.URL.Query()
		if len(params) > 0 {
			injector.Map(params)
		}
		break
	}
	return injector
}

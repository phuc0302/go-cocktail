package cocktail

import "net/http"

/**
 * An interface representing a route in routing layer.
 */
type IRoute interface {

	// Pattern returns the pattern of the route.
	Pattern() string

	// Add handler for method
	AddHandler(method string, handler IHandler)

	// Match request's method & path.
	Match(method string, urlPath string) (bool, PathParams)

	// Invoke handler for method
	InvokeHandler(request *http.Request, response http.ResponseWriter, pathParams PathParams)
}

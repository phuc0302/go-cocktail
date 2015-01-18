package cocktail

import (
	"log"
	"net/http"
	"strings"
)

type Router struct {
	groups []string
	routes []IRoute

	logger *log.Logger
}

// MARK: Struct's constructors
func DefaultRouter(logger *log.Logger) IRouter {
	return &Router{groups: make([]string, 0), logger: logger}
}

// MARK: IRouter interface's members
func (r *Router) Group(urlGroup string, function func(router IRouter)) {
	r.groups = append(r.groups, urlGroup)
	function(r)

	r.groups = r.groups[:len(r.groups)-1]
}

func (r *Router) Get(urlPath string, handler IHandler) {
	r.addRoute(GET, urlPath, handler)
}
func (r *Router) Post(urlPath string, handler IHandler) {
	r.addRoute(POST, urlPath, handler)
}
func (r *Router) Patch(urlPath string, handler IHandler) {
	r.addRoute(PATCH, urlPath, handler)
}
func (r *Router) Delete(urlPath string, handler IHandler) {
	r.addRoute(DELETE, urlPath, handler)
}

func (r *Router) Put(urlPath string, handler IHandler) {
	r.addRoute(PUT, urlPath, handler)
}

func (r *Router) HandleRequest(request *http.Request, response http.ResponseWriter) {
	// Condition validation: Validate request method
	isAllowed := false
	for _, allowedMethod := range HTTP_METHODS {
		if request.Method == allowedMethod {
			isAllowed = true
			break
		}
	}

	if !isAllowed {
		WriteError(response, Error405())
		return
	}

	// Condition validation: Find matched route
	for _, route := range r.routes {
		// Match url & extract path params
		ok, pathParams := route.Match(request.Method, request.URL.Path)
		if !ok {
			continue
		}

		route.InvokeHandler(request, response, pathParams)
		return
	}

	// Not Found
	WriteError(response, Error404())
}

// MARK: Struct's private functions
func (r *Router) addRoute(method string, pattern string, handler IHandler) {
	// Condition validation: If pattern belong to group or not
	if len(r.groups) > 0 {
		groupPattern := ""

		for _, g := range r.groups {
			groupPattern += g
		}
		pattern = groupPattern + pattern
	}

	// Look for existing one before create new
	for _, route := range r.routes {
		if route.Pattern() == pattern {
			route.AddHandler(method, handler)
			r.logger.Printf("%-6s -> %s\n", strings.ToUpper(method), route.Pattern())
			return
		}
	}

	// Create new route
	newRoute := CreateRoute(pattern)
	newRoute.AddHandler(method, handler)
	r.logger.Printf("%-6s -> %s\n", strings.ToUpper(method), newRoute.Pattern())

	// Add to collection
	r.routes = append(r.routes, newRoute)
}

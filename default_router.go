package cocktail

import (
	"fmt"
	"log"
	"strings"
)

type DefaultRouter struct {
	routes  []Route
	groups  []string
	methods []string
	logger  *log.Logger
}

// MARK: Struct's constructors
func createDefaultRouter(logger *log.Logger) *DefaultRouter {
	return &DefaultRouter{
		groups:  make([]string, 0),
		methods: []string{DELETE, GET, PATCH, POST},
		logger:  logger,
	}
}

// MARK: Struct's private functions
func (d *DefaultRouter) addRoute(method string, pattern string, handler interface{}) {
	// Condition validation: If pattern belong to group or not
	if len(d.groups) > 0 {
		groupPattern := ""

		for _, g := range d.groups {
			groupPattern += g
		}
		pattern = fmt.Sprintf("%s%s", groupPattern, pattern)
	}

	// Format pattern before assigned to route
	pattern = FormatPath(pattern)

	// Look for existing one before create new
	for _, route := range d.routes {
		if route.GetPattern() == pattern {
			route.AddHandler(method, handler)
			//			d.Logger.Printf("%-6s -> %s\n", strings.ToUpper(method), route.Pattern)
			return
		}
	}

	// Create new route
	newRoute := createDefaultRoute(pattern)
	newRoute.AddHandler(method, handler)

	// Add to collection
	d.routes = append(d.routes, newRoute)
	d.logger.Printf("%-6s -> %s\n", strings.ToUpper(method), newRoute.Pattern)
}

// MARK: Router's members
func (d *DefaultRouter) ShouldAllow(method string) bool {
	isAllowed := false
	for _, allowedMethod := range d.methods {
		if method == allowedMethod {
			isAllowed = true
			break
		}
	}
	return isAllowed
}
func (d *DefaultRouter) ServeRequest(context *Context) bool {
	isHandled := false

	// FIX FIX FIX: Add priority here so that we can move the mosted used node to top
	for _, route := range d.routes {
		ok, pathParams := route.Match(context.Method, context.UrlPath)
		if !ok {
			continue
		}

		context.Params.PathQueries = pathParams
		route.InvokeHandler(context)
		isHandled = true
		break
	}

	return isHandled
}

func (d *DefaultRouter) Group(urlGroup string, function func(r Router)) {
	d.groups = append(d.groups, urlGroup)
	function(d)

	d.groups = d.groups[:len(d.groups)-1]
}
func (d *DefaultRouter) Delete(urlPath string, handler interface{}) {
	defer RecoveryInternal(d.logger)
	d.addRoute(DELETE, urlPath, handler)
}
func (d *DefaultRouter) Get(urlPath string, handler interface{}) {
	defer RecoveryInternal(d.logger)
	d.addRoute(GET, urlPath, handler)
}
func (d *DefaultRouter) Patch(urlPath string, handler interface{}) {
	defer RecoveryInternal(d.logger)
	d.addRoute(PATCH, urlPath, handler)
}
func (d *DefaultRouter) Post(urlPath string, handler interface{}) {
	defer RecoveryInternal(d.logger)
	d.addRoute(POST, urlPath, handler)
}

func (d *DefaultRouter) Copy(urlPath string, handler interface{}) {
	defer RecoveryInternal(d.logger)
	panic(Status405())
}
func (d *DefaultRouter) Head(urlPath string, handler interface{}) {
	defer RecoveryInternal(d.logger)
	panic(Status405())
}
func (d *DefaultRouter) Link(urlPath string, handler interface{}) {
	defer RecoveryInternal(d.logger)
	panic(Status405())
}
func (d *DefaultRouter) Options(urlPath string, handler interface{}) {
	defer RecoveryInternal(d.logger)
	panic(Status405())
}
func (d *DefaultRouter) Purge(urlPath string, handler interface{}) {
	defer RecoveryInternal(d.logger)
	panic(Status405())
}
func (d *DefaultRouter) Put(urlPath string, handler interface{}) {
	defer RecoveryInternal(d.logger)
	panic(Status405())
}
func (d *DefaultRouter) Unlink(urlPath string, handler interface{}) {
	defer RecoveryInternal(d.logger)
	panic(Status405())
}

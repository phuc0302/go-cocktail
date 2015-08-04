package cocktail

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/phuc0302/go-cocktail/common"
	"github.com/phuc0302/go-cocktail/utils"
)

const (
	DELETE = "DELETE"
	GET    = "GET"
	HEAD   = "HEAD"
	PATCH  = "PATCH"
	POST   = "POST"
	PUT    = "PUT"
)

type Cocktail struct {
	Logger       *log.Logger
	Host         string
	Port         string
	StaticFolder string

	routes  []Route
	groups  []string
	methods []string
}

// MARK: Struct's constructors
func Default() *Cocktail {
	// Define host
	host := os.Getenv("HOST")
	if len(host) == 0 {
		host = "localhost"
	}

	// Define port
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	// Define log
	logger := log.New(os.Stdout, "[Cocktail] ", 0)

	// Define group
	groups := make([]string, 0)

	return &Cocktail{
		Host:    host,
		Port:    port,
		Logger:  logger,
		groups:  groups,
		methods: []string{DELETE, GET, HEAD, PATCH, POST, PUT},
	}
}

// MARK: Struct's public functions
func (c *Cocktail) Run() {
	server := &http.Server{
		Addr:           fmt.Sprintf("%s:%s", c.Host, c.Port),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 5 << 10, // 512kb
		Handler:        c,
	}

	c.Logger.Printf("listening on %s:%s\n", c.Host, c.Port)
	c.Logger.Fatalln(server.ListenAndServe())
}

//func (engine *Engine) RunTLS(addr string, cert string, key string) {
//	if gin_mode == debugCode {
//		fmt.Println("[GIN-debug] Listening and serving HTTPS on " + addr)
//	}
//	if err := http.ListenAndServeTLS(addr, cert, key, engine); err != nil {
//		panic(err)
//	}
//}

func (c *Cocktail) Group(urlGroup string, function func(c *Cocktail)) {
	c.groups = append(c.groups, urlGroup)
	function(c)

	c.groups = c.groups[:len(c.groups)-1]
}

func (c *Cocktail) Delete(urlPath string, handler interface{}) {
	c.addRoute(DELETE, urlPath, handler)
}
func (c *Cocktail) Get(urlPath string, handler interface{}) {
	c.addRoute(GET, urlPath, handler)
}
func (c *Cocktail) Head(urlPath string, handler interface{}) {
	c.addRoute(HEAD, urlPath, handler)
}
func (c *Cocktail) Patch(urlPath string, handler interface{}) {
	c.addRoute(PATCH, urlPath, handler)
}
func (c *Cocktail) Post(urlPath string, handler interface{}) {
	c.addRoute(POST, urlPath, handler)
}
func (c *Cocktail) Put(urlPath string, handler interface{}) {
	c.addRoute(PUT, urlPath, handler)
}

// MARK: Struct's private functions
func (c *Cocktail) addRoute(method string, pattern string, handler interface{}) {
	// Condition validation: If pattern belong to group or not
	if len(c.groups) > 0 {
		groupPattern := ""

		for _, g := range c.groups {
			groupPattern += g
		}
		pattern = groupPattern + pattern
	}

	// Look for existing one before create new
	for _, route := range c.routes {
		if route.Pattern == pattern {
			route.AddHandler(method, handler)
			c.Logger.Printf("%-6s -> %s\n", strings.ToUpper(method), route.Pattern)
			return
		}
	}

	// Create new route
	newRoute := createRoute(pattern)
	newRoute.AddHandler(method, handler)
	c.Logger.Printf("%-6s -> %s\n", strings.ToUpper(method), newRoute.Pattern)

	// Add to collection
	c.routes = append(c.routes, *newRoute)
}

func (c *Cocktail) serveRequest(context *Context) {
	// Condition validation: Validate request method
	isAllowed := false
	for _, allowedMethod := range c.methods {
		if context.Request.Method == allowedMethod {
			isAllowed = true
			break
		}
	}

	if !isAllowed {
		context.RenderError(common.Status405())
		return
	}

	// Condition validation: Find matched route
	for _, route := range c.routes {
		// Match url & extract path params
		ok, pathParams := route.Match(context.Request.Method, context.Request.URL.Path)
		if !ok {
			continue
		}

		context.PathParams = pathParams
		route.InvokeHandler(context)
		return
	}

	// Not Found
	context.RenderError(common.Status404())
}
func (c *Cocktail) serveResource(context *Context) {
	// Condition validation: Only GET is accepted when request a static resources
	if context.Request.Method != GET {
		context.RenderError(common.Status403())
		return
	}
	resourcePath := context.Request.URL.Path[1:]

	// Condition validation: Check if file exist or not
	if !utils.FileExisted(resourcePath) {
		context.RenderError(common.Status404())
		return
	}

	// Open file as read only
	f, err := os.OpenFile(resourcePath, os.O_RDONLY, 0)
	defer f.Close()

	if err != nil {
		context.RenderError(common.Status404())
		return
	}

	// Condition validation: Only serve file, not directory
	fi, _ := f.Stat()
	if fi.IsDir() {
		context.RenderError(common.Status403())
		return
	}

	c.Logger.Printf("serve static: %s", resourcePath)
	http.ServeContent(context.Response, context.Request, resourcePath, fi.ModTime(), f)
}

// MARK: http.Handler's members
func (c *Cocktail) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	request.Method = strings.ToUpper(request.Method)
	request.URL.Path = utils.FormatPath(request.URL.Path)

	// Create context
	context := &Context{Request: request, Response: response}
	defer context.Recovery(c.Logger)

	if len(c.StaticFolder) > 0 && strings.HasPrefix(request.URL.Path, c.StaticFolder) {
		c.serveResource(context)
	} else {
		c.serveRequest(context)
	}
}

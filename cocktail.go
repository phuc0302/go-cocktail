package cocktail

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Cocktail struct {
	Host         string
	Port         string
	StaticFolder string
	Router       Router
	Delegate     CocktailDelegate

	Logger *log.Logger
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

	return &Cocktail{
		Host:   host,
		Port:   port,
		Logger: logger,

		Router: createDefaultRouter(logger),
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
func (c *Cocktail) RunTLS(certFile string, keyFile string) {
	server := &http.Server{
		Addr:           fmt.Sprintf("%s:%s", c.Host, c.Port),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 5 << 10, // 512kb
		Handler:        c,
	}

	c.Logger.Printf("listening on %s:%s\n", c.Host, c.Port)
	c.Logger.Fatalln(server.ListenAndServeTLS(certFile, keyFile))
}

// MARK: Struct's private functions
func (c *Cocktail) serveRequest(context *Context) {
	isAllowed := c.Router.ShouldAllow(context.Method)
	if !isAllowed {
		context.OutputError(Status405())
		return
	}

	// Let delegate decide if the request should be handled or not
	if c.Delegate != nil && !c.Delegate.ShouldServeHTTP(context) {
		context.OutputError(Status404())
		return
	} else if c.Delegate != nil {
		c.Delegate.WillServeHTTP(context)
	}

	isHandled := c.Router.ServeRequest(context)
	if !isHandled {
		context.OutputError(Status404())
	}

	if c.Delegate != nil {
		c.Delegate.DidServeHTTP(context)
	}
}

func (c *Cocktail) serveResource(context *Context, request *http.Request, response http.ResponseWriter) {
	/* Condition validation: Only GET is accepted when request a static resources */
	if request.Method != GET {
		context.OutputError(Status405())
		return
	}
	resourcePath := request.URL.Path[1:]

	/* Condition validation: Check if file exist or not */
	if !FileExisted(resourcePath) {
		context.OutputError(Status404())
		return
	}

	// Open file as read only
	f, err := os.OpenFile(resourcePath, os.O_RDONLY, 0)
	defer f.Close()

	if err != nil {
		context.OutputError(Status404())
		return
	}

	/* Condition validation: Only serve file, not directory */
	fi, _ := f.Stat()
	if fi.IsDir() {
		context.OutputError(Status403())
		return
	}

	http.ServeContent(response, request, resourcePath, fi.ModTime(), f)
}

// MARK: http.Handler's members
func (c *Cocktail) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	request.URL.Path = FormatPath(request.URL.Path)
	request.Method = strings.ToUpper(request.Method)

	// Create context
	context := CreateContext(request, response)
	defer RecoveryRequest(context)

	if len(c.StaticFolder) > 0 && strings.HasPrefix(request.URL.Path, c.StaticFolder) {
		c.serveResource(context, request, response)
	} else {
		c.serveRequest(context)
	}
}

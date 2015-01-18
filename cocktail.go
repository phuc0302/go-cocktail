package cocktail

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

/**
 * Cocktail represents the top level web application.
 */
type Cocktail struct {
	IRouter
	Host   string
	Port   string
	Static string
	Logger *log.Logger
}

// MARK: Struct's constructors
func Classic() *Cocktail {
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
		IRouter: DefaultRouter(logger),
		Host:    host,
		Port:    port,
		Logger:  logger,
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

// MARK: http.Handler's members
func (c *Cocktail) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	request.Method = strings.ToUpper(request.Method)
	defer Recovery(request, response, c.Logger)

	if len(c.Static) > 0 && strings.HasPrefix(request.URL.Path, c.Static) {
		c.serveStatic(request, response)
	} else {
		c.HandleRequest(request, response)
	}
}

// MARK: Struct's private functions
func (c *Cocktail) serveStatic(request *http.Request, response http.ResponseWriter) {
	// Condition validation: Only GET is accepted when request a static resources
	if request.Method != GET {
		WriteError(response, Error403())
		return
	}
	resourcePath := request.URL.Path[1:]

	// Condition validation: Check if file exist or not
	if !FileExist(resourcePath) {
		WriteError(response, Error404())
		return
	}

	// Open file as read only
	f, err := os.OpenFile(resourcePath, os.O_RDONLY, 0)
	defer f.Close()

	if err != nil {
		WriteError(response, Error404())
		return
	}

	// Condition validation: Only serve file, not directory
	fi, _ := f.Stat()
	if fi.IsDir() {
		WriteError(response, Error403())
		return
	}

	c.Logger.Printf("serve static: %s", resourcePath)
	http.ServeContent(response, request, resourcePath, fi.ModTime(), f)
}

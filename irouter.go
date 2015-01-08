package cocktail

import "net/http"

// Router is Martini's de-facto routing interface. Supports HTTP verbs, stacked handlers, and dependency injection.
type IRouter interface {

	// Group adds a group where related routes can be added.
	Group(urlGroup string, function func(router IRouter))

	// CRUD
	Get(urlPath string, handler IHandler)    // Read
	Post(urlPath string, handler IHandler)   // Create
	Patch(urlPath string, handler IHandler)  // Update
	Delete(urlPath string, handler IHandler) // Delete

	// Handle is the entry point for routing. This is used as a martini.Handler
	HandleRequest(request *http.Request, response http.ResponseWriter)
}

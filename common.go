package cocktail

import "mime/multipart"

// Constants
const (
	DELETE = "DELETE"
	GET    = "GET"
	HEAD   = "HEAD"
	PATCH  = "PATCH"
	POST   = "POST"
	PUT    = "PUT"
)

// Valid Http methods
var HTTP_METHODS = [...]string{DELETE, GET, HEAD, PATCH, POST, PUT}

type (
	HandlerFunc func(*Context)
	GroupFunc   func(*Cocktail)

	// Path params from url pattern
	PathParams map[string]string

	// Multipart files from multipart form
	FileParams map[string][]*multipart.FileHeader
)

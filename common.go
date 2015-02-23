package cocktail

import "mime/multipart"

// Constants
const (
	DELETE = "DELETE"
	GET    = "GET"
	PATCH  = "PATCH"
	POST   = "POST"
	PUT    = "PUT"
)

// Valid Http methods
var HTTP_METHODS = [...]string{GET, POST, PATCH, DELETE, PUT}

// Path params from url pattern
type PathParams map[string]string

// Multipart files from multipart form
type FileParams map[string][]*multipart.FileHeader

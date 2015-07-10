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

type (
	// Path params from url pattern
	PathParams map[string]string

	// Multipart files from multipart form
	FileParams map[string][]*multipart.FileHeader
)

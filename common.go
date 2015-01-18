package cocktail

import "mime/multipart"

// Constants
const DELETE = "DELETE"
const GET = "GET"
const PATCH = "PATCH"
const POST = "POST"
const PUT = "PUT"

// Valid Http methods
var HTTP_METHODS = [...]string{GET, POST, PATCH, DELETE, PUT}

// Path params from url pattern
type PathParams map[string]string

// Multipart files from multipart form
type FileParams map[string][]*multipart.FileHeader

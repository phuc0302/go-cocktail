package cocktail

import "net/http"

type HttpStatus struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

// MARK: Struct's constructors
func Status200() *HttpStatus {
	return genericError(http.StatusOK)
}
func Status201() *HttpStatus {
	return genericError(http.StatusCreated)
}

func Error400() *HttpStatus {
	return genericError(http.StatusBadRequest)
}
func Error401() *HttpStatus {
	return genericError(http.StatusUnauthorized)
}
func Error402() *HttpStatus {
	return genericError(http.StatusPaymentRequired)
}
func Error403() *HttpStatus {
	return genericError(http.StatusForbidden)
}
func Error404() *HttpStatus {
	return genericError(http.StatusNotFound)
}
func Error405() *HttpStatus {
	return genericError(http.StatusMethodNotAllowed)
}
func Error406() *HttpStatus {
	return genericError(http.StatusNotAcceptable)
}
func Error407() *HttpStatus {
	return genericError(http.StatusProxyAuthRequired)
}
func Error408() *HttpStatus {
	return genericError(http.StatusRequestTimeout)
}
func Error409() *HttpStatus {
	return genericError(http.StatusConflict)
}
func Error410() *HttpStatus {
	return genericError(http.StatusGone)
}
func Error411() *HttpStatus {
	return genericError(http.StatusLengthRequired)
}
func Error412() *HttpStatus {
	return genericError(http.StatusPreconditionFailed)
}
func Error413() *HttpStatus {
	return genericError(http.StatusRequestEntityTooLarge)
}
func Error414() *HttpStatus {
	return genericError(http.StatusRequestURITooLong)
}
func Error415() *HttpStatus {
	return genericError(http.StatusUnsupportedMediaType)
}
func Error416() *HttpStatus {
	return genericError(http.StatusRequestedRangeNotSatisfiable)
}
func Error417() *HttpStatus {
	return genericError(http.StatusExpectationFailed)
}
func Error418() *HttpStatus {
	return genericError(http.StatusTeapot)
}
func Error422() *HttpStatus {
	return specificError(422, "Unprocessable Entity")
}
func Error423() *HttpStatus {
	return specificError(423, "Locked")
}
func Error424() *HttpStatus {
	return specificError(424, "Failed Dependency")
}
func Error425() *HttpStatus {
	return specificError(425, "Unordered Collection")
}
func Error426() *HttpStatus {
	return specificError(426, "Upgrade Required")
}
func Error428() *HttpStatus {
	return specificError(428, "Precondition Required")
}
func Error429() *HttpStatus {
	return specificError(429, "Too Many Requests")
}
func Error431() *HttpStatus {
	return specificError(431, "Request Header Fields Too Large")
}

func Error500() *HttpStatus {
	return genericError(http.StatusInternalServerError)
}
func Error501() *HttpStatus {
	return genericError(http.StatusNotImplemented)
}
func Error502() *HttpStatus {
	return genericError(http.StatusBadGateway)
}
func Error503() *HttpStatus {
	return genericError(http.StatusServiceUnavailable)
}
func Error504() *HttpStatus {
	return genericError(http.StatusGatewayTimeout)
}
func Error505() *HttpStatus {
	return genericError(http.StatusHTTPVersionNotSupported)
}
func Error506() *HttpStatus {
	return specificError(506, "Variant Also Negotiates")
}
func Error507() *HttpStatus {
	return specificError(507, "Insufficient Storage")
}
func Error508() *HttpStatus {
	return specificError(508, "Loop Detected")
}
func Error511() *HttpStatus {
	return specificError(511, "Network Authentication Required")
}

// MARK: Struct's private constructors
func genericError(statusCode int) *HttpStatus {
	return &HttpStatus{
		Status: statusCode,
		Title:  http.StatusText(statusCode),
		Detail: http.StatusText(statusCode),
	}
}
func specificError(statusCode int, title string) *HttpStatus {
	return &HttpStatus{
		Status: statusCode,
		Title:  title,
		Detail: title,
	}
}

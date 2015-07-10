package common

import "net/http"

type Status struct {
	Status int         `json:"status"`
	Title  string      `json:"title"`
	Detail interface{} `json:"detail"`
}

// MARK: Struct's constructors
func Status200() *Status {
	return genericError(http.StatusOK)
}
func Status201() *Status {
	return genericError(http.StatusCreated)
}

func Status400() *Status {
	return genericError(http.StatusBadRequest)
}
func Status401() *Status {
	return genericError(http.StatusUnauthorized)
}
func Status402() *Status {
	return genericError(http.StatusPaymentRequired)
}
func Status403() *Status {
	return genericError(http.StatusForbidden)
}
func Status404() *Status {
	return genericError(http.StatusNotFound)
}
func Status405() *Status {
	return genericError(http.StatusMethodNotAllowed)
}
func Status406() *Status {
	return genericError(http.StatusNotAcceptable)
}
func Status407() *Status {
	return genericError(http.StatusProxyAuthRequired)
}
func Status408() *Status {
	return genericError(http.StatusRequestTimeout)
}
func Status409() *Status {
	return genericError(http.StatusConflict)
}
func Status410() *Status {
	return genericError(http.StatusGone)
}
func Status411() *Status {
	return genericError(http.StatusLengthRequired)
}
func Status412() *Status {
	return genericError(http.StatusPreconditionFailed)
}
func Status413() *Status {
	return genericError(http.StatusRequestEntityTooLarge)
}
func Status414() *Status {
	return genericError(http.StatusRequestURITooLong)
}
func Status415() *Status {
	return genericError(http.StatusUnsupportedMediaType)
}
func Status416() *Status {
	return genericError(http.StatusRequestedRangeNotSatisfiable)
}
func Status417() *Status {
	return genericError(http.StatusExpectationFailed)
}
func Status422() *Status {
	return specificError(422, "Unprocessable Entity")
}
func Status423() *Status {
	return specificError(423, "Locked")
}
func Status424() *Status {
	return specificError(424, "Failed Dependency")
}
func Status425() *Status {
	return specificError(425, "Unordered Collection")
}
func UpgradeRequired() *Status {
	return specificError(426, "Upgrade Required")
}
func Status428() *Status {
	return specificError(428, "Precondition Required")
}
func Status429() *Status {
	return specificError(429, "Too Many Requests")
}
func Status431() *Status {
	return specificError(431, "Request Header Fields Too Large")
}

func Status500() *Status {
	return genericError(http.StatusInternalServerError)
}
func Status501() *Status {
	return genericError(http.StatusNotImplemented)
}
func Status502() *Status {
	return genericError(http.StatusBadGateway)
}
func Status503() *Status {
	return genericError(http.StatusServiceUnavailable)
}
func Status504() *Status {
	return genericError(http.StatusGatewayTimeout)
}
func Status505() *Status {
	return genericError(http.StatusHTTPVersionNotSupported)
}
func Status506() *Status {
	return specificError(506, "Variant Also Negotiates")
}
func Status507() *Status {
	return specificError(507, "Insufficient Storage")
}
func Status508() *Status {
	return specificError(508, "Loop Detected")
}
func Status511() *Status {
	return specificError(511, "Network Authentication Required")
}

// MARK: Struct's private constructors
func genericError(statusCode int) *Status {
	return &Status{
		Status: statusCode,
		Title:  http.StatusText(statusCode),
		Detail: http.StatusText(statusCode),
	}
}
func specificError(statusCode int, title string) *Status {
	return &Status{
		Status: statusCode,
		Title:  title,
		Detail: title,
	}
}

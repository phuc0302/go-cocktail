package common

import "net/http"

type Status struct {
	Status int         `json:"status,omitempty"`
	Title  string      `json:"title,omitempty"`
	Detail interface{} `json:"detail,omitempty"`
}

// MARK: Struct's constructors
func Status200() *Status {
	return genericStatus(http.StatusOK)
}
func Status201() *Status {
	return genericStatus(http.StatusCreated)
}

func Status400() *Status {
	return genericStatus(http.StatusBadRequest)
}
func Status401() *Status {
	return genericStatus(http.StatusUnauthorized)
}
func Status402() *Status {
	return genericStatus(http.StatusPaymentRequired)
}
func Status403() *Status {
	return genericStatus(http.StatusForbidden)
}
func Status404() *Status {
	return genericStatus(http.StatusNotFound)
}
func Status405() *Status {
	return genericStatus(http.StatusMethodNotAllowed)
}
func Status406() *Status {
	return genericStatus(http.StatusNotAcceptable)
}
func Status407() *Status {
	return genericStatus(http.StatusProxyAuthRequired)
}
func Status408() *Status {
	return genericStatus(http.StatusRequestTimeout)
}
func Status409() *Status {
	return genericStatus(http.StatusConflict)
}
func Status410() *Status {
	return genericStatus(http.StatusGone)
}
func Status411() *Status {
	return genericStatus(http.StatusLengthRequired)
}
func Status412() *Status {
	return genericStatus(http.StatusPreconditionFailed)
}
func Status413() *Status {
	return genericStatus(http.StatusRequestEntityTooLarge)
}
func Status414() *Status {
	return genericStatus(http.StatusRequestURITooLong)
}
func Status415() *Status {
	return genericStatus(http.StatusUnsupportedMediaType)
}
func Status416() *Status {
	return genericStatus(http.StatusRequestedRangeNotSatisfiable)
}
func Status417() *Status {
	return genericStatus(http.StatusExpectationFailed)
}
func Status422() *Status {
	return specificStatus(422, "Unprocessable Entity")
}
func Status423() *Status {
	return specificStatus(423, "Locked")
}
func Status424() *Status {
	return specificStatus(424, "Failed Dependency")
}
func Status425() *Status {
	return specificStatus(425, "Unordered Collection")
}
func UpgradeRequired() *Status {
	return specificStatus(426, "Upgrade Required")
}
func Status428() *Status {
	return specificStatus(428, "Precondition Required")
}
func Status429() *Status {
	return specificStatus(429, "Too Many Requests")
}
func Status431() *Status {
	return specificStatus(431, "Request Header Fields Too Large")
}

func Status500() *Status {
	return genericStatus(http.StatusInternalServerError)
}
func Status501() *Status {
	return genericStatus(http.StatusNotImplemented)
}
func Status502() *Status {
	return genericStatus(http.StatusBadGateway)
}
func Status503() *Status {
	return genericStatus(http.StatusServiceUnavailable)
}
func Status504() *Status {
	return genericStatus(http.StatusGatewayTimeout)
}
func Status505() *Status {
	return genericStatus(http.StatusHTTPVersionNotSupported)
}
func Status506() *Status {
	return specificStatus(506, "Variant Also Negotiates")
}
func Status507() *Status {
	return specificStatus(507, "Insufficient Storage")
}
func Status508() *Status {
	return specificStatus(508, "Loop Detected")
}
func Status511() *Status {
	return specificStatus(511, "Network Authentication Required")
}

// MARK: Struct's private constructors
func genericStatus(statusCode int) *Status {
	return &Status{
		Status: statusCode,
		Title:  http.StatusText(statusCode),
		Detail: http.StatusText(statusCode),
	}
}
func specificStatus(statusCode int, title string) *Status {
	return &Status{
		Status: statusCode,
		Title:  title,
		Detail: title,
	}
}

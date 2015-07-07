package cocktail

import "net/http"

type Status struct {
	Status int         `json:"status"`
	Title  string      `json:"title"`
	Detail interface{} `json:"detail"`
}

// MARK: Struct's constructors
func OK() *Status {
	return genericError(http.StatusOK)
}
func Created() *Status {
	return genericError(http.StatusCreated)
}

func BadRequest() *Status {
	return genericError(http.StatusBadRequest)
}
func Unauthorized() *Status {
	return genericError(http.StatusUnauthorized)
}
func PaymentRequired() *Status {
	return genericError(http.StatusPaymentRequired)
}
func Forbidden() *Status {
	return genericError(http.StatusForbidden)
}
func NotFound() *Status {
	return genericError(http.StatusNotFound)
}
func MethodNotAllowed() *Status {
	return genericError(http.StatusMethodNotAllowed)
}
func NotAcceptable() *Status {
	return genericError(http.StatusNotAcceptable)
}
func ProxyAuthRequired() *Status {
	return genericError(http.StatusProxyAuthRequired)
}
func RequestTimeout() *Status {
	return genericError(http.StatusRequestTimeout)
}
func Conflict() *Status {
	return genericError(http.StatusConflict)
}
func Gone() *Status {
	return genericError(http.StatusGone)
}
func LengthRequired() *Status {
	return genericError(http.StatusLengthRequired)
}
func PreconditionFailed() *Status {
	return genericError(http.StatusPreconditionFailed)
}
func RequestEntityTooLarge() *Status {
	return genericError(http.StatusRequestEntityTooLarge)
}
func RequestURITooLong() *Status {
	return genericError(http.StatusRequestURITooLong)
}
func UnsupportedMediaType() *Status {
	return genericError(http.StatusUnsupportedMediaType)
}
func RequestedRangeNotSatisfiable() *Status {
	return genericError(http.StatusRequestedRangeNotSatisfiable)
}
func ExpectationFailed() *Status {
	return genericError(http.StatusExpectationFailed)
}
func Teapot() *Status {
	return genericError(http.StatusTeapot)
}
func UnprocessableEntity() *Status {
	return specificError(422, "Unprocessable Entity")
}
func Locked() *Status {
	return specificError(423, "Locked")
}
func FailedDependency() *Status {
	return specificError(424, "Failed Dependency")
}
func UnorderedCollection() *Status {
	return specificError(425, "Unordered Collection")
}
func UpgradeRequired() *Status {
	return specificError(426, "Upgrade Required")
}
func PreconditionRequired() *Status {
	return specificError(428, "Precondition Required")
}
func TooManyRequests() *Status {
	return specificError(429, "Too Many Requests")
}
func RequestHeaderFieldsTooLarge() *Status {
	return specificError(431, "Request Header Fields Too Large")
}

func InternalServerError() *Status {
	return genericError(http.StatusInternalServerError)
}
func NotImplemented() *Status {
	return genericError(http.StatusNotImplemented)
}
func BadGateway() *Status {
	return genericError(http.StatusBadGateway)
}
func ServiceUnavailable() *Status {
	return genericError(http.StatusServiceUnavailable)
}
func GatewayTimeout() *Status {
	return genericError(http.StatusGatewayTimeout)
}
func HTTPVersionNotSupported() *Status {
	return genericError(http.StatusHTTPVersionNotSupported)
}
func VariantAlsoNegotiates() *Status {
	return specificError(506, "Variant Also Negotiates")
}
func InsufficientStorage() *Status {
	return specificError(507, "Insufficient Storage")
}
func LoopDetected() *Status {
	return specificError(508, "Loop Detected")
}
func NetworkAuthenticationRequired() *Status {
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

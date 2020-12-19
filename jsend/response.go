package jsend

import "net/http"

func OK(data interface{}) *ResponseMessage {
	return New(http.StatusOK).AddData(data)
}

func Created(data interface{}) *ResponseMessage {
	return New(http.StatusCreated).AddData(data)
}

func Accepted(data interface{}) *ResponseMessage {
	return New(http.StatusAccepted).AddData(data)
}

func NoContent() *ResponseMessage {
	return New(http.StatusNoContent)
}

func NonAuthoritativeInfo(data interface{}) *ResponseMessage {
	return New(http.StatusNonAuthoritativeInfo).AddData(data)
}
func ResetContent(data interface{}) *ResponseMessage {
	return New(http.StatusResetContent).AddData(data)
}
func PartialContent(data interface{}) *ResponseMessage {
	return New(http.StatusPartialContent).AddData(data)
}
func MultiStatus(data interface{}) *ResponseMessage {
	return New(http.StatusMultiStatus).AddData(data)
}
func AlreadyReported(data interface{}) *ResponseMessage {
	return New(http.StatusAlreadyReported).AddData(data)
}
func IMUsed(data interface{}) *ResponseMessage {
	return New(http.StatusIMUsed).AddData(data)
}

func BadRequest(data interface{}) *ResponseMessage {
	return New(http.StatusBadRequest).AddData(data)
}

func Unauthorized(data interface{}) *ResponseMessage {
	return New(http.StatusUnauthorized).AddData(data)
}

func Forbidden(data interface{}) *ResponseMessage {
	return New(http.StatusForbidden).AddData(data)
}

func NotFound(data interface{}) *ResponseMessage {
	return New(http.StatusNotFound).AddData(data)
}

func MethodNotAllowed(data interface{}) *ResponseMessage {
	return New(http.StatusMethodNotAllowed).AddData(data)
}

func NotAcceptable(data interface{}) *ResponseMessage {
	return New(http.StatusNotAcceptable).AddData(data)
}

func ProxyAuthRequired(data interface{}) *ResponseMessage {
	return New(http.StatusProxyAuthRequired).AddData(data)
}
func RequestTimeout(data interface{}) *ResponseMessage {
	return New(http.StatusRequestTimeout).AddData(data)
}
func Conflict(data interface{}) *ResponseMessage {
	return New(http.StatusConflict).AddData(data)
}
func Gone(data interface{}) *ResponseMessage {
	return New(http.StatusGone).AddData(data)
}
func LengthRequired(data interface{}) *ResponseMessage {
	return New(http.StatusLengthRequired).AddData(data)
}
func PreconditionFailed(data interface{}) *ResponseMessage {
	return New(http.StatusPreconditionFailed).AddData(data)
}
func RequestEntityTooLarge(data interface{}) *ResponseMessage {
	return New(http.StatusRequestEntityTooLarge).AddData(data)
}
func RequestURITooLong(data interface{}) *ResponseMessage {
	return New(http.StatusRequestURITooLong).AddData(data)
}
func UnsupportedMediaType(data interface{}) *ResponseMessage {
	return New(http.StatusUnsupportedMediaType).AddData(data)
}
func RequestedRangeNotSatisfiable(data interface{}) *ResponseMessage {
	return New(http.StatusRequestedRangeNotSatisfiable).AddData(data)
}
func ExpectationFailed(data interface{}) *ResponseMessage {
	return New(http.StatusExpectationFailed).AddData(data)
}
func Teapot(data interface{}) *ResponseMessage {
	return New(http.StatusTeapot).AddData(data)
}
func MisdirectedRequest(data interface{}) *ResponseMessage {
	return New(http.StatusMisdirectedRequest).AddData(data)
}
func UnprocessableEntity(data interface{}) *ResponseMessage {
	return New(http.StatusUnprocessableEntity).AddData(data)
}
func Locked(data interface{}) *ResponseMessage {
	return New(http.StatusLocked).AddData(data)
}
func FailedDependency(data interface{}) *ResponseMessage {
	return New(http.StatusFailedDependency).AddData(data)
}
func TooEarly(data interface{}) *ResponseMessage {
	return New(http.StatusTooEarly).AddData(data)
}
func UpgradeRequired(data interface{}) *ResponseMessage {
	return New(http.StatusUpgradeRequired).AddData(data)
}
func PreconditionRequired(data interface{}) *ResponseMessage {
	return New(http.StatusPreconditionRequired).AddData(data)
}
func TooManyRequests(data interface{}) *ResponseMessage {
	return New(http.StatusTooManyRequests).AddData(data)
}
func RequestHeaderFieldsTooLarge(data interface{}) *ResponseMessage {
	return New(http.StatusRequestHeaderFieldsTooLarge).AddData(data)
}
func UnavailableForLegalReasons(data interface{}) *ResponseMessage {
	return New(http.StatusUnavailableForLegalReasons).AddData(data)
}

func InternalServerError(message string) *ResponseMessage {
	return New(http.StatusInternalServerError).AddMessage(message)
}

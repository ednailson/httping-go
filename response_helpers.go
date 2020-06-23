package httping

import "net/http"

func OK(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusOK).AddData(data)
}

func Created(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusCreated).AddData(data)
}

func Accepted(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusAccepted).AddData(data)
}

func NoContent() *ResponseMessage {
	return NewResponse(http.StatusNoContent)
}

func NonAuthoritativeInfo(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusNonAuthoritativeInfo).AddData(data)
}
func ResetContent(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusResetContent).AddData(data)
}
func PartialContent(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusPartialContent).AddData(data)
}
func MultiStatus(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusMultiStatus).AddData(data)
}
func AlreadyReported(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusAlreadyReported).AddData(data)
}
func IMUsed(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusIMUsed).AddData(data)
}

func BadRequest(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusBadRequest).AddData(data)
}

func Unauthorized(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusUnauthorized).AddData(data)
}

func Forbidden(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusForbidden).AddData(data)
}

func NotFound(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusNotFound).AddData(data)
}

func MethodNotAllowed(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusMethodNotAllowed).AddData(data)
}

func NotAcceptable(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusNotAcceptable).AddData(data)
}

func ProxyAuthRequired(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusProxyAuthRequired).AddData(data)
}
func RequestTimeout(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusRequestTimeout).AddData(data)
}
func Conflict(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusConflict).AddData(data)
}
func Gone(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusGone).AddData(data)
}
func LengthRequired(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusLengthRequired).AddData(data)
}
func PreconditionFailed(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusPreconditionFailed).AddData(data)
}
func RequestEntityTooLarge(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusRequestEntityTooLarge).AddData(data)
}
func RequestURITooLong(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusRequestURITooLong).AddData(data)
}
func UnsupportedMediaType(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusUnsupportedMediaType).AddData(data)
}
func RequestedRangeNotSatisfiable(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusRequestedRangeNotSatisfiable).AddData(data)
}
func ExpectationFailed(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusExpectationFailed).AddData(data)
}
func Teapot(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusTeapot).AddData(data)
}
func MisdirectedRequest(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusMisdirectedRequest).AddData(data)
}
func UnprocessableEntity(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusUnprocessableEntity).AddData(data)
}
func Locked(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusLocked).AddData(data)
}
func FailedDependency(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusFailedDependency).AddData(data)
}
func TooEarly(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusTooEarly).AddData(data)
}
func UpgradeRequired(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusUpgradeRequired).AddData(data)
}
func PreconditionRequired(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusPreconditionRequired).AddData(data)
}
func TooManyRequests(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusTooManyRequests).AddData(data)
}
func RequestHeaderFieldsTooLarge(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusRequestHeaderFieldsTooLarge).AddData(data)
}
func UnavailableForLegalReasons(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusUnavailableForLegalReasons).AddData(data)
}

func InternalServerError(message string) *ResponseMessage {
	return NewResponse(http.StatusInternalServerError).AddMessage(message)
}

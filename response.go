package httping

import "net/http"

type Response struct {
	data       interface{}
	headers    map[string][]string
	statusCode int
	cookies    []*http.Cookie
}

func NewResponse(statusCode int) *Response {
	return &Response{headers: make(map[string][]string), statusCode: statusCode}
}

func (r *Response) SetData(data interface{}) *Response {
	r.data = data
	return r
}

func (r *Response) AddHeader(key, value string) *Response {
	r.headers[key] = append(r.headers[key], value)
	return r
}

func (r *Response) AddCookie(cookie *http.Cookie) *Response {
	r.cookies = append(r.cookies, cookie)
	return r
}

func (r *Response) SetCookies(cookies []*http.Cookie) *Response {
	r.cookies = cookies
	return r
}

func (r *Response) StatusCode() int {
	return r.statusCode
}

func (r *Response) Cookies() []*http.Cookie {
	return r.cookies
}

func (r *Response) Headers() map[string][]string {
	return r.headers
}

func (r *Response) Response() interface{} {
	return r.data
}

func OK(data interface{}) *Response {
	return NewResponse(http.StatusOK).SetData(data)
}

func Created(data interface{}) *Response {
	return NewResponse(http.StatusCreated).SetData(data)
}

func Accepted(data interface{}) *Response {
	return NewResponse(http.StatusAccepted).SetData(data)
}

func NoContent() *Response {
	return NewResponse(http.StatusNoContent)
}

func NonAuthoritativeInfo(data interface{}) *Response {
	return NewResponse(http.StatusNonAuthoritativeInfo).SetData(data)
}
func ResetContent(data interface{}) *Response {
	return NewResponse(http.StatusResetContent).SetData(data)
}
func PartialContent(data interface{}) *Response {
	return NewResponse(http.StatusPartialContent).SetData(data)
}
func MultiStatus(data interface{}) *Response {
	return NewResponse(http.StatusMultiStatus).SetData(data)
}
func AlreadyReported(data interface{}) *Response {
	return NewResponse(http.StatusAlreadyReported).SetData(data)
}
func IMUsed(data interface{}) *Response {
	return NewResponse(http.StatusIMUsed).SetData(data)
}

func BadRequest(data interface{}) *Response {
	return NewResponse(http.StatusBadRequest).SetData(data)
}

func Unauthorized(data interface{}) *Response {
	return NewResponse(http.StatusUnauthorized).SetData(data)
}

func Forbidden(data interface{}) *Response {
	return NewResponse(http.StatusForbidden).SetData(data)
}

func NotFound(data interface{}) *Response {
	return NewResponse(http.StatusNotFound).SetData(data)
}

func MethodNotAllowed(data interface{}) *Response {
	return NewResponse(http.StatusMethodNotAllowed).SetData(data)
}

func NotAcceptable(data interface{}) *Response {
	return NewResponse(http.StatusNotAcceptable).SetData(data)
}

func ProxyAuthRequired(data interface{}) *Response {
	return NewResponse(http.StatusProxyAuthRequired).SetData(data)
}
func RequestTimeout(data interface{}) *Response {
	return NewResponse(http.StatusRequestTimeout).SetData(data)
}
func Conflict(data interface{}) *Response {
	return NewResponse(http.StatusConflict).SetData(data)
}
func Gone(data interface{}) *Response {
	return NewResponse(http.StatusGone).SetData(data)
}
func LengthRequired(data interface{}) *Response {
	return NewResponse(http.StatusLengthRequired).SetData(data)
}
func PreconditionFailed(data interface{}) *Response {
	return NewResponse(http.StatusPreconditionFailed).SetData(data)
}
func RequestEntityTooLarge(data interface{}) *Response {
	return NewResponse(http.StatusRequestEntityTooLarge).SetData(data)
}
func RequestURITooLong(data interface{}) *Response {
	return NewResponse(http.StatusRequestURITooLong).SetData(data)
}
func UnsupportedMediaType(data interface{}) *Response {
	return NewResponse(http.StatusUnsupportedMediaType).SetData(data)
}
func RequestedRangeNotSatisfiable(data interface{}) *Response {
	return NewResponse(http.StatusRequestedRangeNotSatisfiable).SetData(data)
}
func ExpectationFailed(data interface{}) *Response {
	return NewResponse(http.StatusExpectationFailed).SetData(data)
}
func Teapot(data interface{}) *Response {
	return NewResponse(http.StatusTeapot).SetData(data)
}
func MisdirectedRequest(data interface{}) *Response {
	return NewResponse(http.StatusMisdirectedRequest).SetData(data)
}
func UnprocessableEntity(data interface{}) *Response {
	return NewResponse(http.StatusUnprocessableEntity).SetData(data)
}
func Locked(data interface{}) *Response {
	return NewResponse(http.StatusLocked).SetData(data)
}
func FailedDependency(data interface{}) *Response {
	return NewResponse(http.StatusFailedDependency).SetData(data)
}
func TooEarly(data interface{}) *Response {
	return NewResponse(http.StatusTooEarly).SetData(data)
}
func UpgradeRequired(data interface{}) *Response {
	return NewResponse(http.StatusUpgradeRequired).SetData(data)
}
func PreconditionRequired(data interface{}) *Response {
	return NewResponse(http.StatusPreconditionRequired).SetData(data)
}
func TooManyRequests(data interface{}) *Response {
	return NewResponse(http.StatusTooManyRequests).SetData(data)
}
func RequestHeaderFieldsTooLarge(data interface{}) *Response {
	return NewResponse(http.StatusRequestHeaderFieldsTooLarge).SetData(data)
}
func UnavailableForLegalReasons(data interface{}) *Response {
	return NewResponse(http.StatusUnavailableForLegalReasons).SetData(data)
}

func InternalServerError(data interface{}) *Response {
	return NewResponse(http.StatusInternalServerError).SetData(data)
}

package httping

import "net/http"

type ResponseMessage struct {
	Status     ResponseStatus `json:"status,omitempty"`
	Data       interface{}    `json:"data,omitempty"`
	Message    string         `json:"message,omitempty"`
	Code       string         `json:"code,omitempty"`
	headers    map[string][]string
	statusCode int
	cookies    []*http.Cookie
}

type ResponseStatus string

const (
	StatusSuccess ResponseStatus = "success"
	StatusError   ResponseStatus = "error"
	StatusFail    ResponseStatus = "fail"
)

// JSend helper. It follows all rules of JSend.
func NewResponse(statusCode int) *ResponseMessage {
	switch {
	case statusCode >= http.StatusInternalServerError:
		return &ResponseMessage{Status: StatusError, headers: make(map[string][]string), statusCode: statusCode}
	case statusCode >= http.StatusBadRequest && statusCode < http.StatusInternalServerError:
		return &ResponseMessage{Status: StatusFail, headers: make(map[string][]string), statusCode: statusCode}
	case statusCode == http.StatusNoContent:
		return &ResponseMessage{statusCode: statusCode}
	default:
		return &ResponseMessage{Status: StatusSuccess, headers: make(map[string][]string), statusCode: statusCode}
	}
}

func (r *ResponseMessage) AddData(data interface{}) *ResponseMessage {
	r.Data = data
	return r
}

func (r *ResponseMessage) AddMessage(message string) *ResponseMessage {
	if r.Status != StatusError {
		return r
	}
	r.Message = message
	return r
}

func (r *ResponseMessage) AddCode(code string) *ResponseMessage {
	if r.Status != StatusError {
		return r
	}
	r.Code = code
	return r
}

func (r *ResponseMessage) AddHeader(key, value string) *ResponseMessage {
	r.headers[key] = append(r.headers[key], value)
	return r
}

func (r *ResponseMessage) AddCookie(cookie *http.Cookie) *ResponseMessage {
	r.cookies = append(r.cookies, cookie)
	return r
}

func (r *ResponseMessage) SetCookies(cookies []*http.Cookie) *ResponseMessage {
	r.cookies = cookies
	return r
}

func (r *ResponseMessage) StatusCode() int {
	return r.statusCode
}

func (r *ResponseMessage) Cookies() []*http.Cookie {
	return r.cookies
}

func (r *ResponseMessage) Headers() map[string][]string {
	return r.headers
}

func (r *ResponseMessage) Response() interface{} {
	return r
}

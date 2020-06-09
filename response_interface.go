package httping

import "net/http"

type IResponse interface {
	Headers() map[string][]string
	Cookies() []*http.Cookie
	Response() interface{}
	StatusCode() int
}

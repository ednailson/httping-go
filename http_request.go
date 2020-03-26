package httping

import (
	"net/http"
	"net/url"
)

type HttpRequest struct {
	Body    []byte
	Query   url.Values
	Params  map[string]string
	Headers map[string][]string
	Cookies []*http.Cookie
}

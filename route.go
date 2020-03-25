package httping

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

type HandlerFunc func(request HttpRequest) (response *ResponseMessage)
type MiddlewareFunc func(request HttpRequest) (response *ResponseMessage, success bool)

type Route struct {
	route      *gin.RouterGroup
	middleware MiddlewareFunc
}

func (r *Route) AddMethod(method string, handler HandlerFunc) {
	method = strings.ToUpper(method)
	switch method {
	case http.MethodGet:
		r.route.GET("", r.getHandleFunc(handler))
	case http.MethodPost:
		r.route.POST("", r.getHandleFunc(handler))
	case http.MethodPut:
		r.route.PUT("", r.getHandleFunc(handler))
	case http.MethodDelete:
		r.route.DELETE("", r.getHandleFunc(handler))
	case http.MethodPatch:
		r.route.PATCH("", r.getHandleFunc(handler))
	case http.MethodHead:
		r.route.HEAD("", r.getHandleFunc(handler))
	case http.MethodOptions:
		r.route.OPTIONS("", r.getHandleFunc(handler))
	}
}

func (r *Route) POST(handler HandlerFunc) {
	r.AddMethod(http.MethodPost, handler)
}

func (r *Route) GET(handler HandlerFunc) {
	r.AddMethod(http.MethodGet, handler)
}

func (r *Route) PUT(handler HandlerFunc) {
	r.AddMethod(http.MethodPut, handler)
}

func (r *Route) DELETE(handler HandlerFunc) {
	r.AddMethod(http.MethodDelete, handler)
}

func (r *Route) PATCH(handler HandlerFunc) {
	r.AddMethod(http.MethodPatch, handler)
}

func (r *Route) HEAD(handler HandlerFunc) {
	r.AddMethod(http.MethodHead, handler)
}

func (r *Route) OPTIONS(handler HandlerFunc) {
	r.AddMethod(http.MethodOptions, handler)
}

func (r *Route) getHandleFunc(handle HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		params := map[string]string{}
		for _, v := range c.Params {
			params[v.Key] = v.Value
		}
		query := c.Request.URL.Query()
		headers := map[string][]string{}
		for k, v := range c.Request.Header {
			headers[k] = v
		}
		if r.middleware != nil {
			message, ok := r.middleware(HttpRequest{
				Body:    body,
				Query:   query,
				Params:  params,
				Headers: headers,
			})
			if !ok {
				if message != nil {
					c.JSON(message.statusCode, message)
					return
				}
				c.JSON(http.StatusOK, nil)
				return
			}
		}
		message := handle(HttpRequest{
			Body:    body,
			Query:   query,
			Params:  params,
			Headers: headers,
		})
		if message != nil {
			for k, v := range message.headers {
				for _, h := range v {
					c.Writer.Header().Add(k, h)
				}
			}
			c.JSON(message.statusCode, message)
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}

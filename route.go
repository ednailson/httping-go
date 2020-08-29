package httping

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

type HandlerFunc func(request HttpRequest) (response IResponse)

type route struct {
	route      *gin.RouterGroup
	middleware []HandlerFunc
}

// Add a HandlerFunc that it will run every time that your server receives a request in the route with the method set up.
func (r *route) AddMethod(method string, handler HandlerFunc) {
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

func (r *route) POST(handler HandlerFunc) {
	r.AddMethod(http.MethodPost, handler)
}

func (r *route) GET(handler HandlerFunc) {
	r.AddMethod(http.MethodGet, handler)
}

func (r *route) PUT(handler HandlerFunc) {
	r.AddMethod(http.MethodPut, handler)
}

func (r *route) DELETE(handler HandlerFunc) {
	r.AddMethod(http.MethodDelete, handler)
}

func (r *route) PATCH(handler HandlerFunc) {
	r.AddMethod(http.MethodPatch, handler)
}

func (r *route) HEAD(handler HandlerFunc) {
	r.AddMethod(http.MethodHead, handler)
}

func (r *route) OPTIONS(handler HandlerFunc) {
	r.AddMethod(http.MethodOptions, handler)
}

func (r *route) SetMiddleware(middleware []HandlerFunc) IRoute {
	r.middleware = middleware
	return r
}

func (r *route) AddMiddleware(middleware HandlerFunc) IRoute {
	r.middleware = append(r.middleware, middleware)
	return r
}

func (r *route) getRoute() *route {
	return r
}

func (r *route) getHandleFunc(handle HandlerFunc) func(c *gin.Context) {
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
		if r.middleware != nil && len(r.middleware) > 0 {
			for _, middleware := range r.middleware {
				if middleware != nil {
					message := middleware(HttpRequest{
						Body:    body,
						Query:   query,
						Params:  params,
						Headers: headers,
						Cookies: c.Request.Cookies(),
					})
					if message != nil {
						c.JSON(message.StatusCode(), message.Response())
						return
					}
				}
			}
		}
		message := handle(HttpRequest{
			Body:    body,
			Query:   query,
			Params:  params,
			Headers: headers,
			Cookies: c.Request.Cookies(),
		})
		if message != nil {
			for k, v := range message.Headers() {
				for _, h := range v {
					c.Writer.Header().Add(k, h)
				}
			}
			for _, v := range message.Cookies() {
				c.SetCookie(v.Name, v.Value, v.MaxAge, v.Path, v.Domain, v.Secure, v.HttpOnly)
			}
			c.JSON(message.StatusCode(), message.Response())
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}

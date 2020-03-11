package httping

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

type HandlerFunc func(request HttpRequest) (response *ResponseMessage)

type Route struct {
	route *gin.RouterGroup
}

func (gp *Route) AddMethod(method string, handler HandlerFunc) {
	method = strings.ToUpper(method)
	switch method {
	case http.MethodGet:
		gp.route.GET("", getHandleFunc(handler))
	case http.MethodPost:
		gp.route.POST("", getHandleFunc(handler))
	case http.MethodPut:
		gp.route.PUT("", getHandleFunc(handler))
	case http.MethodDelete:
		gp.route.DELETE("", getHandleFunc(handler))
	case http.MethodPatch:
		gp.route.PATCH("", getHandleFunc(handler))
	case http.MethodHead:
		gp.route.HEAD("", getHandleFunc(handler))
	case http.MethodOptions:
		gp.route.OPTIONS("", getHandleFunc(handler))
	}
}

func getHandleFunc(handle HandlerFunc) func(c *gin.Context) {
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
		message := handle(HttpRequest{
			Body:    body,
			Query:   query,
			Params:  params,
			Headers: headers,
		})
		for k, v := range message.headers {
			for _, h := range v {
				c.Writer.Header().Add(k, h)
			}
		}
		c.JSON(message.statusCode, message)
	}
}

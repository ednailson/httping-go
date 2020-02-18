package httping

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type HandlerFunc func(request HttpRequest) (statusCode int, response *JSendMessage)
type Methods map[string]bool

type Route struct {
	route   *gin.RouterGroup
	methods Methods
}

func (gp *Route) AddMethod(method string, handler HandlerFunc) error {
	if gp.methods == nil {
		gp.methods = make(map[string]bool)
	}
	value, ok := gp.methods[method]
	if ok == true && value == true {
		return errors.New(fmt.Sprintf("route %s already has a method %s", gp.route.BasePath(), method))
	}
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
	default:
		return &ErrorUnknownMethod{}
	}
	gp.methods[method] = true
	return nil
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
		code, message := handle(HttpRequest{
			Body:   body,
			Query:  query,
			Params: params,
		})
		if message == nil {
			c.Writer.WriteHeader(code)
			return
		}
		c.JSON(code, message)
	}
}

type ErrorUnknownMethod struct{}

func (e *ErrorUnknownMethod) Error() string {
	return "unknown http method"
}

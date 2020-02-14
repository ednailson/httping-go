package httping

import (
	"bytes"
	. "github.com/onsi/gomega"
	"net/http"
	"strconv"
	"testing"
)

const port = 5001
const baseUrl = "http://localhost:5001"

func TestNewServer(t *testing.T) {
	RegisterTestingT(t)
	server := NewHttpServer(port)
	Expect(server.engine).ShouldNot(BeNil())
	Expect(server.server).ShouldNot(BeNil())
}

func TestNewRoute(t *testing.T) {
	RegisterTestingT(t)
	server := NewHttpServer(port)
	Expect(server.engine).ShouldNot(BeNil())
	Expect(server.server).ShouldNot(BeNil())
	route := server.NewRoute(nil, "/test")
	Expect(route.route).ShouldNot(BeNil())
}

func TestNewRouteWithMethod(t *testing.T) {
	RegisterTestingT(t)
	server := NewHttpServer(port)
	Expect(server.engine).ShouldNot(BeNil())
	Expect(server.server).ShouldNot(BeNil())
	route := server.NewRoute(nil, "/test")
	Expect(route.route).ShouldNot(BeNil())
	err := route.AddMethod(http.MethodPost, func(request HttpRequest) (statusCode int, response *JSendMessage) {
		if string(request.Body) == "success" {
			return http.StatusOK, NewJSendMessage(http.StatusOK)
		}
		return http.StatusNotAcceptable, NewJSendMessage(http.StatusNotAcceptable)
	})
	Expect(err).ShouldNot(HaveOccurred())
	closeServer, _ := server.RunServer()
	defer closeServer()
	resp, err := http.Post(baseUrl+"/test", "application/json", bytes.NewReader([]byte("success")))
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.Status).Should(BeEquivalentTo(strconv.Itoa(http.StatusOK) + " OK"))
}

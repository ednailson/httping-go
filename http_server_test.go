package httping

import (
	"bytes"
	"fmt"
	. "github.com/onsi/gomega"
	"net/http"
	"strconv"
	"testing"
)

const port = 5001
const baseUrl = "http://localhost:5001"
const defaultPath = "/test"

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
	route := server.NewRoute(nil, defaultPath)
	Expect(route.route).ShouldNot(BeNil())
}

func TestNewRouteWithPOST(t *testing.T) {
	RegisterTestingT(t)
	server := NewHttpServer(port)
	const method = http.MethodPost
	Expect(server.engine).ShouldNot(BeNil())
	Expect(server.server).ShouldNot(BeNil())
	route := server.NewRoute(nil, defaultPath)
	Expect(route.route).ShouldNot(BeNil())
	err := route.AddMethod(method, func(request HttpRequest) (statusCode int, response *JSendMessage) {
		return http.StatusOK, NewJSendMessage(http.StatusOK)
	})
	Expect(err).ShouldNot(HaveOccurred())
	err = route.AddMethod(method, func(request HttpRequest) (statusCode int, response *JSendMessage) {
		return http.StatusNotAcceptable, NewJSendMessage(http.StatusNotAcceptable)
	})
	Expect(err).Should(HaveOccurred())
	Expect(err.Error()).Should(BeEquivalentTo(fmt.Sprintf("route %s already has a method %s", defaultPath, method)))
}

func TestNewRouteWithGET(t *testing.T) {
	RegisterTestingT(t)
	server := NewHttpServer(port)
	const method = http.MethodGet
	Expect(server.engine).ShouldNot(BeNil())
	Expect(server.server).ShouldNot(BeNil())
	route := server.NewRoute(nil, defaultPath)
	Expect(route.route).ShouldNot(BeNil())
	err := route.AddMethod(method, func(request HttpRequest) (statusCode int, response *JSendMessage) {
		return http.StatusOK, NewJSendMessage(http.StatusOK)
	})
	Expect(err).ShouldNot(HaveOccurred())
	err = route.AddMethod(method, func(request HttpRequest) (statusCode int, response *JSendMessage) {
		return http.StatusNotAcceptable, NewJSendMessage(http.StatusNotAcceptable)
	})
	Expect(err).Should(HaveOccurred())
	Expect(err.Error()).Should(BeEquivalentTo(fmt.Sprintf("route %s already has a method %s", defaultPath, method)))
}

func TestNewRouteWithPUT(t *testing.T) {
	RegisterTestingT(t)
	server := NewHttpServer(port)
	const method = http.MethodPut
	Expect(server.engine).ShouldNot(BeNil())
	Expect(server.server).ShouldNot(BeNil())
	route := server.NewRoute(nil, defaultPath)
	Expect(route.route).ShouldNot(BeNil())
	err := route.AddMethod(method, func(request HttpRequest) (statusCode int, response *JSendMessage) {
		return http.StatusOK, NewJSendMessage(http.StatusOK)
	})
	Expect(err).ShouldNot(HaveOccurred())
	err = route.AddMethod(method, func(request HttpRequest) (statusCode int, response *JSendMessage) {
		return http.StatusNotAcceptable, NewJSendMessage(http.StatusNotAcceptable)
	})
	Expect(err).Should(HaveOccurred())
	Expect(err.Error()).Should(BeEquivalentTo(fmt.Sprintf("route %s already has a method %s", defaultPath, method)))
}

func TestNewRouteWithPATCH(t *testing.T) {
	RegisterTestingT(t)
	server := NewHttpServer(port)
	const method = http.MethodPatch
	Expect(server.engine).ShouldNot(BeNil())
	Expect(server.server).ShouldNot(BeNil())
	route := server.NewRoute(nil, defaultPath)
	Expect(route.route).ShouldNot(BeNil())
	err := route.AddMethod(method, func(request HttpRequest) (statusCode int, response *JSendMessage) {
		return http.StatusOK, NewJSendMessage(http.StatusOK)
	})
	Expect(err).ShouldNot(HaveOccurred())
	err = route.AddMethod(method, func(request HttpRequest) (statusCode int, response *JSendMessage) {
		return http.StatusNotAcceptable, NewJSendMessage(http.StatusNotAcceptable)
	})
	Expect(err).Should(HaveOccurred())
	Expect(err.Error()).Should(BeEquivalentTo(fmt.Sprintf("route %s already has a method %s", defaultPath, method)))
}

func TestNewRouteWithHEAD(t *testing.T) {
	RegisterTestingT(t)
	server := NewHttpServer(port)
	const method = http.MethodHead
	Expect(server.engine).ShouldNot(BeNil())
	Expect(server.server).ShouldNot(BeNil())
	route := server.NewRoute(nil, defaultPath)
	Expect(route.route).ShouldNot(BeNil())
	err := route.AddMethod(method, func(request HttpRequest) (statusCode int, response *JSendMessage) {
		return http.StatusOK, NewJSendMessage(http.StatusOK)
	})
	Expect(err).ShouldNot(HaveOccurred())
	err = route.AddMethod(method, func(request HttpRequest) (statusCode int, response *JSendMessage) {
		return http.StatusNotAcceptable, NewJSendMessage(http.StatusNotAcceptable)
	})
	Expect(err).Should(HaveOccurred())
	Expect(err.Error()).Should(BeEquivalentTo(fmt.Sprintf("route %s already has a method %s", defaultPath, method)))
}

func TestNewRouteWithOPTIONS(t *testing.T) {
	RegisterTestingT(t)
	server := NewHttpServer(port)
	const method = http.MethodOptions
	Expect(server.engine).ShouldNot(BeNil())
	Expect(server.server).ShouldNot(BeNil())
	route := server.NewRoute(nil, defaultPath)
	Expect(route.route).ShouldNot(BeNil())
	err := route.AddMethod(method, func(request HttpRequest) (statusCode int, response *JSendMessage) {
		return http.StatusOK, NewJSendMessage(http.StatusOK)
	})
	Expect(err).ShouldNot(HaveOccurred())
	err = route.AddMethod(method, func(request HttpRequest) (statusCode int, response *JSendMessage) {
		return http.StatusNotAcceptable, NewJSendMessage(http.StatusNotAcceptable)
	})
	Expect(err).Should(HaveOccurred())
	Expect(err.Error()).Should(BeEquivalentTo(fmt.Sprintf("route %s already has a method %s", defaultPath, method)))
}

func TestRunServer(t *testing.T) {
	RegisterTestingT(t)
	server := NewHttpServer(port)
	Expect(server.engine).ShouldNot(BeNil())
	Expect(server.server).ShouldNot(BeNil())
	route := server.NewRoute(nil, defaultPath)
	Expect(route.route).ShouldNot(BeNil())
	err := route.AddMethod(http.MethodPost, func(request HttpRequest) (statusCode int, response *JSendMessage) {
		if string(request.Body) == "success" {
			return http.StatusOK, NewJSendMessage(http.StatusOK)
		}
		return http.StatusNotAcceptable, NewJSendMessage(http.StatusNotAcceptable)
	})
	Expect(err).ShouldNot(HaveOccurred())
	closeServer, chErr := server.RunServer()
	defer closeServer()
	resp, err := http.Post(baseUrl+defaultPath, "application/json", bytes.NewReader([]byte("success")))
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.Status).Should(BeEquivalentTo(strconv.Itoa(http.StatusOK) + " OK"))
	resp, err = http.Post(baseUrl+defaultPath, "application/json", bytes.NewReader([]byte("not success")))
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.Status).Should(BeEquivalentTo(strconv.Itoa(http.StatusNotAcceptable) + " Not Acceptable"))
	Eventually(chErr).ShouldNot(Receive())
	closeServer2, chErr2 := server.RunServer()
	defer closeServer2()
	resp, err = http.Post(baseUrl+defaultPath, "application/json", bytes.NewReader([]byte("success")))
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.Status).Should(BeEquivalentTo(strconv.Itoa(http.StatusOK) + " OK"))
	resp, err = http.Post(baseUrl+defaultPath, "application/json", bytes.NewReader([]byte("not success")))
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.Status).Should(BeEquivalentTo(strconv.Itoa(http.StatusNotAcceptable) + " Not Acceptable"))
	Eventually(chErr2).ShouldNot(Receive())
	Eventually(chErr).ShouldNot(Receive())
}

func TestCloseServerFunc(t *testing.T) {
	RegisterTestingT(t)
	server := NewHttpServer(port)
	Expect(server.engine).ShouldNot(BeNil())
	Expect(server.server).ShouldNot(BeNil())
	route := server.NewRoute(nil, defaultPath)
	Expect(route.route).ShouldNot(BeNil())
	err := route.AddMethod(http.MethodPost, func(request HttpRequest) (statusCode int, response *JSendMessage) {
		if string(request.Body) == "success" {
			return http.StatusOK, NewJSendMessage(http.StatusOK)
		}
		return http.StatusNotAcceptable, NewJSendMessage(http.StatusNotAcceptable)
	})
	Expect(err).ShouldNot(HaveOccurred())
	closeServer, chErr := server.RunServer()
	resp, err := http.Post(baseUrl+defaultPath, "application/json", bytes.NewReader([]byte("success")))
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.Status).Should(BeEquivalentTo(strconv.Itoa(http.StatusOK) + " OK"))
	resp, err = http.Post(baseUrl+defaultPath, "application/json", bytes.NewReader([]byte("not success")))
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.Status).Should(BeEquivalentTo(strconv.Itoa(http.StatusNotAcceptable) + " Not Acceptable"))
	Eventually(chErr).ShouldNot(Receive())
	err = closeServer()
	_, err = http.Post(baseUrl+defaultPath, "application/json", bytes.NewReader([]byte("not success")))
	Expect(err).Should(HaveOccurred())
}

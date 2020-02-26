package httping

import (
	"bytes"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
	"time"
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
		return http.StatusOK, NewJSend(http.StatusOK)
	})
	Expect(err).ShouldNot(HaveOccurred())
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
		return http.StatusOK, NewJSend(http.StatusOK)
	})
	Expect(err).ShouldNot(HaveOccurred())
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
		return http.StatusOK, NewJSend(http.StatusOK)
	})
	Expect(err).ShouldNot(HaveOccurred())
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
		return http.StatusOK, NewJSend(http.StatusOK)
	})
	Expect(err).ShouldNot(HaveOccurred())
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
		return http.StatusOK, NewJSend(http.StatusOK)
	})
	Expect(err).ShouldNot(HaveOccurred())
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
		return http.StatusOK, NewJSend(http.StatusOK)
	})
	Expect(err).ShouldNot(HaveOccurred())
}

func TestRunServer(t *testing.T) {
	RegisterTestingT(t)
	server := NewHttpServer(port)
	Expect(server.engine).ShouldNot(BeNil())
	Expect(server.server).ShouldNot(BeNil())
	route := server.NewRoute(nil, defaultPath)
	Expect(route.route).ShouldNot(BeNil())
	err := route.AddMethod(http.MethodPost, func(request HttpRequest) (int, *JSendMessage) {
		if string(request.Body) == "success" {
			return http.StatusOK, NewJSend(http.StatusOK)
		}
		return http.StatusNotAcceptable, NewJSend(http.StatusNotAcceptable)
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

func TestResponseWithStruct(t *testing.T) {
	RegisterTestingT(t)
	server := NewHttpServer(port)
	Expect(server.engine).ShouldNot(BeNil())
	Expect(server.server).ShouldNot(BeNil())
	route := server.NewRoute(nil, defaultPath)
	Expect(route.route).ShouldNot(BeNil())
	err := route.AddMethod(http.MethodPost, func(request HttpRequest) (int, *JSendMessage) {
		type TestResponse struct {
			Test  string `json:"test"`
			Test2 string `json:"test2"`
		}
		return http.StatusOK, NewJSend(http.StatusOK).AddData(TestResponse{
			Test:  "field 1",
			Test2: "field2",
		})
	})
	Expect(err).ShouldNot(HaveOccurred())
	closeServer, chErr := server.RunServer()
	defer closeServer()
	time.Sleep(5 * time.Millisecond)
	resp, err := http.Post(baseUrl+defaultPath, "application/json", bytes.NewReader([]byte("success")))
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.Status).Should(BeEquivalentTo(strconv.Itoa(http.StatusOK) + " OK"))
	body, err := ioutil.ReadAll(resp.Body)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(body).Should(MatchJSON([]byte(`{"status":"success","data":{"test":"field 1","test2":"field2"}}`)))
	Eventually(chErr).ShouldNot(Receive())
}

func TestRequestAndResponseWithHeaders(t *testing.T) {
	RegisterTestingT(t)
	server := NewHttpServer(port)
	Expect(server.engine).ShouldNot(BeNil())
	Expect(server.server).ShouldNot(BeNil())
	route := server.NewRoute(nil, defaultPath)
	Expect(route.route).ShouldNot(BeNil())
	err := route.AddMethod(http.MethodPost, func(request HttpRequest) (int, *JSendMessage) {
		response := NewJSend(http.StatusOK)
		Expect(request.Headers["Header-Test"][0]).Should(BeEquivalentTo("header test 1"))
		Expect(request.Headers["Header-Test"][1]).Should(BeEquivalentTo("header test 2"))
		Expect(request.Headers["Header-Test2"][0]).Should(BeEquivalentTo("header test 3"))
		return http.StatusOK, response.
			AddData("test").
			AddHeader("Response-Header", "response").
			AddHeader("Response-Header", "response 2")
	})
	Expect(err).ShouldNot(HaveOccurred())
	closeServer, chErr := server.RunServer()
	defer closeServer()
	time.Sleep(5 * time.Millisecond)
	req, err := http.NewRequest(http.MethodPost, baseUrl+defaultPath, bytes.NewReader([]byte("success")))
	Expect(err).ShouldNot(HaveOccurred())
	req.Header.Add("Header-Test", "header test 1")
	req.Header.Add("Header-Test", "header test 2")
	req.Header.Add("Header-Test2", "header test 3")
	client := http.DefaultClient
	resp, err := client.Do(req)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.Status).Should(BeEquivalentTo(strconv.Itoa(http.StatusOK) + " OK"))
	body, err := ioutil.ReadAll(resp.Body)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(body).Should(MatchJSON([]byte(`{"status":"success","data":"test"}`)))
	Expect(resp.Header["Response-Header"][0]).Should(BeEquivalentTo("response"))
	Expect(resp.Header["Response-Header"][1]).Should(BeEquivalentTo("response 2"))
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
			return http.StatusOK, NewJSend(http.StatusOK)
		}
		return http.StatusNotAcceptable, NewJSend(http.StatusNotAcceptable)
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

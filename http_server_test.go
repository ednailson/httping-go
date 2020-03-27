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
	handleFunc := func(request HttpRequest) *ResponseMessage {
		return NewResponse(http.StatusOK)
	}
	route.AddMethod(method, handleFunc)
	Expect(func() {
		route.POST(handleFunc)
	}).To(Panic())
}

func TestNewRouteWithGET(t *testing.T) {
	RegisterTestingT(t)
	server := NewHttpServer(port)
	const method = http.MethodGet
	Expect(server.engine).ShouldNot(BeNil())
	Expect(server.server).ShouldNot(BeNil())
	route := server.NewRoute(nil, defaultPath)
	Expect(route.route).ShouldNot(BeNil())
	handleFunc := func(request HttpRequest) *ResponseMessage {
		return NewResponse(http.StatusOK)
	}
	route.AddMethod(method, handleFunc)
	Expect(func() {
		route.GET(handleFunc)
	}).To(Panic())
}

func TestNewRouteWithPUT(t *testing.T) {
	RegisterTestingT(t)
	server := NewHttpServer(port)
	const method = http.MethodPut
	Expect(server.engine).ShouldNot(BeNil())
	Expect(server.server).ShouldNot(BeNil())
	route := server.NewRoute(nil, defaultPath)
	Expect(route.route).ShouldNot(BeNil())
	handleFunc := func(request HttpRequest) *ResponseMessage {
		return NewResponse(http.StatusOK)
	}
	route.AddMethod(method, handleFunc)
	Expect(func() {
		route.PUT(handleFunc)
	}).To(Panic())
}

func TestNewRouteWithPATCH(t *testing.T) {
	RegisterTestingT(t)
	server := NewHttpServer(port)
	const method = http.MethodPatch
	Expect(server.engine).ShouldNot(BeNil())
	Expect(server.server).ShouldNot(BeNil())
	route := server.NewRoute(nil, defaultPath)
	Expect(route.route).ShouldNot(BeNil())
	handleFunc := func(request HttpRequest) *ResponseMessage {
		return NewResponse(http.StatusOK)
	}
	route.AddMethod(method, handleFunc)
	Expect(func() {
		route.PATCH(handleFunc)
	}).To(Panic())
}

func TestNewRouteWithHEAD(t *testing.T) {
	RegisterTestingT(t)
	server := NewHttpServer(port)
	const method = http.MethodHead
	Expect(server.engine).ShouldNot(BeNil())
	Expect(server.server).ShouldNot(BeNil())
	route := server.NewRoute(nil, defaultPath)
	Expect(route.route).ShouldNot(BeNil())
	handleFunc := func(request HttpRequest) *ResponseMessage {
		return NewResponse(http.StatusOK)
	}
	route.AddMethod(method, handleFunc)
	Expect(func() {
		route.HEAD(handleFunc)
	}).To(Panic())
}

func TestNewRouteWithOPTIONS(t *testing.T) {
	RegisterTestingT(t)
	server := NewHttpServer(port)
	const method = http.MethodOptions
	Expect(server.engine).ShouldNot(BeNil())
	Expect(server.server).ShouldNot(BeNil())
	route := server.NewRoute(nil, defaultPath)
	Expect(route.route).ShouldNot(BeNil())
	handleFunc := func(request HttpRequest) *ResponseMessage {
		return NewResponse(http.StatusOK)
	}
	route.AddMethod(method, handleFunc)
	Expect(func() {
		route.OPTIONS(handleFunc)
	}).To(Panic())
}

func TestRunServer(t *testing.T) {
	RegisterTestingT(t)
	server := NewHttpServer(port)
	Expect(server.engine).ShouldNot(BeNil())
	Expect(server.server).ShouldNot(BeNil())
	route := server.NewRoute(nil, defaultPath)
	Expect(route.route).ShouldNot(BeNil())
	route.AddMethod(http.MethodPost, func(request HttpRequest) *ResponseMessage {
		if string(request.Body) == "success" {
			return NewResponse(http.StatusOK)
		}
		return NewResponse(http.StatusNotAcceptable)
	})
	closeServer, chErr := server.RunServer()
	defer closingServer(closeServer)
	resp, err := http.Post(baseUrl+defaultPath, "application/json", bytes.NewReader([]byte("success")))
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).Should(BeEquivalentTo(http.StatusOK))
	resp, err = http.Post(baseUrl+defaultPath, "application/json", bytes.NewReader([]byte("not success")))
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).Should(BeEquivalentTo(http.StatusNotAcceptable))
	Eventually(chErr).ShouldNot(Receive())
	closeServer2, chErr2 := server.RunServer()
	defer closingServer(closeServer2)
	resp, err = http.Post(baseUrl+defaultPath, "application/json", bytes.NewReader([]byte("success")))
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).Should(BeEquivalentTo(http.StatusOK))
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
	route.AddMethod(http.MethodPost, func(request HttpRequest) *ResponseMessage {
		type TestResponse struct {
			Test  string `json:"test"`
			Test2 string `json:"test2"`
		}
		return NewResponse(http.StatusOK).AddData(TestResponse{
			Test:  "field 1",
			Test2: "field2",
		})
	})
	closeServer, chErr := server.RunServer()
	defer closingServer(closeServer)
	time.Sleep(5 * time.Millisecond)
	resp, err := http.Post(baseUrl+defaultPath, "application/json", bytes.NewReader([]byte("success")))
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).Should(BeEquivalentTo(http.StatusOK))
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
	route.AddMethod(http.MethodPost, func(request HttpRequest) *ResponseMessage {
		response := NewResponse(http.StatusOK)
		Expect(request.Headers["Header-Test"][0]).Should(BeEquivalentTo("header test 1"))
		Expect(request.Headers["Header-Test"][1]).Should(BeEquivalentTo("header test 2"))
		Expect(request.Headers["Header-Test2"][0]).Should(BeEquivalentTo("header test 3"))
		return response.
			AddData("test").
			AddHeader("Response-Header", "response").
			AddHeader("Response-Header", "response 2")
	})
	closeServer, chErr := server.RunServer()
	defer closingServer(closeServer)
	time.Sleep(5 * time.Millisecond)
	req, err := http.NewRequest(http.MethodPost, baseUrl+defaultPath, bytes.NewReader([]byte("success")))
	Expect(err).ShouldNot(HaveOccurred())
	req.Header.Add("Header-Test", "header test 1")
	req.Header.Add("Header-Test", "header test 2")
	req.Header.Add("Header-Test2", "header test 3")
	client := http.DefaultClient
	resp, err := client.Do(req)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).Should(BeEquivalentTo(http.StatusOK))
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
	route.AddMethod(http.MethodPost, func(request HttpRequest) *ResponseMessage {
		if string(request.Body) == "success" {
			return NewResponse(http.StatusOK)
		}
		return NewResponse(http.StatusNotAcceptable)
	})
	closeServer, chErr := server.RunServer()
	resp, err := http.Post(baseUrl+defaultPath, "application/json", bytes.NewReader([]byte("success")))
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).Should(BeEquivalentTo(http.StatusOK))
	resp, err = http.Post(baseUrl+defaultPath, "application/json", bytes.NewReader([]byte("not success")))
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).Should(BeEquivalentTo(http.StatusNotAcceptable))
	Eventually(chErr).ShouldNot(Receive())
	err = closeServer()
	_, err = http.Post(baseUrl+defaultPath, "application/json", bytes.NewReader([]byte("not success")))
	Expect(err).Should(HaveOccurred())
}

func TestServerWithMiddleware(t *testing.T) {
	RegisterTestingT(t)
	const token = "b4357690-1a01-4fc5-8243-2c2f32b9fc26"
	server := NewHttpServer(port).SetMiddleware(func(request HttpRequest) (*ResponseMessage, bool) {
		if request.Headers["Authorization"][0] != token {
			return Unauthorized("not authorized"), false
		}
		return nil, true
	})
	server.NewRoute(nil, defaultPath).AddMethod(http.MethodPost, func(request HttpRequest) *ResponseMessage {
		return OK("middleware ok")
	})
	closeServer, chErr := server.RunServer()
	defer closingServer(closeServer)
	time.Sleep(5 * time.Millisecond)
	req, err := http.NewRequest(http.MethodPost, baseUrl+defaultPath, nil)
	Expect(err).ShouldNot(HaveOccurred())
	req.Header.Add("Authorization", token)
	client := http.DefaultClient
	resp, err := client.Do(req)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).Should(BeEquivalentTo(http.StatusOK))
	body, err := ioutil.ReadAll(resp.Body)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(body).Should(MatchJSON([]byte(`{"status":"success","data":"middleware ok"}`)))
	req.Header.Del("Authorization")
	req.Header.Add("Authorization", "wrong token")
	resp, err = client.Do(req)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).Should(BeEquivalentTo(http.StatusUnauthorized))
	Eventually(chErr).ShouldNot(Receive())
}

func TestNullResponsesOnMiddleware(t *testing.T) {
	RegisterTestingT(t)
	RegisterTestingT(t)
	server := NewHttpServer(port).SetMiddleware(func(request HttpRequest) (*ResponseMessage, bool) {
		return nil, false
	})
	server.NewRoute(nil, defaultPath).AddMethod(http.MethodPost, func(request HttpRequest) *ResponseMessage {
		return OK("success")
	})
	closeServer, chErr := server.RunServer()
	defer closingServer(closeServer)
	time.Sleep(5 * time.Millisecond)
	resp, err := http.Post(baseUrl+defaultPath, "application/json", nil)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).Should(BeEquivalentTo(http.StatusOK))
	body, err := ioutil.ReadAll(resp.Body)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(string(body)).Should(BeEquivalentTo("null\n"))
	Eventually(chErr).ShouldNot(Receive())
}

func TestNullResponsesOnHandler(t *testing.T) {
	RegisterTestingT(t)
	RegisterTestingT(t)
	server := NewHttpServer(port)
	server.NewRoute(nil, defaultPath).AddMethod(http.MethodPost, func(request HttpRequest) *ResponseMessage {
		return nil
	})
	closeServer, chErr := server.RunServer()
	defer closingServer(closeServer)
	time.Sleep(5 * time.Millisecond)
	resp, err := http.Post(baseUrl+defaultPath, "application/json", nil)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).Should(BeEquivalentTo(http.StatusOK))
	body, err := ioutil.ReadAll(resp.Body)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(string(body)).Should(BeEquivalentTo("null\n"))
	Eventually(chErr).ShouldNot(Receive())
}

func TestRequestAndResponseWithCookies(t *testing.T) {
	RegisterTestingT(t)
	server := NewHttpServer(port)
	cookie := &http.Cookie{
		Name:  "test-cookie",
		Value: "value-cookie",
	}
	server.NewRoute(nil, defaultPath).POST(func(request HttpRequest) *ResponseMessage {
		Expect(request.Cookies[0].Name).To(BeEquivalentTo(cookie.Name))
		Expect(request.Cookies[0].Value).To(BeEquivalentTo(cookie.Value))
		return OK("test").AddCookie(cookie)
	})
	closeServer, chErr := server.RunServer()
	defer closingServer(closeServer)
	time.Sleep(5 * time.Millisecond)
	req, err := http.NewRequest(http.MethodPost, baseUrl+defaultPath, nil)
	Expect(err).ShouldNot(HaveOccurred())
	req.AddCookie(cookie)
	client := http.DefaultClient
	resp, err := client.Do(req)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).Should(BeEquivalentTo(http.StatusOK))
	Expect(resp.Cookies()[0].Name).To(BeEquivalentTo(cookie.Name))
	Expect(resp.Cookies()[0].Value).To(BeEquivalentTo(cookie.Value))
	body, err := ioutil.ReadAll(resp.Body)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(body).Should(MatchJSON([]byte(`{"status":"success","data":"test"}`)))
	Eventually(chErr).ShouldNot(Receive())
}

func closingServer(closeServerFn ServerCloseFunc) {
	err := closeServerFn()
	Expect(err).ShouldNot(HaveOccurred())
}

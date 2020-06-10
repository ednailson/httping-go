package main

import (
	"fmt"
	"github.com/ednailson/httping-go"
	"net/http"
)

//The httping package has a helper for responses. But you can build your own response builder

func main() {

	//Creating a server localhost port 8080
	server := httping.NewHttpServer("", 8080)

	//Setting up a route with a custom response
	server.NewRoute(nil, "/example").POST(func(request httping.HttpRequest) httping.IResponse {
		return &response{
			data:       "OK!",
			statusCode: http.StatusOK,
		}
	})

	//Running the server
	//If you close the server, you just need to call the closeServerFunc()
	//chErr will receive any error that happened on the server
	//We will have 1 routes on this server:
	//POST /example
	closeServeFunc, chErr := server.RunServer()

	<-chErr
	err := closeServeFunc()
	if err != nil {
		fmt.Sprintln(err.Error())
	}
}

//To create a custom response it is only needed to implement the interface IResponse
type response struct {
	data       string
	statusCode int
}

func (r *response) Headers() map[string][]string {
	return nil
}

func (r *response) Cookies() []*http.Cookie {
	return nil
}

func (r *response) Response() interface{} {
	return r.data
}

func (r *response) StatusCode() int {
	return r.statusCode
}

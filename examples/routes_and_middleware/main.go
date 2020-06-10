package main

import (
	"fmt"
	"github.com/ednailson/httping-go"
)

func main() {

	//Creating a server localhost port 8080
	server := httping.NewHttpServer("", 8080)

	//If your server will be accessed by browsers we recommend you to enable cors
	server.EnableCORS()

	//Adding a middleware func to the server. It means the every request to the server will pass by the middleware
	//It is possible to remove the middleware to a specific route. The example is down below.
	server.AddMiddleware(func(request httping.HttpRequest) httping.IResponse {
		if request.Body == nil {

			//This lib has some helpers to create response with JSend pattern.
			//Here we are returning a BadRequest with a json data.
			return httping.BadRequest(`{"body": "body is required"}`)
		}

		//If the middleware handle function return nil
		//It means the the request will continue to the next handle functions
		return nil
	})

	//Setting up a route
	v1Route := server.NewRoute(nil, "/v1")

	//Setting up a new route with a route base
	//The route will be localhost:8080/v1/example
	exampleRoute := server.NewRoute(v1Route, "/example")

	//Adding a method POST to the exampleRoute with the handleFunc()
	exampleRoute.POST(handleFunc())

	//Adding a method GET to the same exampleRoute with the handleFunc2()
	exampleRoute.GET(handleFunc2())

	//Creating a new route with a route base without the server middleware handle function
	//Adding already a POST method to the route with the handleFunc2()
	server.NewRoute(v1Route, "/noMiddleware").SetMiddleware(nil).POST(handleFunc2())

	//Running the server
	//If you close the server, you just need to call the closeServerFunc()
	//chErr will receive any error that happened on the server
	//We will have 3 routes on this server:
	//POST /v1/example
	//GET /v1/example
	//POST /v1/noMiddleware
	closeServeFunc, chErr := server.RunServer()

	<-chErr
	err := closeServeFunc()
	if err != nil {
		fmt.Sprintln(err.Error())
	}
}

func handleFunc() httping.HandlerFunc {
	return func(request httping.HttpRequest) httping.IResponse {
		if request.Params["authorization"] == "" {
			return httping.Unauthorized(map[string]string{
				"authorization": "unauthorized",
			})
		}

		//Anything with the request can be made

		return httping.NoContent()
	}
}

func handleFunc2() httping.HandlerFunc {
	return func(request httping.HttpRequest) httping.IResponse {

		//Anything with the request can be made

		return httping.OK("OK!")
	}
}

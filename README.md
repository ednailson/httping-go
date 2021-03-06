# Httping go

[![Release](https://img.shields.io/badge/release-0.15.0-blue)](https://github.com/ednailson/httping-go/releases)
![Coverage](https://img.shields.io/badge/coverage-100%25-success)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)
[![GoDoc](https://img.shields.io/badge/godoc-reference-9cf)](https://godoc.org/github.com/ednailson/httping-go)

A library for creating API's on golang that comes with a [JSend responses](https://github.com/omniti-labs/jsend) helper

* **[CHANGELOG](CHANGELOG.md)**

Checkout some [example](examples).

## Getting started

**Download**

```bash
go get github.com/ednailson/httping-go
```

### Creating a server

```go
server := httping.NewHttpServer("", 3000)
``` 

Here a **http server** was created on the port `3000` and from the server is possible to create _http routes_.

The server has support to CORS. To enable it, you just need to do

```go
server.EnableCORS()
```

### Creating a route

```go
routeExample := server.NewRoute(nil, "/example")
```

Now there a **route**. It is possible to add http methods to this **route** with a treatment function. 

#### Creating a route from a route

It is possible to create a **route** from a **route**. This allows to create **routes** from a unique **route**.

```go
routeCreateExample := server.NewRoute(routeExample, "/create")
```

```go
routeUpdateExample := server.NewRoute(routeExample, "/update")
```

So now there are two new **routes**: `http://localhost:3000/example/create` and `http://localhost:3000/example/update`.

### Adding a method on the route

```go
routeExample.AddMethod("POST", func(request httping.HttpRequest) httping.IResponse {
    if len(request.body) == 0 {
        return httping.NewResponse(404)
    }
    return httping.NewResponse(200)
})
```

A **method** `POST` is now available on the **route** `http://localhost:3000/example`.

_p.s.: only http methods and http codes are allowed_

And it is possible to add different **methods** on the same **route**. 

```go
routeExample.AddMethod("GET", func(request httping.HttpRequest) httping.IResponse {
    if len(request.body) == 0 {
        return httping.NewResponse(404)
    }
    return httping.NewResponse(200)
})
```

Now the route `http://localhost:3000/example` has the **methods** `GET` and `POST`.

If you will not use the route two or more times you can directly create a route and add a method 

```go
server.NewRoute(nil, "/create").AddMethod("POST", func(request httping.HttpRequest) httping.IResponse {
		return httping.NewResponse(200)
	})
```

### Responding

Responses for the `handleFunc()`

For creating a `ResponseMessage`

```go
response := httping.NewResponse(200)
```

This will build a Response message with the status correct according with the http status code and [jsend](https://github.com/omniti-labs/jsend) pattern.

**Example**

```go
server.NewRoute(nil, "/create").POST(func(request httping.HttpRequest) httping.IResponse {
		return httping.NewResponse(200).AddData("success")
	})
```

It respects the [jsend](https://github.com/omniti-labs/jsend)'s pattern. 

On **responses** it also possible to add Headers, Code and Message.

There are a few helpers for the most commons http status codes.

**Example**

```go
server.NewRoute(nil, "/create").POST(func(request httping.HttpRequest) httping.IResponse {
		return httping.OK("data example")
	})
```

It will return a status code ok (200) with the data `data example`
You can check all the helpers [here](CHANGELOG.md#050).

### Middleware on the server or route

It is possible to add a middleware handler function to the server or a route

```go
server := httping.NewHttpServer(3000).AddMiddleware(
    func(request httping.HttpRequest) httping.IResponse {
        if request.Headers["Authorization"][0] != "token"{
            return httping.Unauthorized("not authorized")
        }
        return nil
    }
)
```

If you return `IResponse`: The server will **not** let the request proceed and it will return the response returned.

If you return `nil`, the server will let the request proceed to the route's `handleFunc`

Middleware can also be applied only on a route. If the server has a middleware, the function will be added by the route's middleware function.

It is possible to set middleware. It will replaced all the middleware functions on that route or server.

# Developer

[Júnior Vilas Boas](http://ednailson.github.io)
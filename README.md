# Httping go

A helper to create APIs on golang with [JSend responses](https://github.com/omniti-labs/jsend)

## Getting started

**Download**

```bash
go get github.com/ednailson/httping-go
```

### Creating a server

```go
server := httping.NewHttpServer(3000)
``` 

Here a **http server** was created on the port `3000` and from the server is possible to create _http routes_.

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
err := routeExample.AddMethod("POST", func(request HttpRequest) (int, *JSendMessage) {
    if len(request.body) === 0 {
        return 404, httping.NewJSend(404)
    }
    return 200, httping.NewJSend(200)
})
```

A **method** `POST` is now available on the **route** `http://localhost:3000/example`.

_p.s.: only http methods and http codes are allowed_

And it is possible to add different **methods** on the same **route**. 

```go
err := routeExample.AddMethod("GET", func(request HttpRequest) (int, *JSendMessage) {
    if len(request.body) === 0 {
        return 404, httping.NewJSend(404)
    }
    return 200, httping.NewJSend(200)
})
```

Now the route `http://localhost:3000/example` has the **methods** `GET` and `POST`.

If you will not use the route two or more times you can directly create a route and add a method 

```go
err := server.NewRoute(nil, "/create").AddMethod("POST", func(request httping.HttpRequest) (int, *httping.JSendMessage) {
		return http.StatusOK, httping.NewJSend(http.StatusOK)
	})
```

### Response helpers

This lib also brings some helpers for the response for the `handleFunc()`

For creating a `JSendResponse`

```go
response := httping.NewJSend(200)
```

This will build a JSend message with the status correct according with the http status code.

**Example**

```go
err := server.NewRoute(nil, "/create").AddMethod("POST", func(request httping.HttpRequest) (int, *httping.JSendMessage) {
		return http.StatusOK, httping.NewJSend(200).AddData("success")
	})
```

# Developer

[Júnior Vilas Boas](http://ednailson.github.io)
# CHANGELOG 

## 0.1.0

* Server creator
* Route creator
* Add method to route creator
* Run server feature 

## 0.2.0

* Server does not accept two methods at the same route. `panic.Err()` is generated
* `NewJSend` method renamed
* `JSendMessage` methods to add data and message

## 0.3.0

* Headers has been added into the **request** and the **response**
* Creating responses now is `NewResponse` instead of `NewJSend`

### 0.3.1

* Fix the error report on the channel received from `RunServer()`

## 0.4.0

* Only `NewResponse()` method is required on the response when is added a new method

## 0.5.0

* `Code` in `ResponseMessage` for error responses
* Helpers for responses. **Status**:
    * OK
    * Created
    * Accepted
    * NoContent
    * BadRequest
    * Unauthorized
    * Forbidden
    * NotFound
    * MethodNotAllowed
    * NotAcceptable
    * InternalServerError

### 0.5.1

* Code and message can only be added on the `StatusError` messages

## 0.6.0

* `AddMethod` does not returns error. If the method is not known it will not be added any method

## 0.7.0

* It is possible to add a middleware function to the **http server**
* Responses **null** will return http status code of success (200)

### 0.7.1

* Helpers for adding methods on routes
    * POST
    * GET
    * DELETE
    * PUT
    * PATCH
    * HEAD
    * OPTIONS

### 0.7.2

* Cookies is manageable on the **request** and **response** now on.

## 0.8.0

* Get for `status`, `headers` and `cookies` on `ResponseMessage`.
* Fixing response no content(_status 204_) to respond no body

### 0.8.1

* Go mod fixing

## 0.9.0

* Middleware functions only returns `ResponseMessage`. If it is the `MiddlewareFunc` returns `null` the middleware will
continue the request for the `HandleFunc`
* `NewServer` requires a host. It can run local if it receives `""` as host parameter.
* Middleware function can be set on IRoute. It will be replace the server middleware function only on that route.
* `NewServer` and `NewRoute` return only interfaces now on.

### 0.9.1

* CORS support on the server

## 0.10.0

* Support a many middleware functions
* `SetMiddleware` sets a slice of `HandleFunc`
* `AddMiddleware` add a new function to the middleware

## 0.11.0

* Fixing go mod

## 0.12.0

* Fixing go mod
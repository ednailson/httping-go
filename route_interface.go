package httping

type Route interface {
	AddMethod(method string, handler HandlerFunc)
	POST(handler HandlerFunc)
	GET(handler HandlerFunc)
	DELETE(handler HandlerFunc)
	PUT(handler HandlerFunc)
	PATCH(handler HandlerFunc)
	HEAD(handler HandlerFunc)
	OPTIONS(handler HandlerFunc)
	SetMiddleware(middleware []HandlerFunc) Route
	AddMiddleware(middleware HandlerFunc) Route
	getRoute() *route
}

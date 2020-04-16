package httping

type IRoute interface {
	AddMethod(method string, handler HandlerFunc)
	POST(handler HandlerFunc)
	GET(handler HandlerFunc)
	DELETE(handler HandlerFunc)
	PUT(handler HandlerFunc)
	PATCH(handler HandlerFunc)
	HEAD(handler HandlerFunc)
	OPTIONS(handler HandlerFunc)
	SetMiddleware(middleware []HandlerFunc) IRoute
	AddMiddleware(middleware HandlerFunc) IRoute
	getRoute() *route
}

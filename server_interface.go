package httping

type IServer interface {
	NewRoute(group *Route, path string) *Route
	RunServer() (ServerCloseFunc, chan error)
	SetMiddleware(middleware MiddlewareFunc) *HttpServer
}

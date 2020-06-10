package httping

type IServer interface {
	NewRoute(group IRoute, path string) IRoute
	RunServer() (ServerCloseFunc, chan error)
	SetMiddleware(middleware []HandlerFunc) IServer
	AddMiddleware(middleware HandlerFunc) IServer
	EnableCORS() IServer
}

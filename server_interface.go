package httping

type Server interface {
	NewRoute(group Route, path string) Route
	RunServer() (ServerCloseFunc, chan error)
	SetMiddleware(middleware []HandlerFunc) Server
	AddMiddleware(middleware HandlerFunc) Server
	EnableCORS() Server
}

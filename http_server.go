package httping

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func NewHttpServer(host string, port int) IServer {
	engine := gin.Default()
	engine.HandleMethodNotAllowed = true
	server := &http.Server{
		Addr:    host + ":" + strconv.Itoa(port),
		Handler: engine,
	}
	return &httpServer{server: server, engine: engine}
}

type httpServer struct {
	server     *http.Server
	engine     *gin.Engine
	middleware HandlerFunc
}

func (server *httpServer) NewRoute(baseRoute IRoute, path string) IRoute {
	if baseRoute != nil {
		g := baseRoute.getRoute().route.Group(path)
		return &route{route: g, middleware: baseRoute.getRoute().middleware}
	}
	g := server.engine.Group(path)
	return &route{route: g, middleware: server.middleware}
}

func (server *httpServer) RunServer() (ServerCloseFunc, chan error) {
	chErr := make(chan error)
	go func(server *http.Server) {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			chErr <- err
		}
	}(server.server)
	return func() error {
		return server.server.Close()
	}, chErr
}

func (server *httpServer) SetMiddleware(middleware HandlerFunc) IServer {
	server.middleware = middleware
	return server
}

type ServerCloseFunc func() error

package httping

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func NewHttpServer(port int) *HttpServer {
	engine := gin.Default()
	engine.HandleMethodNotAllowed = true
	server := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: engine,
	}
	return &HttpServer{server: server, engine: engine}
}

type HttpServer struct {
	server *http.Server
	engine *gin.Engine
}

func (server *HttpServer) NewRoute(baseRoute *Route, path string) *Route {
	if baseRoute != nil {
		g := baseRoute.route.Group(path)
		return &Route{route: g}
	}
	g := server.engine.Group(path)
	return &Route{route: g}
}

func (server *HttpServer) RunServer() (ServerCloseFunc, chan error) {
	chErr := make(chan error)
	go func(server *http.Server) {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			<-chErr
		}
	}(server.server)
	return func() error {
		return server.server.Close()
	}, chErr
}

type ServerCloseFunc func() error

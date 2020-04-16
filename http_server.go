package httping

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func NewHttpServer(host string, port int, cors ...bool) IServer {
	engine := gin.Default()
	engine.HandleMethodNotAllowed = true
	if len(cors) > 0 && cors[0] == true {
		engine.Use(corsMiddleware())
	}
	server := &http.Server{
		Addr:    host + ":" + strconv.Itoa(port),
		Handler: engine,
	}
	return &httpServer{server: server, engine: engine, middleware: []HandlerFunc{}}
}

type httpServer struct {
	server     *http.Server
	engine     *gin.Engine
	middleware []HandlerFunc
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

func (server *httpServer) SetMiddleware(middleware []HandlerFunc) IServer {
	server.middleware = middleware
	return server
}

func (server *httpServer) AddMiddleware(middleware HandlerFunc) IServer {
	server.middleware = append(server.middleware, middleware)
	return server
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

type ServerCloseFunc func() error

package httpserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ServerMode string

const (
	// DebugMode indicates gin mode is debug.
	DebugMode ServerMode = "debug"
	// ReleaseMode indicates gin mode is release.
	ReleaseMode ServerMode = "release"
	// TestMode indicates gin mode is test.
	TestMode ServerMode = "test"
)

type Server struct {
	handler http.Handler
}

func New(mode ...ServerMode) Server {
	if len(mode) > 0 {
		gin.SetMode(string(mode[0]))
	} else {
		gin.SetMode(string(ReleaseMode))
	}
	return Server{
		handler: gin.Default(),
	}
}

func (s *Server) AddHandle(method, path string, inputStructPtr any, f func(inputStruct any) (any, error)) {
	s.handler.(*gin.Engine).Handle(method, path, func(c *gin.Context) {
		if err := c.ShouldBindJSON(inputStructPtr); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		data, err := f(inputStructPtr)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	})
}

func (s *Server) AddGinHandle(method, path string, f gin.HandlerFunc) {
	s.handler.(*gin.Engine).Handle(method, path, f)
}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.handler)
}

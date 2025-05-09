package httpserver

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type ServerMode string

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

const (
	DebugMode   ServerMode = "debug"
	ReleaseMode ServerMode = "release"
	TestMode    ServerMode = "test"
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

func (s *Server) AddHandle(path string, f func(input []byte) (Response, error)) {
	s.handler.(*gin.Engine).Handle(http.MethodPost, path, func(c *gin.Context) {
		body := c.Request.Body
		bodyByte, _ := io.ReadAll(body)
		data, err := f(bodyByte)
		if err != nil {
			c.JSON(http.StatusOK, err.Error())
			return
		}
		c.JSON(http.StatusOK, data)
	})
}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.handler)
}

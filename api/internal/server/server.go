package server

import (

	"github.com/labstack/echo/v4"
)

type Server struct {
	echoServer *echo.Echo
}

func (s *Server) Run(port string, echoServer *echo.Echo) error {
	s.echoServer = echoServer

	return s.echoServer.Start(port)
}
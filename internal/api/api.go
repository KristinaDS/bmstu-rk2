package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Server struct {
	server  *echo.Echo
	address string

	uc Usecase
}

func NewServer(ip string, port int, uc Usecase) *Server {
	api := Server{
		uc: uc,
	}

	api.server = echo.New()
	api.server.POST("/users", api.CreateUser)
	api.server.GET("/users", api.ListUsers)
	api.server.GET("/users/:id", api.GetUser)
	api.server.PUT("/users/:id", api.UpdateUser)
	api.server.DELETE("/users/:id", api.DeleteUser)

	api.server.POST("/events", api.CreateEvent)
	api.server.GET("/events", api.GetEvents)
	api.server.PUT("/events/:id", api.UpdateEvent)
	api.server.DELETE("/events/:id", api.DeleteEvent)

	api.address = fmt.Sprintf("%s:%d", ip, port)

	return &api
}

func (s *Server) Run() {
	s.server.Logger.Fatal(s.server.Start(s.address))
}

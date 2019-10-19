package server

import (
	"StreamChannelSwitch/config"
	"StreamChannelSwitch/controller"
	"StreamChannelSwitch/datasource"
	"errors"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type Server struct {
	addr        string
	e           *echo.Echo
	controllers map[string]controller.Controller
	datasource  *datasource.RedisSouce
}

type SDKDetail struct {
	uid       string
	sid       string
	timestamp string
	token     string
	sdklist   []string
}



func NewServer(config *config.Config, redis *datasource.RedisSouce) *Server {

	return &Server{addr: config.ServerAddr, datasource: redis, e: echo.New(), controllers: map[string]controller.Controller{}}
}

func (s *Server) AddControl(prefix string, controller controller.Controller) error {
	if controller == nil {
		return errors.New("controller can not nil")
	}
	s.controllers[prefix] = controller
	g := s.e.Group(prefix)
	controller.Set(g)
	return nil
}

func (s *Server) Start() error {
	s.e.Logger.SetLevel(log.DEBUG)

	for _, v := range s.controllers {
		_ = v.Init()
	}

	s.e.Logger.Fatal(s.e.Start(":18989"))
	return nil
}

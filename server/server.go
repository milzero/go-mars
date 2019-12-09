package server

import (
	"StreamChannelSwitch/config"
)

type Server struct {

}




func NewServer(config *config.Config, redis *datasource.RedisSouce) *Server {

	return &Server{}
}


func (s *Server) Start() error {
	return nil
}

package main

import (
	"StreamChannelSwitch/config"
	"StreamChannelSwitch/controller"
	"StreamChannelSwitch/datasource"
	"StreamChannelSwitch/server"
	"StreamChannelSwitch/util"
)

func main() {
	util.Logger.Info("start")
	conf := config.NewConfig("config.cfg")
	_ = server2.Start()
}

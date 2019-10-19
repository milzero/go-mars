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
	redis := datasource.NewRedisSouce(conf)
	mongo := datasource.NewMongoSource(conf)
	server2 := server.NewServer(conf, redis)
	before := controller.NewBeforeClass(mongo, redis, "123")
	in := controller.NewInClass(mongo, redis, nil , "123")
	_ = server2.AddControl("/switcher/before", before)
	_ = server2.AddControl("/switcher/in", in)
	_ = server2.Start()
}

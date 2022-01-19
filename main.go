package main

import (
	"chat/conf"
	"chat/router"
	"chat/service"
)

func main() {
	go service.Manager.Start()
	r := router.NewRouter()
	_ = r.Run(conf.HttpPort)
}

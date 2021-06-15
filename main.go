package main

import (
	_ "hello/orthodonticsAdmin/bootstrap"
	"hello/orthodonticsAdmin/global/variable"
	"hello/orthodonticsAdmin/router"
)

func main() {
	r := router.InitRouter()
	address := variable.ConfigViper.GetString("httpserver.address")
	port := variable.ConfigViper.GetString("httpserver.port")
	r.Run(address + ":" + port)
}

package main

import (
	"fmt"
	"go/sevice/config"
	"go/sevice/web/routers"
	// "server_demo/utils"
)

var(
	conf = config.GetConfig()
	// logger = utils.GetLogger("main")
)

func main(){
	server := routers.RegisterRoutes()
	// fmt.Println("Service starting")
	err := server.Run(fmt.Sprintf(":%d", conf.Port))
	if err != nil {
		fmt.Println("Service startup failed!")
 	}
}

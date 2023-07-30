package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"math2do.in/notification/configs"
	"math2do.in/notification/handlers"
	"math2do.in/notification/utils"
)

func main() {
	utils.InitLogger()

	err := configs.LoadConfigs()
	if err != nil {
		log.Fatalf("err:%v", err)
	}
	config := configs.Config

	router := gin.Default()

	handlers.AddRoutes(router)
	addr := fmt.Sprintf("%s:%d", config.ServiceURL, config.Port)
	log.Infof("Server running in %s", addr)

	router.Run(addr)
	if err != nil {
		log.Fatalf("err:%v", err)
	}
}

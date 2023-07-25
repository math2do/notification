package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"math2do.in/notification/configs"
	"math2do.in/notification/handlers"
)

func main() {
	log := log.Default()

	err := configs.LoadConfigs(log)
	if err != nil {
		log.Fatalf("err:%v", err)
	}

	router := gin.Default()

	handlers.AddRoutes(router)

	router.Run(fmt.Sprintf(":%d", configs.Config.Port))
}

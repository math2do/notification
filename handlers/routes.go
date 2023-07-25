package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.Engine) {
	route := router.Group("/api/v1")

	route.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Group various endpoints
	{

	}
}

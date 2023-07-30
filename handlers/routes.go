package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.Engine) {
	route := router.Group("/api/v1")

	route.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Group various endpoints
	email := route.Group("/email")
	{
		email.POST("/verification-email", HandleVerificationEmail)
		email.POST("/verify", HandleVerificationEmail)
	}
}

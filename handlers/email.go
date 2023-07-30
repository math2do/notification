package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"math2do.in/notification/dtos"
	"math2do.in/notification/service/email"
)

func HandleVerificationEmail(c *gin.Context) {
	var req dtos.VerifyEmailReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// process the request by email service
	service := email.NewGmailSender()

	err = service.SendEmailForVerification(req)
	if err != nil {
		log.Errorf("error in send email, err:%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

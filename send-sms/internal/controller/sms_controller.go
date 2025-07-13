package controller

import (
	"net/http"
	"send-sms/internal/repository"
	"send-sms/internal/service"

	"github.com/gin-gonic/gin"
)

type SMSRequest struct {
	To      string `json:"to"`
	Message string `json:"message"`
}

func SendSMSHandler(redisHost string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req SMSRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		if err := repository.StoreSMS(redisHost, req.To, req.Message); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to store sms"})
			return
		}

		if err := service.SendSMS(req.To, req.Message); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send sms"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "sms sent"})
	}
}

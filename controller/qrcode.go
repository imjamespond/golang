package controller

import (
	"4d-qrcode/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenQrcode(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{model.Userkey: nil, model.Expiredkey: nil})
}

func Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
}

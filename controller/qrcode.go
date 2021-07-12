package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenQrcode(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

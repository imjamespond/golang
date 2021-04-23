package main

import (
	"net/http"
	"test-gin-auth/model"
	"time"

	"github.com/gin-gonic/gin"
)

func Version(c *gin.Context) {
	// https://github.com/gin-gonic/gin/issues/1335
	// gin is based on net/http package, and that serves each request by an individual goroutine.
	time.Sleep(3 * time.Second)
	c.JSON(http.StatusOK, gin.H{"ver": model.Version})
}

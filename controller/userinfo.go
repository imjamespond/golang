package controller

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"test-gin-auth/model"
)

func Me(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(model.Userkey)
	expired := session.Get(model.Expiredkey)
	c.JSON(http.StatusOK, gin.H{model.Userkey: user, model.Expiredkey: expired})
}

func Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
}

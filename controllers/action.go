package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func Action(c *gin.Context) {
	var body struct {
		Type    string `json:"type"`
		Payload any    `json:"payload"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var result any
	var err error

	switch body.Type {
	case "save_url":
		result, err = SaveUrl(body.Payload)
	default:
		c.JSON(400, gin.H{"error": "unknown type"})
		return
	}

	if err != nil {
		if errors.Is(err, ErrInvalidPayload) {
			c.JSON(400, gin.H{"error": err.Error()})
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(200, result)
}

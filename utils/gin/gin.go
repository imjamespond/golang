package gin

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error, status int) bool {
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		c.Abort()
		return true
	}
	return false
}

func HandleBadRequest(c *gin.Context, err error) bool {
	return HandleError(c, err, http.StatusBadRequest)
}

func HandleUnauthorized(c *gin.Context, err error) bool {
	return HandleError(c, err, http.StatusUnauthorized)
}

func Handle404(c *gin.Context, err string) bool {
	return HandleError(c, errors.New(err), http.StatusNotFound)
}

func HandleJSON(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}

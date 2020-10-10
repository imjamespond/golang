package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//
// APP error definition
//
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//
// Middleware Error Handler in server package
//
func JSONAppErrorReporter() gin.HandlerFunc {
	return jsonAppErrorReporterT(gin.ErrorTypeAny)
}

// Error implements error interface
func (v *AppError) Error() string {
	return ""
}

func jsonAppErrorReporterT(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedErrors := c.Errors.ByType(errType)

		log.Println("Handle APP error")
		if len(detectedErrors) > 0 {
			err := detectedErrors[0].Err
			var parsedError *AppError
			switch err.(type) {
			case *AppError:
				parsedError = err.(*AppError)
			default:
				parsedError = &AppError{
					Code: http.StatusInternalServerError, Message: "Internal Server Error",
				}
			}
			// Put the error into response
			c.IndentedJSON(parsedError.Code, parsedError)
			c.Abort()
			// or c.AbortWithStatusJSON(parsedError.Code, parsedError)
			return
		}

	}
}

func main() {
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.Use(JSONAppErrorReporter())
	r.GET("/test", func(c *gin.Context) {
		//c.Error(&AppError{Code: 401, Message: "a prank"})
		c.Error(errors.New("a prank"))
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

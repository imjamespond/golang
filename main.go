package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// UserData 模拟数据库或内存中的用户数据
func EmailPtr(e string) *openapi_types.Email {
    email := openapi_types.Email(e)
    return &email
}
var UserData = []User{
	{Id: 1, Name: "Alice", Email: EmailPtr("alice@example.com")},
	{Id: 2, Name: "Bob", Email: EmailPtr("bob@example.com")},
}

// Ensure our handler struct satisfies the ServerInterface
var _ ServerInterface = (*Server)(nil)

// Server defines our application's server
type Server struct{}

// GetUsers handles the GET /users endpoint
func (s *Server) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, UserData)
}

// CreateUser handles the POST /users endpoint
func (s *Server) CreateUser(c *gin.Context) {
	var newUser NewUser
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lastID := int64(0)
	if len(UserData) > 0 {
		lastID = UserData[len(UserData)-1].Id
	}

	user := User{
		Id:    lastID + 1,
		Name:  newUser.Name,
		Email: newUser.Email,
	}

	UserData = append(UserData, user)

	c.JSON(http.StatusCreated, user)
}

func main() {
	// Create a new router
	router := gin.Default()

	// Get our OpenAPI server interface
	server := &Server{}

	// Register our handlers
	RegisterHandlers(router, server)

	fmt.Println("Server is running on http://localhost:8080")
	router.Run(":8080")
}
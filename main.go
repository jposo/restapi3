package main

import (
	"net/http"

	"slices"

	"github.com/gin-gonic/gin"
)

func main() {
	// Static data
	users := []map[string]interface{}{
		{"id": 1, "name": "Alice", "email": "alice@example.com"},
		{"id": 2, "name": "Bob", "email": "bob@example.com"},
	}

	// Create a router
	gin.SetMode(gin.ReleaseMode) // Set Gin to release mode for production
	router := gin.Default()

	// Define API endpoints
	router.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	})

	router.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, user := range users {
			if user["id"] == id {
				c.JSON(http.StatusOK, user)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
	})

	router.POST("/users", func(c *gin.Context) {
		var newUser struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Add new user to the static data
		newUserID := len(users) + 1
		users = append(users, map[string]any{
			"id":    newUserID,
			"name":  newUser.Name,
			"email": newUser.Email,
		})

		c.JSON(http.StatusOK, gin.H{"status": "user added", "id": newUserID})
	})

	router.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedUser struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}
		if err := c.ShouldBindJSON(&updatedUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for i, user := range users {
			if user["id"] == id {
				users[i]["name"] = updatedUser.Name
				users[i]["email"] = updatedUser.Email
				c.JSON(http.StatusOK, gin.H{"status": "user updated"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
	})

	router.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, user := range users {
			if user["id"] == id {
				users = slices.Delete(users, i, i+1)
				c.JSON(http.StatusOK, gin.H{"status": "user deleted"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
	})

	// Start the server
	router.Run(":8080")
}

package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    // Static data
    users := []map[string]interface{}{
        {"id": 1, "name": "Alice", "email": "alice@example.com"},
        {"id": 2, "name": "Bob", "email": "bob@example.com"},
    }

    // Create a router
    router := gin.Default()

    // Define API endpoints
    router.GET("/users", func(c *gin.Context) {
        c.JSON(http.StatusOK, users)
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
        users = append(users, map[string]interface{}{
            "id":    newUserID,
            "name":  newUser.Name,
            "email": newUser.Email,
        })

        c.JSON(http.StatusOK, gin.H{"status": "user added", "id": newUserID})
    })

    // Start the server
    router.Run(":8080")
}

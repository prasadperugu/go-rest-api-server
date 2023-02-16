package main

import (
	//go get github.com/gin-gonic/gin
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{ID: 1, Name: "Nandu", Email: "nandu.ahmed@onymos.com"},
	{ID: 2, Name: "Prasad", Email: "prasad.perugu@onymos.com"},
}

func main() {
	r := gin.Default()

	// Route to get all users
	r.GET("/users", getUsers)

	// Route to get a specific user by ID
	r.GET("/users/:id", getUser) //curl http://localhost:9090/users/1

	// Route to create a new user
	r.POST("/users", createUser)

	// Route to update an existing user
	r.PUT("/users/:id", updateUser)

	// Route to delete an existing user
	r.DELETE("/users/:id", deleteUser)

	r.Run(":9090")
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func getUser(c *gin.Context) {
	id := c.Param("id")

	for _, user := range users {
		if strconv.Itoa(user.ID) == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = len(users) + 1
	users = append(users, user)

	c.JSON(http.StatusCreated, user)
}

func updateUser(c *gin.Context) {
	id := c.Param("id")

	for i, user := range users {
		if strconv.Itoa(user.ID) == id {
			var updatedUser User
			if err := c.ShouldBindJSON(&updatedUser); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			updatedUser.ID = user.ID
			users[i] = updatedUser

			c.JSON(http.StatusOK, updatedUser)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")

	for i, user := range users {
		if strconv.Itoa(user.ID) == id {
			users = append(users[:i], users[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

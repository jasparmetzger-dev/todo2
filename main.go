package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var store *Store = NewStore() // initialize the store, could be a db

	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//auth logic
	auth := r.Group("/auth")
	{
		auth.POST("/login", Login(store))
		auth.POST("/register", Register(store))
	}

	//protected logic
	api := r.Group("/api")
	api.Use(AuthMiddleware(store)) // apply the auth middleware to all /api routes
	{
		// example protected route
		api.GET("/profile", Profile(store))
		api.PATCH("/profile", UpdateProfile(store))

		api.GET("/todos", ListTodos(store))
		api.POST("/todos", CreateTodo(store))
		api.GET("/todos/fulfilled", ListFulfilled(store))

		api.PATCH("/todos/:id", UpdateTodo(store))
		api.DELETE("/todos/:id", DeleteTodo(store))
		api.PATCH("/todos/:id/fulfill", FulfillTodo(store))

	}

	r.Run("localhost:8080") //  default to :8080
}

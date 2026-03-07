package main

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(store *Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		//validate header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "authorization header is required"})
			c.Abort()
			return
		}
		//create token
		var tokenString string = authHeader[len("Bearer "):]
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(MY_SECRET_KEY), nil
		})
		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}
		claims, ok := token.Claims.(*Claims)
		if !ok {
			c.JSON(401, gin.H{"error": "invalid token claims"})
			c.Abort()
			return
		}

		fmt.Printf("JWT extreacted user_id: %d, \n", claims.UserId)
		c.Set("user_id", claims.UserId)

		//validate id
		if _, err := store.GetUser(claims.UserId); err != nil {
			c.JSON(401, gin.H{"error": "user not found"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log the request method and path
		// Note: No real Logging
		//Missing: log to a file, log level, log format, time to complete, etc.

		fmt.Printf("Request: %s %s\n", c.Request.Method, c.Request.URL.Path)
		c.Next()
	}
}

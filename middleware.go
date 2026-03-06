package main

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "authorization header is required"})
			c.Abort()
			return
		}
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
		c.Set("user_id", claims.UserId)
		c.Next()
	}
}

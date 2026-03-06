package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const MY_SECRET_KEY string = "shhhh, im secure"

type Claims struct {
	UserId uint64 `json:"user_id"`
	jwt.RegisteredClaims
}

func generateJWT(userId uint64) (string, error) {
	claims := Claims{
		UserId: userId,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(MY_SECRET_KEY))
}

func Register(s *Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		//validate input
		var body struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.BindJSON(&body); err != nil || body.Username == "" || body.Password == "" {
			c.JSON(400, gin.H{"error": "invalid body, username and password are required"})
			return
		}
		//check if username already exists
		for _, user := range s.UserMap {
			if user.Username == body.Username {
				c.JSON(400, gin.H{"error": "username already exists"})
				return
			}
		}
		//create user and generate token
		user := CreateUser(body.Username, body.Password)

		s.AddUser(user)
		token, err := generateJWT(user.Id)

		if err != nil {
			c.JSON(500, gin.H{"error": "token error"})
			return
		}
		c.JSON(201, gin.H{"token": token})
	}
}

func Login(s *Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		//validate input
		var body struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.BindJSON(&body); err != nil || body.Username == "" || body.Password == "" {
			c.JSON(400, gin.H{"error": "invalid body, correct username and password are required"})
			return
		}
		//generate token
		for _, user := range s.UserMap {
			if user.Username == body.Username && user.Password == body.Password {
				token, err := generateJWT(user.Id)
				if err != nil {
					c.JSON(500, gin.H{"error": "token error"})
					return
				}
				c.JSON(200, gin.H{"token": token})
				return
			}
		}
		c.JSON(401, gin.H{"error": "invalid username or password"})
	}
}

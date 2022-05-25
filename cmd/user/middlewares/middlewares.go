package middlewares

import (
	"errors"
	"go-blog/cmd/user"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type JWTClaimm struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func Auth(c *fiber.Ctx) {
	tokenString := c.Get("Authorization")
	newToken := strings.Split(tokenString, "Bearer ")[1]
	println(newToken)
	if tokenString == "" {
		c.Status(401).JSON(user.Response{Status: false})
	}
	err := ValidateToken(newToken)
	if err != nil {

		return
	}
	c.Next()
	return
}

func ValidateToken(signedToken string) (c *fiber.Ctx) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaimm{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		},
	)
	if err != nil {
		errors.New("Invalid token")
		return
	}
	claims, ok := token.Claims.(*JWTClaimm)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}

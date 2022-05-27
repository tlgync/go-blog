package middlewares

import (
	"errors"
	"go-blog/cmd/user"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type JWTClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type NewAuthCheck interface {
	Auth(c *fiber.Ctx)
}

func Auth(c *fiber.Ctx) {
	tokenString := c.Get("Authorization")
	newToken := strings.Split(tokenString, "Bearer ")[1]
	if tokenString == "" {
		err := c.Status(401).JSON(user.Response{Status: false})
		if err != nil {
			panic("FATAL ERROR")
		}
	}
	err := ValidateToken(newToken)
	if err != nil {
		err := c.Status(401).JSON(user.Response{Status: false})
		if err != nil {
			panic("FATAL ERROR")
		}
	}
	return
}

func ValidateToken(signedToken string) (c *fiber.Ctx) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		},
	)
	if err != nil {
		err := errors.New("invalid token")
		if err != nil {
			return nil
		}
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
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

package auth

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("SecretYouShouldHide")

type JWTClaims struct {
	jwt.MapClaims
}

func Authorize(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()

	bearerSplit := strings.SplitAfter(headers["Authorization"], "Bearer ")
	if len(bearerSplit) != 2 {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	tokenString := bearerSplit[1]
	jwtClaims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, jwtClaims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if token.Valid {
		return c.Next()
	} else {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// //To use claims, use this
	// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	// 	fmt.Printf("Claims JWT: %+v", claims)
	// 	return c.Next()
	// } else {
	// 	return c.SendStatus(fiber.StatusUnauthorized)
	// }
}

func Generate() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "go-crud",
		// Simple claims but you set the defaults
	})
	tokenString, err := token.SignedString(secretKey)
	return tokenString, err
}

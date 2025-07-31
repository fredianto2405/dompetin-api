package jwt

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type Claims struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(id int, email string) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	claims := Claims{
		ID:    id,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "dompetin",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ValidateJWT(tokenStr string) (*Claims, error) {
	secret := os.Getenv("JWT_SECRET")

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(MsgUnexpectedSigningMethod)
		}
		return []byte(secret), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New(MsgInvalidToken)
		}
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New(MsgInvalidToken)
	}

	return claims, nil
}

func GetUserClaims(c *gin.Context) (*Claims, bool) {
	user, exists := c.Get("user")
	if !exists {
		return nil, false
	}

	claims, ok := user.(*Claims)
	return claims, ok
}

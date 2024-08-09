package validation

import (
	"avito_bootcamp/internal/entity"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

var jwtKey = []byte("secret_key")

type Claims struct {
	Email    string `json:"email"`
	UserType string `json:"user_type"`
	jwt.RegisteredClaims
}

func GetToken(c echo.Context) (string, error) {
	cookie, err := c.Cookie("token")
	if err != nil {
		return "", fmt.Errorf("no cookie, %d", http.StatusUnauthorized)
	}
	return cookie.Value, nil
}

func CreateToken(user entity.User) (string, error) {
	claims := &Claims{
		Email:    user.Email,
		UserType: user.UserType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func CheckToken(tokenString string) (Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return Claims{}, fmt.Errorf("signature invalid, %d", http.StatusUnauthorized)
		}
		return Claims{}, fmt.Errorf("bad Request, %d", http.StatusBadRequest)
	}

	if !token.Valid {
		return Claims{}, fmt.Errorf("token is invalid, %d", http.StatusUnauthorized)
	}
	return *claims, nil
}

func CheckModerator(tokenString string) error {
	claims, err := CheckToken(tokenString)
	if err != nil {
		return err
	}
	if claims.UserType == "moderator" {
		return nil
	}
	return fmt.Errorf("token is invalid, %d", http.StatusForbidden)
}

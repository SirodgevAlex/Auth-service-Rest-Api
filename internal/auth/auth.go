package auth

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserID int
	jwt.StandardClaims
}

var jwtKey = []byte("1234")

func IsEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func IsPasswordSafe(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasDigit && hasSpecial
}

func ExtractUserIdFromToken(r *http.Request) (int, error) {
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return 0, fmt.Errorf("ошибка при парсинге токена: %v", err)
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		user_id, err := strconv.Atoi(claims.Subject)
		if err != nil {
			return 0, fmt.Errorf("ошибка при извлечении идентификатора пользователя из токена: %w", err)
		}
		return user_id, nil
	}

	return 0, errors.New("невозможно извлечь идентификатор пользователя из токена")
}
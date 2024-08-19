package utils

import (
	"time"

	"github.com/KozlovNikolai/test-task/internal/pkg/config"
	"github.com/golang-jwt/jwt/v5"
)

// Claims is ...
type Claims struct {
	Login string `json:"login"`
	jwt.RegisteredClaims
}

// GenerateJWT is ...
// func GenerateJWT(user model.User) (string, time.Time, error) {
// 	//TokenTimeDuration := 30 * time.Minute
// 	// Create JWT token
// 	expirationTime := time.Now().Add(config.Cfg.TokenTimeDuration)
// 	//expirationTime := time.Now().Add(TokenTimeDuration)
// 	claims := &Claims{
// 		Login: user.Login,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(expirationTime),
// 		},
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString(config.JwtKey)
// 	if err != nil {
// 		return "", time.Time{}, err
// 	}
// 	return tokenString, expirationTime, nil
// }

// GenerateJWT is ...
func GenerateJWT(login, role string) (string, error) {
	claims := jwt.MapClaims{
		"login": login,
		"role":  role,
		"exp":   time.Now().Add(config.Cfg.TokenTimeDuration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JwtKey)
}

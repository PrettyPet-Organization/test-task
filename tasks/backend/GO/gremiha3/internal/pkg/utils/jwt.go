package utils

import (
	"time"

	"github.com/KozlovNikolai/test-task/internal/pkg/config"
	"github.com/golang-jwt/jwt/v5"
)

// Claims is ...
type Claims struct {
	AuthLogin string `json:"auth_login"`
	AuthRole  string `json:"auth_role"`
	jwt.RegisteredClaims
}

// GenerateJWT is ...
func GenerateJWT(login, role string) (string, error) {
	claims := jwt.MapClaims{
		"auth_login": login,
		"auth_role":  role,
		"exp":        time.Now().Add(config.Cfg.TokenTimeDuration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JwtKey)
}

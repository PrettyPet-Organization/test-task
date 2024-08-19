package middlewares

import (
	"net/http"

	"github.com/KozlovNikolai/test-task/internal/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Claims is ...
type Claims struct {
	Login string `json:"login"`
	jwt.RegisteredClaims
}

// AuthMiddleware is ...
// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		cookie, err := c.Request.Cookie("token")
// 		fmt.Println("###########################")
// 		fmt.Printf("cookie = %v\n", cookie.Value)
// 		fmt.Println("###########################")
// 		if err != nil {
// 			if err == http.ErrNoCookie {
// 				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 				return
// 			}
// 			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
// 			return
// 		}

// 		// Parse JWT token
// 		tknStr := cookie.Value
// 		claims := &Claims{}
// 		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
// 			return config.JwtKey, nil
// 		})

// 		if err != nil {
// 			if err == jwt.ErrSignatureInvalid {
// 				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 				return
// 			}
// 			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
// 			return
// 		}

// 		if !tkn.Valid {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 			return
// 		}

// 		// Set username in context
// 		c.Set("login", claims.Login)

// 		c.Next()
// 	}
// }

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		tokenString := authHeader[len("Bearer "):]
		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return config.JwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
			return
		}

		if !tkn.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Сохранение имени пользователя в контексте запроса
		c.Set("login", claims.Login)

		c.Next()
	}
}

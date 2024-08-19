package utils

import "github.com/gin-gonic/gin"

func GetLevel(c *gin.Context) (string, string) {
	login, _ := c.Get("auth_login")
	role, _ := c.Get("auth_role")
	return login.(string), role.(string)
}

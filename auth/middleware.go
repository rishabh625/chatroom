package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthenticateHeader(c *gin.Context) {
	//c.Set("clientid", "rishabh")
	//c.Next()
	authHeader := c.Request.Header["Authorization"]
	sentToken := ""
	if len(authHeader) > 0 {
		bearer := strings.Split(authHeader[0], " ")
		if len(bearer) > 0 {
			if bearer[0] == "Bearer" {
				sentToken = bearer[1]
				if sentToken == "" {
					c.JSON(http.StatusUnauthorized, "{\"message\":\"Token not found\"}")
					c.Abort()
				}
			} else {
				c.JSON(http.StatusUnauthorized, "{\"message\":\"Authorization header must start with \" Bearer\"\"}")
				c.Abort()
			}
		} else {
			c.JSON(http.StatusUnauthorized, "{\"message\":\"Authorization header must start with \" Bearer\"\"}")
			c.Abort()
		}
	} else {
		c.JSON(http.StatusUnauthorized, "{\"message\":\"Authorization header is expected\"}")
		c.Abort()
	}
	if sentToken != "" {
		clientid, err := TokenValid(sentToken)
		if err != nil || clientid == "" {
			c.JSON(http.StatusUnauthorized, "{\"message\":\"Authorization Failed, invalid token\"}")
			c.Abort()
		} else {
			c.Set("clientid", clientid)
			c.Set("token", sentToken)
			c.Next()
		}
	}
}

func CheckAdmin(c *gin.Context) {
	clientid := c.MustGet("clientid").(string)
	if clientid != "admin" {
		c.JSON(http.StatusUnauthorized, "")
		c.Abort()
	}
}

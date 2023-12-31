package middleware

import (
	"fmt"
	"net/http"

	helper "github.com/MrSharma003/GO-JWT/helpers"
	"github.com/gin-gonic/gin"
)

func Autheticates() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Unauthorized access")})
			c.Abort()
			return
		}

		claims, msg := helper.ValidateToken(clientToken)
		if msg != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("uid", claims.Uid)
		c.Set("user_type", claims.User_type)
		c.Next()

	}
}

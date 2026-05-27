package middleware

import (
	"net/http"
	"strings"
	"todo-app/internal/config"
	"todo-app/internal/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {

		//grab the full "Bearer  eyjhbg...from the request header"
		authHeader := c.GetHeader("Authorization")

		//if no token was sent at all block immediately//

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "authorization header required",
			})
			c.Abort()
			return
		}

		//split "Bearer eyjhbg...into["bearer","eyjhbg..."]//
		//the empty string here must have a space else would split every single character//
		parts := strings.Split(authHeader, " ")

		//must have exactly two parts and first part must be Bearer//
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization format",
			})
			c.Abort()
			return
		}
		//so middleware extract the header and pass the token into validatejwt//
		//parts[1] is the raw token "eyjhbg..."
		tokenString := parts[1]

		//pass raw token to ValidateJWT to verify and extract UserId from validate token//
		//the userId comes from validate token function which returned the id//
		userId, err := utils.ValidateJWT(tokenString, cfg)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid or expired token",
			})
			c.Abort()
			return
		}
		//store userId in context so handlers can access it//
		//means save authenticated user's id inside this request context//
		//to still help handlers know who user is after middleware finishes//
		c.Set("userId", userId)

		//everything passed then let request through to handler//
		c.Next()
	}
}

package middlewares

import (
	"log"
	"net/http"

	"code.jtg.tools/ayush.singhal/notifications-microservice/configuration"
	"code.jtg.tools/ayush.singhal/notifications-microservice/constants"
	"code.jtg.tools/ayush.singhal/notifications-microservice/shared/auth"
	"github.com/gin-gonic/gin"
)

// AuthorizeJWT validates and authorizes the requests
func AuthorizeJWT() gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader(constants.Authorization)
		headerPrefix := configuration.GetResp().Token.HeaderPrefix
		headerCheck := len(authHeader) < (len(headerPrefix)+2) || authHeader[:len(headerPrefix)] != headerPrefix

		if headerCheck {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString := authHeader[len(headerPrefix)+1:]

		token, err := auth.ValidateToken(tokenString)
		if err != nil {
			log.Println(err.Error())
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(*auth.CustomClaims)
		if token.Valid && claims.TokenType == constants.TokenType().Access {
			c.Set(constants.ID, claims.UserID)
			c.Set(constants.Role, claims.Role)
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

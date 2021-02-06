package auth

import (
	"log"
	"net/http"

	"code.jtg.tools/ayush.singhal/notifications-microservice/db/models"
	"code.jtg.tools/ayush.singhal/notifications-microservice/features/services/users"
	"code.jtg.tools/ayush.singhal/notifications-microservice/shared/auth"
	"github.com/gin-gonic/gin"
)

//ValidateEmail Controller verifies the email after checking the token
func ValidateEmail(c *gin.Context) {

	tokenString := c.Param("token")

	token, err := auth.ValidateToken(tokenString)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims := token.Claims.(*auth.CustomClaims)
	var userDetails models.User

	if token.Valid && claims.TokenType == "validation" {

		err = users.GetUserWithID(&userDetails, claims.UserID)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"internal_error": "Internal Server Error"})
			return
		}

		userDetails.Verified = true
		err = users.PatchUser(&userDetails)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"internal_error": "Internal Server Error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"email_id_verified": "your email id was verified"})

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}

package notifications

import (
	"net/http"

	"code.jtg.tools/ayush.singhal/notifications-microservice/app/services/notifications"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GetAPILastRoute is used to get last 8 characters of API Key
func GetAPILastRoute(router *gin.RouterGroup) {
	router.GET("", GetAPILast)
}

// GetAPILast function is a controller for the get notifications/ route
func GetAPILast(c *gin.Context) {
	apiLast, err := notifications.GetAPILast()
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No API Key exists",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"api_last": apiLast,
	})
}
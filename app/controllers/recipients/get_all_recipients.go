package recipients

import (
	"net/http"

	"code.jtg.tools/ayush.singhal/notifications-microservice/app/serializers"
	"code.jtg.tools/ayush.singhal/notifications-microservice/app/services/recipients"
	"code.jtg.tools/ayush.singhal/notifications-microservice/constants"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GetAllRecipientRoute is used to get recipients from database
func GetAllRecipientRoute(router *gin.RouterGroup) {
	router.GET("", GetAllRecipient)
}

// GetAllRecipient Controller for get /recipient route
func GetAllRecipient(c *gin.Context) {

	recipientArray, err := recipients.GetAllRecipients()
	if err != nil && err != gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": constants.Errors().InternalError})
		return
	}

	var info serializers.RecipientInfo
	var infoArray []serializers.RecipientInfo

	for _, recipient := range recipientArray {
		serializers.RecipientModelToRecipientInfo(&info, &recipient)
		infoArray = append(infoArray, info)
	}

	c.JSON(http.StatusOK, infoArray)
}

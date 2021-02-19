package recipients

import (
	"log"
	"net/http"
	"strconv"

	"code.jtg.tools/ayush.singhal/notifications-microservice/app/controllers/preflight"
	"code.jtg.tools/ayush.singhal/notifications-microservice/app/serializers"
	"code.jtg.tools/ayush.singhal/notifications-microservice/app/services/channels"
	"code.jtg.tools/ayush.singhal/notifications-microservice/app/services/recipients"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GetRecipientRoute is used to get recipients from database
func GetRecipientRoute(router *gin.RouterGroup) {
	router.GET("/:id", GetRecipient)
	router.OPTIONS("/:id", preflight.Preflight)
}

// GetRecipient Controller for get /recipient/:id route
func GetRecipient(c *gin.Context) {
	recipientID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID should be a unigned integer"})
		log.Println("String Conversion Error")
		return
	}
	recipient, err := recipients.GetRecipientWithID(uint64(recipientID))
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id not in database"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	var info serializers.RecipientInfo
	serializers.RecipientModelToRecipientInfo(&info, recipient)

	if info.PreferredChannelID != 0 {
		var channelInfo serializers.ChannelInfo
		channel, err := channels.GetChannelWithID(uint(info.PreferredChannelID))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		serializers.ChannelModelToChannelInfo(&channelInfo, channel)
		c.JSON(http.StatusOK, gin.H{
			"recipient_details": info,
			"preferred_channel": channelInfo,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"recipient_details": info,
	})
}

package channels

import (
	"log"
	"net/http"

	"code.jtg.tools/ayush.singhal/notifications-microservice/app/serializers"
	"code.jtg.tools/ayush.singhal/notifications-microservice/app/services/channels"
	"github.com/gin-gonic/gin"
)

// GetAllChannelsRoute is used to get all the channels from the database
func GetAllChannelsRoute(router *gin.RouterGroup) {
	router.GET("", GetAllChannels)
}

// GetAllChannels function is a controller for the get channels/ route
func GetAllChannels(c *gin.Context) {

	channelList, err := channels.GetAllChannels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		log.Println("find all channels query error")
		return
	}

	var infoArray []serializers.ChannelInfo
	var info serializers.ChannelInfo

	for _, channel := range channelList {
		serializers.ChannelModelToChannelInfo(&info, &channel)
		infoArray = append(infoArray, info)
	}
	c.JSON(http.StatusOK, infoArray)
}
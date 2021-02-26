package recipients

import (
	"log"
	"net/http"

	"code.jtg.tools/ayush.singhal/notifications-microservice/app/serializers"
	"code.jtg.tools/ayush.singhal/notifications-microservice/app/services/recipients"
	"github.com/gin-gonic/gin"
)

// AddUpdateRecipientRoute is used to allow creation and updation of recipients from csv
func AddUpdateRecipientRoute(router *gin.RouterGroup) {
	router.POST("/csv", AddUpdateRecipient)
}

// AddUpdateRecipient controller for post /recipient/csv route
func AddUpdateRecipient(c *gin.Context) {

	rFile, err := c.FormFile("recipients")

	if err != nil {
		log.Println(err)
		var errorList serializers.ErrorInfo
		errorList.Error = map[int][]string{
			1: {"File Format error"},
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errorList)
		return
	}
	recipientRecords, err := recipients.ReadCSV(rFile)
	if err != nil {
		var errorList serializers.ErrorInfo
		errorList.Error = map[int][]string{
			1: {err.Error()},
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errorList)
		return
	}

	status, errorList := recipients.AddUpdateRecipients(recipientRecords)

	if status == http.StatusOK {
		c.JSON(http.StatusOK, gin.H{
			"status":           "OK",
			"records_affected": len(*recipientRecords),
		})
	} else {
		c.AbortWithStatusJSON(status, errorList)
	}
}

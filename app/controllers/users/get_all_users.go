package users

import (
	"log"
	"net/http"

	"code.jtg.tools/ayush.singhal/notifications-microservice/app/serializers"
	"code.jtg.tools/ayush.singhal/notifications-microservice/app/serializers/filter"
	"code.jtg.tools/ayush.singhal/notifications-microservice/app/services/users"
	"code.jtg.tools/ayush.singhal/notifications-microservice/constants"
	"github.com/gin-gonic/gin"
)

// GetAllUsersRoute is used to get all users from database
func GetAllUsersRoute(router *gin.RouterGroup) {
	router.GET("", GetAllUsers)
}

// GetAllUsers Controller for get /users/ route
func GetAllUsers(c *gin.Context) {

	var err error
	var pagination serializers.Pagination
	err = c.BindQuery(&pagination)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": constants.Errors().InvalidPagination,
		})
		return
	}

	var userFilter filter.User
	err = c.BindQuery(&userFilter)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": constants.Errors().InvalidFilter,
		})
		return
	}
	filter.ConvertUserStringToLower(&userFilter)

	recordsCount, err := users.GetAllUsersCount(&userFilter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.Errors().InternalError})
		log.Println("find all users query error")
		return
	}
	var infoArray []serializers.UserInfo
	userListResponse := serializers.UserListResponse{
		RecordsAffected: recordsCount,
		UserInfo:        infoArray,
	}
	if recordsCount == 0 {
		c.JSON(http.StatusOK, userListResponse)
		return
	}

	usersArray, err := users.GetAllUsers(&pagination, &userFilter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.Errors().InternalError})
		log.Println("find all users query error", err.Error())
		return
	}

	var info *serializers.UserInfo

	for _, user := range usersArray {
		info = serializers.UserModelToUserInfo(&user)
		infoArray = append(infoArray, *info)
	}

	userListResponse = serializers.UserListResponse{
		RecordsAffected: recordsCount,
		UserInfo:        infoArray,
	}
	c.JSON(http.StatusOK, userListResponse)
}

package serializers

import "code.jtg.tools/ayush.singhal/notifications-microservice/db/models"

//SignupInfo serializer to bind request data
type SignupInfo struct {
	FirstName	string		`json:"first_name" binding:"required"`
	LastName	string		`json:"last_name"`
	Email        string      `json:"email" binding:"required"`
	Password     string      `json:"password,omitempty" binding:"required"`
	Role		int			`json:"role"`
}

//SignupInfoToUserModel converts SignupInfo serializer to User model
func SignupInfoToUserModel(info SignupInfo,user *models.User){
	user.FirstName = info.FirstName
	user.LastName = info.LastName
	user.Email = info.Email
	user.Password = info.Password
	user.Verified = false
	user.Role = info.Role
}
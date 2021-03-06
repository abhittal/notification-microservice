package serializers

import (
	"time"

	"code.jtg.tools/ayush.singhal/notifications-microservice/db/models"
)

// UserInfo serializer to bind request data
type UserInfo struct {
	ID        uint      `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name,omitempty"`
	Email     string    `json:"email"`
	Verified  bool      `json:"verified"`
	Role      int       `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserListResponse serializer for user list response
type UserListResponse struct {
	RecordsAffected int64      `json:"records_count"`
	UserInfo        []UserInfo `json:"users"`
}

// UserModelToUserInfo converts the user model to UserInfo struct
func UserModelToUserInfo(user *models.User) *UserInfo {
	var info UserInfo
	info.ID = user.ID
	info.FirstName = user.FirstName
	info.LastName = user.LastName
	info.Email = user.Email
	info.Verified = user.Verified
	info.Role = user.Role
	info.CreatedAt = user.CreatedAt
	info.UpdatedAt = user.UpdatedAt
	return &info
}

package models

import (
	"github.com/jinzhu/gorm"
)

// Organisation model to create 'organisations' table in the database
type Organisation struct {
	gorm.Model
	APILast string `gorm:"not null;type:varchar(10)"`
	APIKey  string `gorm:"unique;not null"`
}

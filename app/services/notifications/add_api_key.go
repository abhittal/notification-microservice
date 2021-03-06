package notifications

import (
	"code.jtg.tools/ayush.singhal/notifications-microservice/configuration"
	"code.jtg.tools/ayush.singhal/notifications-microservice/constants"
	"code.jtg.tools/ayush.singhal/notifications-microservice/db"
	"code.jtg.tools/ayush.singhal/notifications-microservice/db/models"
	"code.jtg.tools/ayush.singhal/notifications-microservice/shared/hash"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// GetAPIKey creates a new API Key and returns it
func GetAPIKey() (string, error) {
	var organisation models.Organisation
	err := db.Get().First(&organisation).Error
	apiKey := hash.GenerateSecureToken(constants.APIKeyLength)
	apiLast := apiKey[len(apiKey)-8:]
	if err != gorm.ErrRecordNotFound && err != nil {
		return "", errors.Wrap(err, "Get API Key error")
	} else if err != gorm.ErrRecordNotFound {
		organisation.APIKey, err = hash.Message(apiKey, configuration.GetResp().APIHash)
		if err != nil {
			return "", errors.Wrap(err, "Hashing the Key error")
		}
		organisation.APILast = apiLast
		err = db.Get().Save(&organisation).Error
		if err != nil {
			return "", errors.Wrap(err, "Updating API Key error")
		}
		return apiKey, nil
	}

	organisation = models.Organisation{}

	organisation.APIKey, err = hash.Message(apiKey, configuration.GetResp().APIHash)
	if err != nil {
		return "", errors.Wrap(err, "Hashing the Key error")
	}
	organisation.APILast = apiLast
	err = db.Get().Create(&organisation).Error
	if err != nil {
		return "", errors.Wrap(err, "Creating new API Key error")
	}
	return apiKey, nil
}

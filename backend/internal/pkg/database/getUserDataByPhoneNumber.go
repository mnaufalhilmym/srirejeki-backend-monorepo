package database

import (
	"greenhouse-monitoring-iot/internal/pkg/database/model"
	"greenhouse-monitoring-iot/internal/pkg/errorHandler"

	"gorm.io/gorm"
)

func GetUserDataByPhoneNumber(pgDB *gorm.DB, phoneNumber *string) (*model.User, error) {
	user := model.User{}
	result := pgDB.Where("phone_number = ?", phoneNumber).Take(&user)
	if result.Error != nil {
		errorHandler.LogErrorThenContinue("helpers/GetUserDataByPhoneNumber1", result.Error)
		return nil, result.Error
	}
	return &user, nil
}

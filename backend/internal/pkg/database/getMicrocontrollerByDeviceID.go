package database

import (
	"greenhouse-monitoring-iot/internal/pkg/database/model"
	"greenhouse-monitoring-iot/internal/pkg/errorHandler"

	"gorm.io/gorm"
)

func GetMicrocontrollerByDeviceID(pgDB *gorm.DB, DeviceID *string) (*model.Microcontroller, error) {
	microcontroller := model.Microcontroller{}
	result := pgDB.Where("device_id = ?", DeviceID).Take(&microcontroller)
	if result.Error != nil {
		errorHandler.LogErrorThenContinue("getMicrocontrollerByDeviceID1", result.Error)
		return nil, result.Error
	}
	return &microcontroller, nil
}

package database

import (
	"greenhouse-monitoring-iot/internal/pkg/database/model"
	"greenhouse-monitoring-iot/internal/pkg/errorHandler"

	"gorm.io/gorm"
)

func GetFarmlandByID(pgDB *gorm.DB, ID *uint) (*model.Farmland, error) {
	farmland := model.Farmland{}
	result := pgDB.Where("id = ?", ID).Take(&farmland)
	if result.Error != nil {
		errorHandler.LogErrorThenContinue("getFarmlandByID1", result.Error)
		return nil, result.Error
	}
	return &farmland, nil
}

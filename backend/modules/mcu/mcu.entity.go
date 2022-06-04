package mcu

import (
	"greenhouse-monitoring-iot/internal/pkg/database"
	"greenhouse-monitoring-iot/internal/pkg/database/model"
	"greenhouse-monitoring-iot/internal/pkg/errorHandler"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Microcontroller struct {
	FiberRouter fiber.Router
	PgDB        *gorm.DB
	RDB         *redis.Client
}

func (m *Microcontroller) getMicrocontrollersByUserID(userID *uint) (*[]model.Microcontroller, error) {
	microcontrollers := []model.Microcontroller{}
	result := m.PgDB.Where("user_id = ?", userID).Find(&microcontrollers)
	if result.Error != nil {
		errorHandler.LogErrorThenContinue("getMicrocontrollersByUserID1", result.Error)
		return nil, result.Error
	}
	return &microcontrollers, nil
}

func (m *Microcontroller) getMicrocontrollersByFarmlandID(ID *uint) (*[]model.Microcontroller, error) {
	microcontroller := []model.Microcontroller{}
	result := m.PgDB.Where("farmland_id = ?", ID).Find(&microcontroller)
	if result.Error != nil {
		errorHandler.LogErrorThenContinue("getMicrocontrollersByFarmlandID1", result.Error)
		return nil, result.Error
	}
	return &microcontroller, nil
}

func (m *Microcontroller) getMicrocontrollerByID(ID *uint) (*model.Microcontroller, error) {
	microcontroller := model.Microcontroller{}
	result := m.PgDB.Where("id = ?", ID).Take(&microcontroller)
	if result.Error != nil {
		errorHandler.LogErrorThenContinue("getMicrocontrollerByID1", result.Error)
		return nil, result.Error
	}
	return &microcontroller, nil
}

func (m *Microcontroller) getMicrocontrollerByDeviceID(DeviceID *string) (*model.Microcontroller, error) {
	microcontroller, err := database.GetMicrocontrollerByDeviceID(m.PgDB, DeviceID)
	if err != nil {
		errorHandler.LogErrorThenContinue("getMicrocontrollerByDeviceID1", err)
		return nil, err
	}
	return microcontroller, nil
}

func (m *Microcontroller) getFarmlandByID(ID *uint) (*model.Farmland, error) {
	farmland, err := database.GetFarmlandByID(m.PgDB, ID)
	if err != nil {
		errorHandler.LogErrorThenContinue("getFarmlandByID1", err)
		return nil, err
	}
	return farmland, nil
}

func (m *Microcontroller) createMicrocontroller(microcontroller *model.Microcontroller) error {
	result := m.PgDB.Create(&microcontroller)
	if result.Error != nil {
		errorHandler.LogErrorThenContinue("createMicrocontroller1", result.Error)
		return result.Error
	}
	return nil
}

func (m *Microcontroller) updateMicrocontrollerByID(ID *uint, toBeUpdatedMicrocontroller *model.Microcontroller) error {
	result := m.PgDB.Model(&model.Microcontroller{}).Where("id = ?", ID).Updates(&toBeUpdatedMicrocontroller)
	if result.Error != nil {
		errorHandler.LogErrorThenContinue("updateMicrocontrollerByID1", result.Error)
		return result.Error
	}
	return nil
}

func (m *Microcontroller) deleteMicrocontrollerByID(ID *uint) error {
	result := m.PgDB.Delete(&model.Microcontroller{}, ID)
	if result.Error != nil {
		errorHandler.LogErrorThenContinue("deleteMicrocontrollerByID1", result.Error)
		return result.Error
	}
	return nil
}

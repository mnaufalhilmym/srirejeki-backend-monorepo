package farmland

import (
	"greenhouse-monitoring-iot/internal/pkg/database"
	"greenhouse-monitoring-iot/internal/pkg/database/model"
	"greenhouse-monitoring-iot/internal/pkg/errorHandler"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Farmland struct {
	FiberRouter fiber.Router
	PgDB        *gorm.DB
	RDB         *redis.Client
}

func (f *Farmland) getFarmlandsByUserID(userID *uint) (*[]model.Farmland, error) {
	farmlands := []model.Farmland{}
	result := f.PgDB.Where("user_id = ?", userID).Find(&farmlands)
	if result.Error != nil {
		errorHandler.LogErrorThenContinue("getFarmlandsByUserID1", result.Error)
		return nil, result.Error
	}
	return &farmlands, nil
}

func (f *Farmland) getFarmlandByID(ID *uint) (*model.Farmland, error) {
	farmland, err := database.GetFarmlandByID(f.PgDB, ID)
	if err != nil {
		errorHandler.LogErrorThenContinue("getFarmlandByID1", err)
		return nil, err
	}
	return farmland, nil
}

func (f *Farmland) createFarmland(farmland *model.Farmland) error {
	result := f.PgDB.Create(&farmland)
	if result.Error != nil {
		errorHandler.LogErrorThenContinue("createFarmland1", result.Error)
		return result.Error
	}
	return nil
}

func (f *Farmland) updateFarmlandByID(ID *uint, toBeUpdatedFarmland *model.Farmland, UserID *uint) error {
	result := f.PgDB.Model(&model.Farmland{}).Where("id = ?", ID).Updates(&toBeUpdatedFarmland)
	if result.Error != nil {
		errorHandler.LogErrorThenContinue("updateFarmlandByID1", result.Error)
		return result.Error
	}
	return nil
}

func (f *Farmland) deleteFarmlandByID(ID *uint) error {
	result := f.PgDB.Delete(&model.Farmland{}, ID)
	if result.Error != nil {
		errorHandler.LogErrorThenContinue("deleteFarmlandByID1", result.Error)
		return result.Error
	}
	return nil
}

func (f *Farmland) deleteMicrocontrollerByFarmlandID(farmlandID *uint) error {
	result := f.PgDB.Where("farmland_id = ?", farmlandID).Delete(&model.Microcontroller{})
	if result.Error != nil {
		errorHandler.LogErrorThenContinue("deleteMicrocontrollerByFarmlandID1", result.Error)
		return result.Error
	}
	return nil
}

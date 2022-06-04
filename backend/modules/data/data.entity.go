package data

import (
	"greenhouse-monitoring-iot/internal/pkg/database"
	"greenhouse-monitoring-iot/internal/pkg/database/model"
	"greenhouse-monitoring-iot/internal/pkg/errorHandler"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Data struct {
	FiberRouter fiber.Router
	PgDB        *gorm.DB
	RDB         *redis.Client
}

func (d *Data) getMicrocontrollerByDeviceID(deviceID *string) (*model.Microcontroller, error) {
	microcontroller, err := database.GetMicrocontrollerByDeviceID(d.PgDB, deviceID)
	if err != nil {
		errorHandler.LogErrorThenContinue("getMicrocontrollerByDeviceID1", err)
		return nil, err
	}
	return microcontroller, nil
}

func (d *Data) getSnapshotsByDeviceIDAndDurationAndTypeWithLimit(deviceID *string, duration *string, dataType *string, limit *int) (*[]model.Snapshot, error) {
	snapshots := []model.Snapshot{}
	result := d.PgDB.Limit(*limit).Order("id desc").Select("created_at", "data").Where("device_id = ? AND duration = ? AND type = ?", deviceID, duration, dataType).Find(&snapshots)
	if result.Error != nil {
		errorHandler.LogErrorThenContinue("getSnapshotsByTypeAndDurationAndDeviceID1", result.Error)
		return nil, result.Error
	}
	return &snapshots, nil
}

func (d *Data) createSnapshot(snapshot *model.Snapshot) error {
	result := d.PgDB.Create(&snapshot)
	if result.Error != nil {
		errorHandler.LogErrorThenContinue("createSnapshot1", result.Error)
		return result.Error
	}
	return nil
}

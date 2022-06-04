package profile

import (
	"encoding/json"
	"greenhouse-monitoring-iot/internal/pkg/database"
	"greenhouse-monitoring-iot/internal/pkg/database/model"
	"greenhouse-monitoring-iot/internal/pkg/errorHandler"
	"time"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Profile struct {
	FiberRouter fiber.Router
	PgDB        *gorm.DB
	RDB         *redis.Client
}

func (p *Profile) getUserDataByPhoneNumber(phoneNumber *string) (*model.User, error) {
	user, err := database.GetUserDataByPhoneNumber(p.PgDB, phoneNumber)
	if err != nil {
		errorHandler.LogErrorThenContinue("getUserDataByPhoneNumber1", err)
		return nil, err

	}
	return user, nil
}

func (p *Profile) updateUserDataByPhoneNumber(phoneNumber *string, toBeUpdatedUser *model.User) error {
	result := p.PgDB.Model(&model.User{}).Where("phone_number = ?", &phoneNumber).Updates(&toBeUpdatedUser)
	if result.Error != nil {
		errorHandler.LogErrorThenContinue("updateUserDataByPhoneNumber2", result.Error)
		return result.Error
	}
	return nil
}

func (p *Profile) updateUserDataByPhoneNumberSession(authBearerToken *string, userSession *model.UserSession, toBeUpdatedUserSession *model.UserSession) error {
	val, err := json.Marshal(&toBeUpdatedUserSession)
	if err != nil {
		errorHandler.LogErrorThenContinue("updateUserDataByPhoneNumber3", err)
		return err
	}
	if err := p.RDB.Set("user-"+*authBearerToken, val, 3*24*time.Hour).Err(); err != nil {
		errorHandler.LogErrorThenContinue("updateUserDataByPhoneNumber4", err)
		return err
	}
	return nil
}

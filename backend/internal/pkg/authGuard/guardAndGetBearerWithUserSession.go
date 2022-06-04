package authguard

import (
	"encoding/json"
	"errors"
	"greenhouse-monitoring-iot/internal/pkg/database/model"
	"greenhouse-monitoring-iot/internal/pkg/errorHandler"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

func GuardAndGetBearerWithUserSession(c *fiber.Ctx, rDB *redis.Client) (*string, *model.UserSession, error) {
	auth := c.GetReqHeaders()["Authorization"]
	if len(auth) == 0 || auth[:6] != "Bearer" {
		err := errors.New("no authorization bearer header")
		errorHandler.LogErrorThenContinue("authguard/GuardAndGetUserSession1", err)
		return nil, nil, err
	}
	bearer := auth[7:]
	var userSession *model.UserSession
	val, err := rDB.Get("user-" + bearer).Bytes()
	if err != nil {
		if err == redis.Nil {
			err := errors.New("user session not exist")
			errorHandler.LogErrorThenContinue("authGuard/GuardAndGetUserSession2", err)
			return nil, nil, err
		}
		errorHandler.LogErrorThenContinue("authGuard/GuardAndGetUserSession2", err)
		return nil, nil, err
	}
	if err := json.Unmarshal(val, &userSession); err != nil {
		errorHandler.LogErrorThenContinue("authGuard/GuardAndGetUserSession3", err)
		return nil, nil, err
	}
	return &bearer, userSession, nil
}

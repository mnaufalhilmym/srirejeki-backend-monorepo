package authguard

import (
	"errors"
	"greenhouse-monitoring-iot/internal/pkg/errorHandler"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

func Guard(c *fiber.Ctx, rDB *redis.Client) error {
	auth := c.GetReqHeaders()["Authorization"]
	if len(auth) == 0 || auth[:6] != "Bearer" {
		err := errors.New("no authorization bearer header")
		errorHandler.LogErrorThenContinue("authguard/Guard1", err)
		return err
	}
	bearer := auth[7:]
	val, err := rDB.Exists("user-" + bearer).Result()
	if err != nil {
		errorHandler.LogErrorThenContinue("authguard/Guard2", err)
		return err
	}
	if val == 0 {
		err := errors.New("user session not exist")
		errorHandler.LogErrorThenContinue("authguard/Guard3", err)
		return err
	}
	return nil
}

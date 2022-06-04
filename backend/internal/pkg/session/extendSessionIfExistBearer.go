package session

import (
	"encoding/json"
	authguard "greenhouse-monitoring-iot/internal/pkg/authGuard"
	"greenhouse-monitoring-iot/internal/pkg/errorHandler"
	"time"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

type response struct {
	Payload interface{} `json:"payload"`
}

func ExtendSessionIfExistBearer(rDB *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if len(c.GetReqHeaders()["Authorization"]) == 0 {
			return c.Next()
		}
		authBearerToken, userSession, err := authguard.GuardAndGetBearerWithUserSession(c, rDB)
		if err != nil {
			errorHandler.LogErrorThenContinue("session/ExtendSessionIfExistBearer", err)
			return c.Status(fiber.StatusUnauthorized).JSON(&response{Payload: err.Error()})
		}
		val, err := json.Marshal(userSession)
		if err != nil {
			errorHandler.LogErrorThenContinue("createUserResetPassword1", err)
			return c.Status(fiber.StatusInternalServerError).JSON(&response{Payload: err.Error()})
		}
		if err := rDB.Set("user-"+*authBearerToken, val, 3*24*time.Hour).Err(); err != nil {
			errorHandler.LogErrorThenContinue("createUserResetPassword2", err)
			return c.Status(fiber.StatusInternalServerError).JSON(&response{Payload: err.Error()})
		}
		return c.Next()
	}
}

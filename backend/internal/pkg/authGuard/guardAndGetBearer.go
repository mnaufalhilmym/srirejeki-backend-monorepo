package authguard

import (
	"errors"
	"greenhouse-monitoring-iot/internal/pkg/errorHandler"

	"github.com/gofiber/fiber/v2"
)

func GuardAndGetBearer(c *fiber.Ctx) (*string, error) {
	auth := c.GetReqHeaders()["Authorization"]
	if len(auth) == 0 || auth[:6] != "Bearer" {
		err := errors.New("no authorization bearer header")
		errorHandler.LogErrorThenContinue("authguard/GuardAndGetUserSession1", err)
		return nil, err
	}
	bearer := auth[7:]
	return &bearer, nil
}

package data

import (
	authguard "greenhouse-monitoring-iot/internal/pkg/authGuard"
	"greenhouse-monitoring-iot/internal/pkg/errorHandler"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
)

// getSubscribe godoc
// @Summary      Get mqtt data in specified topic
// @Description  Get mqtt data in specified topic using authorization bearer header
// @Tags         Data
// @Produce      json
// @Param        topic    query     string  true  "Mqtt topic"
// @Success      200      {object}  dataResponse
// @Failure      400      {object}  dataResponse
// @Failure      401      {object}  dataResponse
// @Failure      500      {object}  dataResponse
// @Router       /data/subscribe [get]
func getSubscribe(d *Data) {
	d.FiberRouter.Get("/subscribe", timeout.New(func(c *fiber.Ctx) error {
		userSession, err := authguard.GuardAndGetUserSession(c, d.RDB)
		if err != nil {
			errorHandler.LogErrorThenContinue("getSubscribe1", err)
			return c.Status(fiber.StatusUnauthorized).JSON(&dataResponse{Payload: err.Error()})
		}
		dataSubscribeRequest := new(dataSubscribeRequest)
		if err := c.QueryParser(dataSubscribeRequest); err != nil {
			errorHandler.LogErrorThenContinue("getSubscribe2", err)
			return c.Status(fiber.StatusBadRequest).JSON(&dataResponse{Payload: err.Error()})
		}
		validator := validator.New()
		if err := validator.Struct(dataSubscribeRequest); err != nil {
			errorHandler.LogErrorThenContinue("getSubscribe3", err)
			return c.Status(fiber.StatusInternalServerError).JSON(&dataResponse{Payload: err.Error()})
		}
		result := d.getSubscribeService(userSession, dataSubscribeRequest)
		return c.Status(result.code).JSON(&dataResponse{Payload: result.payload})
	}, 10*time.Second))
}

// getSnapshot godoc
// @Summary      Get microcontroller snapshot data
// @Description  Get microcontroller snapshot data using authorization bearer header
// @Tags         Data
// @Produce      json
// @Param        type     query     string  true  "Type of microcontroller sensor"
// @Param        deviceId query     string  true  "deviceId/clientId of the microcontroller"
// @Param        duration query     string  false  "Duration of snapshot data: hour | day | month. Default: month"
// @Param        limit    query     string  false  "Limit of snapshot data. It must be string form of a number. Default: 30"
// @Success      200      {object}  dataResponse
// @Failure      400      {object}  dataResponse
// @Failure      401      {object}  dataResponse
// @Failure      500      {object}  dataResponse
// @Router       /data/snapshot [get]
func getSnapshot(d *Data) {
	d.FiberRouter.Get("/snapshot", func(c *fiber.Ctx) error {
		userSession, err := authguard.GuardAndGetUserSession(c, d.RDB)
		if err != nil {
			errorHandler.LogErrorThenContinue("getSnapshot1", err)
			return c.Status(fiber.StatusUnauthorized).JSON(&dataResponse{Payload: err.Error()})
		}
		dataSnapshotGetRequest := new(dataSnapshotGetRequest)
		if err = c.QueryParser(dataSnapshotGetRequest); err != nil {
			errorHandler.LogErrorThenContinue("getSnapshot2", err)
			return c.Status(fiber.StatusBadRequest).JSON(&dataResponse{Payload: err.Error()})
		}
		validator := validator.New()
		if err := validator.Struct(dataSnapshotGetRequest); err != nil {
			errorHandler.LogErrorThenContinue("getSnapshot3", err)
			return c.Status(fiber.StatusInternalServerError).JSON(&dataResponse{Payload: err.Error()})
		}
		result := d.getSnapshotSevice(userSession, dataSnapshotGetRequest)
		return c.Status(result.code).JSON(&dataResponse{Payload: result.payload})
	})
}

// postSnapshot godoc
// @Summary      Post microcontroller snapshot data | INTERNAL USE ON BACKEND ONLY
// @Description  Post microcontroller snapshot data | INTERNAL USE ON BACKEND ONLY
// @Tags         Data
// @Accept       json
// @Produce      json
// @Param        dataSnapshotPostRequest  body      dataSnapshotPostRequest  true  "Required data to post microcontroller snapshot data"
// @Success      200      {object}  dataResponse
// @Failure      400      {object}  dataResponse
// @Failure      500      {object}  dataResponse
// @Router       /data/snapshot [post]
func postSnapshot(d *Data) {
	d.FiberRouter.Post("/snapshot", func(c *fiber.Ctx) error {
		var dataSnapshotPostRequest *dataSnapshotPostRequest
		if err := c.BodyParser(&dataSnapshotPostRequest); err != nil {
			errorHandler.LogErrorThenContinue("postSnapshot1", err)
			return c.Status(fiber.StatusBadRequest).JSON(&dataResponse{Payload: err.Error()})
		}
		validator := validator.New()
		if err := validator.Struct(dataSnapshotPostRequest); err != nil {
			errorHandler.LogErrorThenContinue("postSnapshot2", err)
			return c.Status(fiber.StatusInternalServerError).JSON(&dataResponse{Payload: err.Error()})
		}
		result := d.postSnapshotService(dataSnapshotPostRequest)
		return c.Status(result.code).JSON(&dataResponse{Payload: result.payload})
	})
}

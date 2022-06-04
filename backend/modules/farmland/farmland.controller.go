package farmland

import (
	authguard "greenhouse-monitoring-iot/internal/pkg/authGuard"
	"greenhouse-monitoring-iot/internal/pkg/errorHandler"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// getFarmlands godoc
// @Summary      Get all user's farmlands data
// @Description  Get all user's farmlands data using authorization bearer header
// @Tags         Farmland
// @Produce      json
// @Success      200      {object}  farmlandResponse
// @Failure      401      {object}  farmlandResponse
// @Failure      500      {object}  farmlandResponse
// @Router       /farmland/user [get]
func getFarmlands(f *Farmland) {
	f.FiberRouter.Get("/user", func(c *fiber.Ctx) error {
		userSession, err := authguard.GuardAndGetUserSession(c, f.RDB)
		if err != nil {
			errorHandler.LogErrorThenContinue("getFarmlands1", err)
			return c.Status(fiber.StatusUnauthorized).JSON(&farmlandResponse{Payload: err.Error()})
		}
		result := f.getFarmlandsService(userSession)
		return c.Status(result.code).JSON(&farmlandResponse{Payload: result.payload})
	})
}

// getFarmland godoc
// @Summary      Get one farmlands data
// @Description  Get one farmlands data using authorization bearer header
// @Tags         Farmland
// @Produce      json
// @Param        id    query     string  true  "Farmland Id"
// @Success      200      {object}  farmlandResponse
// @Failure      400      {object}  farmlandResponse
// @Failure      401      {object}  farmlandResponse
// @Failure      500      {object}  farmlandResponse
// @Router       /farmland [get]
func getFarmland(f *Farmland) {
	f.FiberRouter.Get("", func(c *fiber.Ctx) error {
		if err := authguard.Guard(c, f.RDB); err != nil {
			errorHandler.LogErrorThenContinue("getFarmland1", err)
			return c.Status(fiber.StatusUnauthorized).JSON(&farmlandResponse{Payload: err.Error()})
		}
		farmlandGetRequest := new(farmlandGetRequest)
		if err := c.QueryParser(farmlandGetRequest); err != nil {
			errorHandler.LogErrorThenContinue("getFarmland2", err)
			return c.Status(fiber.StatusBadRequest).JSON(&farmlandResponse{Payload: err.Error()})
		}
		validator := validator.New()
		if err := validator.Struct(farmlandGetRequest); err != nil {
			errorHandler.LogErrorThenContinue("getFarmland3", err)
			return c.Status(fiber.StatusInternalServerError).JSON(&farmlandResponse{Payload: err.Error()})
		}
		result := f.getFarmlandService(farmlandGetRequest)
		return c.Status(result.code).JSON(&farmlandResponse{Payload: result.payload})
	})
}

// postFarmland godoc
// @Summary      Add new farmland data
// @Description  Add new farmland data using authorization bearer header
// @Tags         Farmland
// @Accept       json
// @Produce      json
// @Param        farmlandAddRequest  body      farmlandAddRequest  true  "Required data to add new farmland"
// @Success      201      {object}  farmlandResponse
// @Failure      400      {object}  farmlandResponse
// @Failure      401      {object}  farmlandResponse
// @Failure      500      {object}  farmlandResponse
// @Router       /farmland [post]
func postFarmland(f *Farmland) {
	f.FiberRouter.Post("", func(c *fiber.Ctx) error {
		userSession, err := authguard.GuardAndGetUserSession(c, f.RDB)
		if err != nil {
			errorHandler.LogErrorThenContinue("postFarmland1", err)
			return c.Status(fiber.StatusUnauthorized).JSON(&farmlandResponse{Payload: err.Error()})
		}
		var farmlandAddRequest *farmlandAddRequest
		if err := c.BodyParser(&farmlandAddRequest); err != nil {
			errorHandler.LogErrorThenContinue("postFarmland2", err)
			return c.Status(fiber.StatusBadRequest).JSON(&farmlandResponse{Payload: err.Error()})
		}
		validator := validator.New()
		if err := validator.Struct(farmlandAddRequest); err != nil {
			errorHandler.LogErrorThenContinue("postFarmland3", err)
			return c.Status(fiber.StatusInternalServerError).JSON(&farmlandResponse{Payload: err.Error()})
		}
		result := f.postFarmlandService(userSession, farmlandAddRequest)
		return c.Status(result.code).JSON(&farmlandResponse{Payload: result.payload})
	})
}

// patchFarmland godoc
// @Summary      Edit existing farmland data
// @Description  Edit existing farmland data using authorization bearer header
// @Tags         Farmland
// @Accept       json
// @Produce      json
// @Param        farmlandEditRequest  body      farmlandEditRequest  true  "Required data to edit existing farmland"
// @Success      200      {object}  farmlandResponse
// @Failure      400      {object}  farmlandResponse
// @Failure      401      {object}  farmlandResponse
// @Failure      500      {object}  farmlandResponse
// @Router       /farmland [patch]
func patchFarmland(f *Farmland) {
	f.FiberRouter.Patch("", func(c *fiber.Ctx) error {
		userSession, err := authguard.GuardAndGetUserSession(c, f.RDB)
		if err != nil {
			errorHandler.LogErrorThenContinue("patchFarmland1", err)
			return c.Status(fiber.StatusUnauthorized).JSON(&farmlandResponse{Payload: err.Error()})
		}
		var farmlandEditRequest *farmlandEditRequest
		if err := c.BodyParser(&farmlandEditRequest); err != nil {
			errorHandler.LogErrorThenContinue("patchFarmland2", err)
			return c.Status(fiber.StatusBadRequest).JSON(&farmlandResponse{Payload: err.Error()})
		}
		validator := validator.New()
		if err := validator.Struct(farmlandEditRequest); err != nil {
			errorHandler.LogErrorThenContinue("patchFarmland3", err)
			return c.Status(fiber.StatusInternalServerError).JSON(&farmlandResponse{Payload: err.Error()})
		}
		result := f.patchFarmlandService(userSession, farmlandEditRequest)
		return c.Status(result.code).JSON(&farmlandResponse{Payload: result.payload})
	})
}

// deleteFarmland godoc
// @Summary      Delete existing farmland data
// @Description  Delete existing farmland data using authorization bearer header
// @Tags         Farmland
// @Accept       json
// @Produce      json
// @Param        farmlandDeleteRequest  body      farmlandDeleteRequest  true  "Required data to delete existing farmland"
// @Success      200      {object}  farmlandResponse
// @Failure      401      {object}  farmlandResponse
// @Failure      500      {object}  farmlandResponse
// @Router       /farmland [delete]
func deleteFarmland(f *Farmland) {
	f.FiberRouter.Delete("", func(c *fiber.Ctx) error {
		if err := authguard.Guard(c, f.RDB); err != nil {
			errorHandler.LogErrorThenContinue("deleteFarmland1", err)
			return c.Status(fiber.StatusUnauthorized).JSON(&farmlandResponse{Payload: err.Error()})
		}
		var farmlandDeleteRequest *farmlandDeleteRequest
		if err := c.BodyParser(&farmlandDeleteRequest); err != nil {
			errorHandler.LogErrorThenContinue("deleteFarmland2", err)
			return c.Status(fiber.StatusUnauthorized).JSON(&farmlandResponse{Payload: err.Error()})
		}
		validator := validator.New()
		if err := validator.Struct(farmlandDeleteRequest); err != nil {
			errorHandler.LogErrorThenContinue("deleteFarmland3", err)
			return c.Status(fiber.StatusInternalServerError).JSON(&farmlandResponse{Payload: err.Error()})
		}
		result := f.deleteFarmlandService(farmlandDeleteRequest)
		return c.Status(result.code).JSON(&farmlandResponse{Payload: result.payload})
	})
}

package mcu

import (
	authguard "greenhouse-monitoring-iot/internal/pkg/authGuard"
	"greenhouse-monitoring-iot/internal/pkg/errorHandler"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// getAllMicrocontrollers godoc
// @Summary      Get all user's microcontrollers data
// @Description  Get all user's microcontrollers data using authorization bearer header
// @Tags         Microcontroller
// @Produce      json
// @Success      200      {object}  microcontrollerResponse
// @Failure      401      {object}  microcontrollerResponse
// @Failure      500      {object}  microcontrollerResponse
// @Router       /mcu/user [get]
func getAllMicrocontrollers(m *Microcontroller) {
	m.FiberRouter.Get("/user", func(c *fiber.Ctx) error {
		userSession, err := authguard.GuardAndGetUserSession(c, m.RDB)
		if err != nil {
			errorHandler.LogErrorThenContinue("getAllMicrocontrollers1", err)
			return c.Status(fiber.StatusUnauthorized).JSON(&microcontrollerResponse{Payload: err.Error()})
		}
		result := m.getAllMicrocontrollersService(userSession)
		return c.Status(result.code).JSON(&microcontrollerResponse{Payload: result.payload})
	})
}

// getFarmlandMicrocontrollers godoc
// @Summary      Get all farmland's microcontrollers data
// @Description  Get all farmland's microcontrollers data using authorization bearer header
// @Tags         Microcontroller
// @Produce      json
// @Param        id    query     string  true  "Farmland Id"
// @Success      200      {object}  microcontrollerResponse
// @Failure      400      {object}  microcontrollerResponse
// @Failure      401      {object}  microcontrollerResponse
// @Failure      500      {object}  microcontrollerResponse
// @Router       /mcu/farmland [get]
func getFarmlandMicrocontrollers(m *Microcontroller) {
	m.FiberRouter.Get("/farmland", func(c *fiber.Ctx) error {
		if err := authguard.Guard(c, m.RDB); err != nil {
			errorHandler.LogErrorThenContinue("getFarmlandMicrocontrollers1", err)
			return c.Status(fiber.StatusUnauthorized).JSON(&microcontrollerResponse{Payload: err.Error()})
		}
		microcontrollersFarmlandGetRequest := new(microcontrollersFarmlandGetRequest)
		if err := c.QueryParser(microcontrollersFarmlandGetRequest); err != nil {
			errorHandler.LogErrorThenContinue("getFarmlandMicrocontrollers2", err)
			return c.Status(fiber.StatusBadRequest).JSON(&microcontrollerResponse{Payload: err.Error()})
		}
		validator := validator.New()
		if err := validator.Struct(microcontrollersFarmlandGetRequest); err != nil {
			errorHandler.LogErrorThenContinue("getFarmlandMicrocontrollers3", err)
			return c.Status(fiber.StatusInternalServerError).JSON(&microcontrollerResponse{Payload: err.Error()})
		}
		result := m.getFarmlandMicrocontrollersService(microcontrollersFarmlandGetRequest)
		return c.Status(result.code).JSON(&microcontrollerResponse{Payload: result.payload})
	})
}

// getMicrocontroller godoc
// @Summary      Get one microcontroller data
// @Description  Get one microcontroller data using authorization bearer header
// @Tags         Microcontroller
// @Produce      json
// @Param        id    query     string  true  "Microcontroller Id"
// @Success      200      {object}  microcontrollerResponse
// @Failure      400      {object}  microcontrollerResponse
// @Failure      401      {object}  microcontrollerResponse
// @Failure      500      {object}  microcontrollerResponse
// @Router       /mcu [get]
func getMicrocontroller(m *Microcontroller) {
	m.FiberRouter.Get("", func(c *fiber.Ctx) error {
		if err := authguard.Guard(c, m.RDB); err != nil {
			errorHandler.LogErrorThenContinue("getMicrocontroller1", err)
			return c.Status(fiber.StatusUnauthorized).JSON(&microcontrollerResponse{Payload: err.Error()})
		}
		microcontrollerGetRequest := new(microcontrollerGetRequest)
		if err := c.QueryParser(microcontrollerGetRequest); err != nil {
			errorHandler.LogErrorThenContinue("getMicrocontroller2", err)
			return c.Status(fiber.StatusBadRequest).JSON(&microcontrollerResponse{Payload: err.Error()})
		}
		validator := validator.New()
		if err := validator.Struct(microcontrollerGetRequest); err != nil {
			errorHandler.LogErrorThenContinue("getMicrocontroller3", err)
			return c.Status(fiber.StatusInternalServerError).JSON(&microcontrollerResponse{Payload: err.Error()})
		}
		result := m.getMicrocontrollerService(microcontrollerGetRequest)
		return c.Status(result.code).JSON(&microcontrollerResponse{Payload: result.payload})
	})
}

// postMicrocontroller godoc
// @Summary      Add new microcontroller data
// @Description  Add new microcontroller data using authorization bearer header
// @Tags         Microcontroller
// @Accept       json
// @Produce      json
// @Param        microcontrollerAddRequest  body      microcontrollerAddRequest  true  "Required data to add new microcontroller"
// @Success      201      {object}  microcontrollerResponse
// @Failure      400      {object}  microcontrollerResponse
// @Failure      401      {object}  microcontrollerResponse
// @Failure      500      {object}  microcontrollerResponse
// @Router       /mcu [post]
func postMicrocontroller(m *Microcontroller) {
	m.FiberRouter.Post("", func(c *fiber.Ctx) error {
		userSession, err := authguard.GuardAndGetUserSession(c, m.RDB)
		if err != nil {
			errorHandler.LogErrorThenContinue("postMicrocontroller1", err)
			return c.Status(fiber.StatusUnauthorized).JSON(&microcontrollerResponse{Payload: err.Error()})
		}
		var microcontrollerAddRequest *microcontrollerAddRequest
		if err := c.BodyParser(&microcontrollerAddRequest); err != nil {
			errorHandler.LogErrorThenContinue("postMicrocontroller2", err)
			return c.Status(fiber.StatusBadRequest).JSON(&microcontrollerResponse{Payload: err.Error()})
		}
		validator := validator.New()
		if err := validator.Struct(microcontrollerAddRequest); err != nil {
			errorHandler.LogErrorThenContinue("postMicrocontroller3", err)
			return c.Status(fiber.StatusInternalServerError).JSON(&microcontrollerResponse{Payload: err.Error()})
		}
		result := m.postMicrocontrollerService(userSession, microcontrollerAddRequest)
		return c.Status(result.code).JSON(&microcontrollerResponse{Payload: result.payload})
	})
}

// patchMicrocontroller godoc
// @Summary      Edit existing microcontroller data
// @Description  Edit existing microcontroller data using authorization bearer header
// @Tags         Microcontroller
// @Accept       json
// @Produce      json
// @Param        microcontrollerEditRequest  body      microcontrollerEditRequest  true  "Required data to edit existing microcontroller"
// @Success      200      {object}  microcontrollerResponse
// @Failure      400      {object}  microcontrollerResponse
// @Failure      401      {object}  microcontrollerResponse
// @Failure      500      {object}  microcontrollerResponse
// @Router       /mcu [patch]
func patchMicrocontroller(m *Microcontroller) {
	m.FiberRouter.Patch("", func(c *fiber.Ctx) error {
		if err := authguard.Guard(c, m.RDB); err != nil {
			errorHandler.LogErrorThenContinue("patchMicrocontroller1", err)
			return c.Status(fiber.StatusUnauthorized).JSON(&microcontrollerResponse{Payload: err.Error()})
		}
		var microcontrollerEditRequest *microcontrollerEditRequest
		if err := c.BodyParser(&microcontrollerEditRequest); err != nil {
			errorHandler.LogErrorThenContinue("patchMicrocontroller2", err)
			return c.Status(fiber.StatusBadRequest).JSON(&microcontrollerResponse{Payload: err.Error()})
		}
		validator := validator.New()
		if err := validator.Struct(microcontrollerEditRequest); err != nil {
			errorHandler.LogErrorThenContinue("patchMicrocontroller3", err)
			return c.Status(fiber.StatusInternalServerError).JSON(&microcontrollerResponse{Payload: err.Error()})
		}
		result := m.patchMicrocontrollerService(microcontrollerEditRequest)
		return c.Status(result.code).JSON(&microcontrollerResponse{Payload: result.payload})
	})
}

// deleteMicrocontroller godoc
// @Summary      Delete existing microcontroller data
// @Description  Delete existing microcontroller data
// @Tags         Microcontroller
// @Accept       json
// @Produce      json
// @Param        microcontrollerDeleteRequest  body      microcontrollerDeleteRequest  true  "Required data to delete existing microcontroller"
// @Success      200      {object}  microcontrollerResponse
// @Failure      401      {object}  microcontrollerResponse
// @Failure      500      {object}  microcontrollerResponse
// @Router       /mcu [delete]
func deleteMicrocontroller(m *Microcontroller) {
	m.FiberRouter.Delete("", func(c *fiber.Ctx) error {
		if err := authguard.Guard(c, m.RDB); err != nil {
			errorHandler.LogErrorThenContinue("deleteMicrocontroller1", err)
			return c.Status(fiber.StatusUnauthorized).JSON(&microcontrollerResponse{Payload: err.Error()})
		}
		var microcontrollerDeleteRequest *microcontrollerDeleteRequest
		if err := c.BodyParser(&microcontrollerDeleteRequest); err != nil {
			errorHandler.LogErrorThenContinue("deleteMicrocontroller2", err)
			return c.Status(fiber.StatusUnauthorized).JSON(&microcontrollerResponse{Payload: err.Error()})
		}
		validator := validator.New()
		if err := validator.Struct(microcontrollerDeleteRequest); err != nil {
			errorHandler.LogErrorThenContinue("deleteMicrocontroller3", err)
			return c.Status(fiber.StatusInternalServerError).JSON(&microcontrollerResponse{Payload: err.Error()})
		}
		result := m.deleteMicrocontrollerService(microcontrollerDeleteRequest)
		return c.Status(result.code).JSON(&microcontrollerResponse{Payload: result.payload})
	})
}

// authMicrocontroller godoc
// @Summary      Authenticate microcontroller | INTERNAL USE ON BACKEND ONLY
// @Description  Authenticate microcontroller | INTERNAL USE ON BACKEND ONLY
// @Tags         Microcontroller
// @Accept       json
// @Produce      json
// @Param        microcontrollerAuthRequest  body      microcontrollerAuthRequest  true  "Required data to authenticate microcontroller"
// @Success      200      {object}  microcontrollerResponse
// @Failure      400      {object}  microcontrollerResponse
// @Failure      401      {object}  microcontrollerResponse
// @Router       /mcu/auth [post]
func authMicrocontroller(m *Microcontroller) {
	m.FiberRouter.Post("/auth", func(c *fiber.Ctx) error {
		var microcontrollerAuthRequest *microcontrollerAuthRequest
		if err := c.BodyParser(&microcontrollerAuthRequest); err != nil {
			errorHandler.LogErrorThenContinue("authMicrocontroller1", err)
			return c.Status(fiber.StatusBadRequest).JSON(&microcontrollerResponse{Payload: err.Error()})
		}
		validator := validator.New()
		if err := validator.Struct(microcontrollerAuthRequest); err != nil {
			errorHandler.LogErrorThenContinue("authMicrocontroller2", err)
			return c.Status(fiber.StatusInternalServerError).JSON(&microcontrollerResponse{Payload: err.Error()})
		}
		result := m.authMicrocontrollerService(microcontrollerAuthRequest)
		return c.Status(result.code).JSON(&microcontrollerResponse{Payload: result.payload})
	})
}

// postSendDataToMcu godoc
// @Summary      Send data to microcontroller
// @Description  Send data to microcontroller using authorization bearer header
// @Tags         Microcontroller
// @Accept       json
// @Produce      json
// @Param        microcontrollerSendDataToMcuRequest  body      microcontrollerSendDataToMcuRequest  true  "Required data to send data to microcontroller"
// @Success      200      {object}  microcontrollerResponse
// @Failure      400      {object}  microcontrollerResponse
// @Failure      401      {object}  microcontrollerResponse
// @Failure      500      {object}  microcontrollerResponse
// @Router       /mcu/send [post]
func postSendDataToMcu(m *Microcontroller) {
	m.FiberRouter.Post("/send", func(c *fiber.Ctx) error {
		userSession, err := authguard.GuardAndGetUserSession(c, m.RDB)
		if err != nil {
			errorHandler.LogErrorThenContinue("forwardDataToMcu1", err)
			return c.Status(fiber.StatusUnauthorized).JSON(&microcontrollerResponse{Payload: err.Error()})
		}
		var microcontrollerSendDataToMcuRequest *microcontrollerSendDataToMcuRequest
		if err := c.BodyParser(&microcontrollerSendDataToMcuRequest); err != nil {
			errorHandler.LogErrorThenContinue("forwardDataToMcu2", err)
			return c.Status(fiber.StatusInternalServerError).JSON(&microcontrollerResponse{Payload: err.Error()})
		}
		validator := validator.New()
		if err := validator.Struct(microcontrollerSendDataToMcuRequest); err != nil {
			errorHandler.LogErrorThenContinue("forwardDataToMcu3", err)
			return c.Status(fiber.StatusInternalServerError).JSON(&microcontrollerResponse{Payload: err.Error()})
		}
		result := m.sendDataToMcuService(userSession, microcontrollerSendDataToMcuRequest)
		return c.Status(result.code).JSON(&microcontrollerResponse{Payload: result.payload})
	})
}

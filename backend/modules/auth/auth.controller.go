package auth

import (
	authguard "greenhouse-monitoring-iot/internal/pkg/authGuard"
	"greenhouse-monitoring-iot/internal/pkg/errorHandler"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// getUserSession godoc
// @Summary      Get user session data
// @Description  Get user session data using authorization bearer header
// @Tags         Auth
// @Produce      json
// @Success      200      {object}  authResponse
// @Failure      401      {object}  authResponse
// @Router       /auth [get]
func getUserSession(a *Auth) {
	a.FiberRouter.Get("", func(c *fiber.Ctx) error {
		userSession, err := authguard.GuardAndGetUserSession(c, a.RDB)
		if err != nil {
			errorHandler.LogErrorThenContinue("getUserSession1", err)
			return c.Status(fiber.StatusBadRequest).JSON(&authResponse{Payload: err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(&authResponse{Payload: &userSession})
	})
}

// postSignUp godoc
// @Summary      Sign up an account
// @Description  Sign up an account
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        authSignUpRequest  body      authSignUpRequest  true  "Required account data"
// @Success      200      {object}  authResponse
// @Failure      400      {object}  authResponse
// @Failure      500      {object}  authResponse
// @Router       /auth/signup [post]
func postSignUp(a *Auth) {
	a.FiberRouter.Post("/signup", func(c *fiber.Ctx) error {
		var authSignUpRequest *authSignUpRequest
		if err := c.BodyParser(&authSignUpRequest); err != nil {
			errorHandler.LogErrorThenContinue("postSignUp1", err)
			return c.Status(fiber.StatusBadRequest).JSON(&authResponse{Payload: err.Error()})
		}
		validator := validator.New()
		if err := validator.Struct(authSignUpRequest); err != nil {
			errorHandler.LogErrorThenContinue("postSignUp2", err)
			return c.Status(fiber.StatusBadRequest).JSON(&authResponse{Payload: err.Error()})
		}
		result := a.postSignUpService(authSignUpRequest)
		return c.Status(result.code).JSON(&authResponse{Payload: result.payload})
	})
}

// postVerifySignUp godoc
// @Summary      Verify signed up account
// @Description  Verify signed up account
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        authVerificationRequest  body      authVerificationRequest  true  "Code to verify signed up account"
// @Success      201      {object}  authResponse
// @Failure      400      {object}  authResponse
// @Failure      401      {object}  authResponse
// @Failure      403      {object}  authResponse
// @Failure      500      {object}  authResponse
// @Router       /auth/verify-signup [post]
func postVerifySignUp(a *Auth) {
	a.FiberRouter.Post("/verify-signup", func(c *fiber.Ctx) error {
		var authVerificationRequest *authVerificationRequest
		if err := c.BodyParser(&authVerificationRequest); err != nil {
			errorHandler.LogErrorThenContinue("postVerifySignUp1", err)
			return c.Status(fiber.StatusBadRequest).JSON(&authResponse{Payload: err.Error()})
		}
		validate := validator.New()
		if err := validate.Struct(authVerificationRequest); err != nil {
			errorHandler.LogErrorThenContinue("postVerifySignUp2", err)
			return c.Status(fiber.StatusInternalServerError).JSON(&authResponse{Payload: err.Error()})
		}
		result := a.postVerifySignUpService(authVerificationRequest)
		return c.Status(result.code).JSON(&authResponse{Payload: result.payload})
	})
}

// postSignIn godoc
// @Summary      Sign in
// @Description  Sign in
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        authSignInRequest  body      authSignInRequest  true  "Required sign in data"
// @Success      200      {object}  authResponse
// @Failure      400      {object}  authResponse
// @Failure      401      {object}  authResponse
// @Failure      500      {object}  authResponse
// @Router       /auth/signin [post]
func postSignIn(a *Auth) {
	a.FiberRouter.Post("/signin", func(c *fiber.Ctx) error {
		var authSignInRequest *authSignInRequest
		if err := c.BodyParser(&authSignInRequest); err != nil {
			errorHandler.LogErrorThenContinue("postSignIn1", err)
			return c.Status(fiber.StatusBadRequest).JSON(&authResponse{Payload: err.Error()})
		}
		validator := validator.New()
		if err := validator.Struct(authSignInRequest); err != nil {
			errorHandler.LogErrorThenContinue("postSignIn2", err)
			return c.Status(fiber.StatusInternalServerError).JSON(&authResponse{Payload: err.Error()})
		}
		result := a.postSignInService(authSignInRequest)
		return c.Status(result.code).JSON(&authResponse{Payload: result.payload})
	})
}

// postRequestResetPassword godoc
// @Summary      Send request reset user password
// @Description  Send request reset user password
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        authRequestResetPasswordRequest  body      authRequestResetPasswordRequest  true  "Phone number to send request reset password"
// @Success      200      {object}  authResponse
// @Failure      400      {object}  authResponse
// @Router       /auth/request-reset-password [post]
func postRequestResetPassword(a *Auth) {
	a.FiberRouter.Post("/request-reset-password", func(c *fiber.Ctx) error {
		var authRequestResetPasswordRequest *authRequestResetPasswordRequest
		if err := c.BodyParser(&authRequestResetPasswordRequest); err != nil {
			errorHandler.LogErrorThenContinue("postRequestResetPassword1", err)
			return c.Status(fiber.StatusBadRequest).JSON(&authResponse{Payload: err.Error()})
		}
		validator := validator.New()
		if err := validator.Struct(authRequestResetPasswordRequest); err != nil {
			errorHandler.LogErrorThenContinue("postRequestResetPassword2", err)
			return c.Status(fiber.StatusInternalServerError).JSON(&authResponse{Payload: err.Error()})
		}
		result := a.postRequestResetPasswordService(authRequestResetPasswordRequest)
		return c.Status(result.code).JSON(&authResponse{Payload: result.payload})
	})
}

// postVerifyRequestResetPassword godoc
// @Summary      Verify request reset password
// @Description  Verify request reset password
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        authVerificationRequest  body      authVerificationRequest  true  "Required data to verify request reset password"
// @Success      200      {object}  authResponse
// @Failure      400      {object}  authResponse
// @Router       /auth/verify-request-reset-password [post]
func postVerifyRequestResetPassword(a *Auth) {
	a.FiberRouter.Post("/verify-request-reset-password", func(c *fiber.Ctx) error {
		var authVerificationRequest *authVerificationRequest
		if err := c.BodyParser(&authVerificationRequest); err != nil {
			errorHandler.LogErrorThenContinue("postVerifyRequestResetPassword1", err)
			return c.Status(fiber.StatusBadRequest).JSON(&authResponse{Payload: err.Error()})
		}
		validator := validator.New()
		if err := validator.Struct(authVerificationRequest); err != nil {
			errorHandler.LogErrorThenContinue("postVerifyRequestResetPassword2", err)
			return c.Status(fiber.StatusInternalServerError).JSON(&authResponse{Payload: err.Error()})
		}
		result := a.postVerifyRequestResetPasswordService(authVerificationRequest)
		return c.Status(result.code).JSON(&authResponse{Payload: result.payload})
	})
}

// patchResetPassword godoc
// @Summary      Reset user's password
// @Description  Reset user's password
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        authResetPasswordRequest  body      authResetPasswordRequest  true  "Required data to reset password"
// @Success      200      {object}  authResponse
// @Failure      400      {object}  authResponse
// @Failure      401      {object}  authResponse
// @Failure      500      {object}  authResponse
// @Router       /auth/reset-password [patch]
func patchResetPassword(a *Auth) {
	a.FiberRouter.Patch("/reset-password", func(c *fiber.Ctx) error {
		var authResetPasswordRequest *authResetPasswordRequest
		if err := c.BodyParser(&authResetPasswordRequest); err != nil {
			errorHandler.LogErrorThenContinue("patchResetPassword1", err)
			return c.Status(fiber.StatusBadRequest).JSON(&authResponse{Payload: err.Error()})
		}
		validator := validator.New()
		if err := validator.Struct(authResetPasswordRequest); err != nil {
			errorHandler.LogErrorThenContinue("patchResetPassword2", err)
			return c.Status(fiber.StatusInternalServerError).JSON(&authResponse{Payload: err.Error()})
		}
		result := a.patchResetPasswordService(authResetPasswordRequest)
		return c.Status(result.code).JSON(&authResponse{Payload: result.payload})
	})
}

// getSignOut godoc
// @Summary      Sign out
// @Description  Sign out using authorization bearer header
// @Tags         Auth
// @Produce      json
// @Success      200      {object}  authResponse
// @Router       /auth/signout [get]
func getSignOut(a *Auth) {
	a.FiberRouter.Get("/signout", func(c *fiber.Ctx) error {
		bearer, err := authguard.GuardAndGetBearer(c)
		if err != nil {
			errorHandler.LogErrorThenContinue("getSignOut1", err)
			// Pass
		}
		result := a.getSignOutService(*bearer)
		return c.Status(result.code).JSON(&authResponse{Payload: result.payload})
	})
}

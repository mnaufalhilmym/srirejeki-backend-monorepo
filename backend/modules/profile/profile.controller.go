package profile

import (
	"errors"
	authguard "greenhouse-monitoring-iot/internal/pkg/authGuard"
	"greenhouse-monitoring-iot/internal/pkg/errorHandler"

	"github.com/gofiber/fiber/v2"
)

// getProfile    godoc
// @Summary      Get profile data
// @Description  Get profile data using authorization Bearer header
// @Tags         Profile
// @Produce      json
// @Success      200      {object}  profileResponse
// @Failure      401      {object}  profileResponse
// @Failure      500      {object}  profileResponse
// @Router       /profile [get]
func getProfile(p *Profile) {
	p.FiberRouter.Get("", func(c *fiber.Ctx) error {
		userSession, err := authguard.GuardAndGetUserSession(c, p.RDB)
		if err != nil {
			errorHandler.LogErrorThenContinue("getProfile1", err)
			return c.Status(fiber.StatusUnauthorized).JSON(&profileResponse{Payload: err.Error()})
		}
		result := p.getProfileService(userSession)
		return c.Status(result.code).JSON(&profileResponse{Payload: result.payload})
	})

}

// editProfile godoc
// @Summary      Edit profile data
// @Description  Edit profile data using authorization Bearer header
// @Tags         Profile
// @Accept       json
// @Produce      json
// @Param        profileEditRequest  body      profileEditRequest  true  "Required data to edit profile"
// @Success      200      {object}  profileResponse
// @Failure      400      {object}  profileResponse
// @Failure      401      {object}  profileResponse
// @Failure      500      {object}  profileResponse
// @Router       /profile [patch]
func editProfile(p *Profile) {
	p.FiberRouter.Patch("", func(c *fiber.Ctx) error {
		bearer, userSession, err := authguard.GuardAndGetBearerWithUserSession(c, p.RDB)
		if err != nil {
			errorHandler.LogErrorThenContinue("patchProfile1", err)
			return c.Status(fiber.StatusUnauthorized).JSON(&profileResponse{Payload: err.Error()})
		}
		var profileRequestPatch *profileEditRequest
		if err := c.BodyParser(&profileRequestPatch); err != nil {
			errorHandler.LogErrorThenContinue("patchProfile2", err)
			return c.Status(fiber.StatusBadRequest).JSON(&profileResponse{Payload: err.Error()})
		}
		if profileRequestPatch == nil {
			err := errors.New("nothing to update")
			errorHandler.LogErrorThenContinue("patchProfile3", err)
			return c.Status(fiber.StatusBadRequest).JSON(&profileResponse{Payload: err.Error()})
		}
		result := p.editProfileService(bearer, userSession, profileRequestPatch)
		return c.Status(result.code).JSON(&profileResponse{Payload: result.payload})
	})
}

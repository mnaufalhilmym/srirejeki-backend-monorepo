package health

import "github.com/gofiber/fiber/v2"

// health godoc
// @Summary      Check application health
// @Description  Check application health
// @Tags         Info
// @Produce      plain
// @Success      200      {string}  string
// @Router       /health [get]
func health(h *Health) {
	h.FiberRouter.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("SriRejeki backend is running normally.")
	})
}

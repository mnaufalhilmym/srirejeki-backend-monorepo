package metrics

import "github.com/gofiber/fiber/v2/middleware/monitor"

// metrics godoc
// @Summary      Check application metrics
// @Description  Check application metrics
// @Tags         Info
// @Produce      html
// @Success      200
// @Router       /metrics [get]
func metrics(m *Metrics) {
	m.FiberRouter.Get("/", monitor.New(monitor.Config{Title: "SriRejeki Metrics"}))
}

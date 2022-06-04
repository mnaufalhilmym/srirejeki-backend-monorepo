package docs

import (
	_ "greenhouse-monitoring-iot/docs"

	"github.com/gofiber/swagger"
)

// getDocs godoc
// @Summary      See the page you are currently on
// @Description  See the page you are currently on
// @Tags         Info
// @Produce      html
// @Router       /docs [get]
func getDocs(d *Docs) {
	d.FiberRouter.Get("/*", swagger.HandlerDefault)
}

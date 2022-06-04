package main

import (
	"greenhouse-monitoring-iot/modules/auth"
	"greenhouse-monitoring-iot/modules/data"
	"greenhouse-monitoring-iot/modules/docs"
	"greenhouse-monitoring-iot/modules/farmland"
	"greenhouse-monitoring-iot/modules/health"
	"greenhouse-monitoring-iot/modules/mcu"
	"greenhouse-monitoring-iot/modules/metrics"
	"greenhouse-monitoring-iot/modules/profile"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func module(app *fiber.App, pgDb *gorm.DB, rDB *redis.Client) {
	app.Get("/", func(c *fiber.Ctx) error { return c.Redirect("/health") })

	authModule := &auth.Auth{FiberRouter: app.Group("/auth"), PgDB: pgDb, RDB: rDB}
	auth.Module(authModule)

	dataModule := &data.Data{FiberRouter: app.Group("/data"), PgDB: pgDb, RDB: rDB}
	data.Module(dataModule)

	docsModule := &docs.Docs{FiberRouter: app.Group("/docs")}
	docs.Module(docsModule)

	farmlandModule := &farmland.Farmland{FiberRouter: app.Group("/farmland"), PgDB: pgDb, RDB: rDB}
	farmland.Module(farmlandModule)

	healthModule := &health.Health{FiberRouter: app.Group("/health")}
	health.Module(healthModule)

	mcuModule := &mcu.Microcontroller{FiberRouter: app.Group("/mcu"), PgDB: pgDb, RDB: rDB}
	mcu.Module(mcuModule)

	metricModule := &metrics.Metrics{FiberRouter: app.Group("/metrics")}
	metrics.Module(metricModule)

	profileModule := &profile.Profile{FiberRouter: app.Group("/profile"), PgDB: pgDb, RDB: rDB}
	profile.Module(profileModule)
}

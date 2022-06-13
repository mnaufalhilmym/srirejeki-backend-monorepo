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

// All module used in SriRejeki backend
func module(app *fiber.App, pgDb *gorm.DB, rDB *redis.Client) {
	app.Get("/", func(c *fiber.Ctx) error { return c.Redirect("/health") })

	// Authentication module used to authenticate user
	authModule := &auth.Auth{FiberRouter: app.Group("/auth"), PgDB: pgDb, RDB: rDB}
	auth.Module(authModule)

	// Data module to subscribe and publish data by frontend and ESP
	dataModule := &data.Data{FiberRouter: app.Group("/data"), PgDB: pgDb, RDB: rDB}
	data.Module(dataModule)

	// Documentation module
	docsModule := &docs.Docs{FiberRouter: app.Group("/docs")}
	docs.Module(docsModule)

	// Farmland module
	farmlandModule := &farmland.Farmland{FiberRouter: app.Group("/farmland"), PgDB: pgDb, RDB: rDB}
	farmland.Module(farmlandModule)

	// Health module to check backend status
	healthModule := &health.Health{FiberRouter: app.Group("/health")}
	health.Module(healthModule)

	// MCU (Microcontroller Unit) module
	mcuModule := &mcu.Microcontroller{FiberRouter: app.Group("/mcu"), PgDB: pgDb, RDB: rDB}
	mcu.Module(mcuModule)

	// Metric module to monitor backend server
	metricModule := &metrics.Metrics{FiberRouter: app.Group("/metrics")}
	metrics.Module(metricModule)

	// Profile module
	profileModule := &profile.Profile{FiberRouter: app.Group("/profile"), PgDB: pgDb, RDB: rDB}
	profile.Module(profileModule)
}

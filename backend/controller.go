package main

import (
	"greenhouse-monitoring-iot/config"
	"greenhouse-monitoring-iot/internal/pkg/errorHandler"
	filehandler "greenhouse-monitoring-iot/internal/pkg/fileHandler"
	"greenhouse-monitoring-iot/internal/pkg/session"
	"os"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

// @title SriRejeki Backend API Documentation
// @version 1.0
// @description This is an API documentation for SriRejeki IoT Greenhouse Monitoring
// @contact.name API Support
// @contact.email mail@hilmy.dev
// @host localhost:80
// @BasePath /
//
func controller(pgDb *gorm.DB, rDB *redis.Client) {
	app := fiber.New()

	file, err := os.OpenFile("./app.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		errorHandler.LogErrorThenExit("controller1", err)
	}
	defer file.Close()
	filehandler.CloseFileOnInterupt(file)

	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip}:${port} ${status} - ${latency} ${method} ${path}\n",
		Output: file,
	}))

	app.Use(cors.New())

	app.Use(session.ExtendSessionIfExistBearer(rDB))

	module(app, pgDb, rDB)

	port := config.GetEnv(config.EnvEnum.Port).(string)
	app.Listen(port)
}

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
	// Initialize fiber as http framework
	app := fiber.New()

	// Use file logging saved in ./app.log
	file, err := os.OpenFile("./app.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		errorHandler.LogErrorThenExit("controller1", err)
	}
	// Close file on exit
	defer file.Close()
	filehandler.CloseFileOnInterupt(file)

	// Define logger formatting
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip}:${port} ${status} - ${latency} ${method} ${path}\n",
		Output: file,
	}))

	// Use CORS (Cross-origin resource sharing)
	app.Use(cors.New())

	// Extend User Session
	app.Use(session.ExtendSessionIfExistBearer(rDB))

	// Call module function
	module(app, pgDb, rDB)

	// http server listen to port defined in environment variable (.env)
	port := config.GetEnv(config.EnvEnum.Port).(string)
	app.Listen(port)
}

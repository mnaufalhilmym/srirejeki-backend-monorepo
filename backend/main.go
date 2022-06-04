package main

import (
	"fmt"
	"greenhouse-monitoring-iot/config"
	"greenhouse-monitoring-iot/internal/pkg/database/model"
	"greenhouse-monitoring-iot/internal/pkg/errorHandler"
	"strconv"

	"github.com/go-redis/redis"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	pgHost := config.GetEnv(config.EnvEnum.PostgresHost).(string)
	pgUser := config.GetEnv(config.EnvEnum.PostgresUser).(string)
	pgPass := config.GetEnv(config.EnvEnum.PostgresPassword).(string)
	pgDbname := config.GetEnv(config.EnvEnum.PostgresDb).(string)
	pgPort := config.GetEnv(config.EnvEnum.PostgresPort).(string)
	redisHost := config.GetEnv(config.EnvEnum.RedisHost).(string)
	redisPort := config.GetEnv(config.EnvEnum.RedisPort).(string)
	redisPass := config.GetEnv(config.EnvEnum.RedisPassword).(string)
	redisDbname, err := strconv.Atoi(config.GetEnv(config.EnvEnum.RedisDb).(string))
	if err != nil {
		errorHandler.LogErrorThenContinue("main1", err)
		redisDbname = 0
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", pgHost, pgUser, pgPass, pgDbname, pgPort)
	pgDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		errorHandler.LogErrorThenExit("main2", err)
	}
	if err := pgDb.AutoMigrate(&model.User{}, &model.Farmland{}, &model.Microcontroller{}, &model.Snapshot{}); err != nil {
		errorHandler.LogErrorThenExit("main3", err)
	}

	rDb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: redisPass,
		DB:       redisDbname,
	})

	controller(pgDb, rDb)
}

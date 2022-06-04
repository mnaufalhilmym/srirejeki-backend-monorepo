package config

import (
	"errors"
	"greenhouse-monitoring-iot/internal/pkg/errorHandler"
	"os"
)

type env struct {
	Mode             int
	Port             int
	PostgresHost     int
	PostgresPort     int
	PostgresUser     int
	PostgresPassword int
	PostgresDb       int
	RedisHost        int
	RedisPort        int
	RedisPassword    int
	RedisDb          int
	MQTTBrokerHost   int
	MQTTBrokerPort   int
	MQTTClientID     int
}

var EnvEnum = env{Mode: 0, Port: 1, PostgresHost: 2, PostgresPort: 3, PostgresUser: 4, PostgresPassword: 5, PostgresDb: 6, RedisHost: 7, RedisPort: 8, RedisPassword: 9, RedisDb: 10, MQTTBrokerHost: 11, MQTTBrokerPort: 12, MQTTClientID: 13}

var envKey = map[int]string{
	EnvEnum.Mode:             "MODE",
	EnvEnum.Port:             "PORT",
	EnvEnum.PostgresHost:     "POSTGRES_HOST",
	EnvEnum.PostgresPort:     "POSTGRES_PORT",
	EnvEnum.PostgresUser:     "POSTGRES_USER",
	EnvEnum.PostgresPassword: "POSTGRES_PASSWORD",
	EnvEnum.PostgresDb:       "POSTGRES_DB",
	EnvEnum.RedisHost:        "REDIS_HOST",
	EnvEnum.RedisPort:        "REDIS_PORT",
	EnvEnum.RedisPassword:    "REDIS_PASSWORD",
	EnvEnum.RedisDb:          "REDIS_DB",
	EnvEnum.MQTTBrokerHost:   "MQTT_BROKER_HOST",
	EnvEnum.MQTTBrokerPort:   "MQTT_BROKER_PORT",
	EnvEnum.MQTTClientID:     "MQTT_SERVER_CLIENT_ID",
}

var envFunc = map[int]interface{}{
	EnvEnum.Mode:             mode,
	EnvEnum.Port:             port,
	EnvEnum.PostgresHost:     postgresHost,
	EnvEnum.PostgresPort:     postgresPort,
	EnvEnum.PostgresUser:     postgresUser,
	EnvEnum.PostgresPassword: postgresPassword,
	EnvEnum.PostgresDb:       postgresDb,
	EnvEnum.RedisHost:        redisHost,
	EnvEnum.RedisPort:        redisPort,
	EnvEnum.RedisPassword:    redisPassword,
	EnvEnum.RedisDb:          redisDb,
	EnvEnum.MQTTBrokerHost:   mqttBrokerHost,
	EnvEnum.MQTTBrokerPort:   mqttBrokerPort,
	EnvEnum.MQTTClientID:     mqttClientID,
}

func GetAllEnv() map[int]interface{} {
	env := make(map[int]interface{}, 0)

	for key := range envKey {
		env[key] = GetEnv(key)
	}

	return env
}

func GetEnv(key int) interface{} {
	if len(envKey[key]) == 0 {
		err := errors.New("the environment variable is not listed")
		errorHandler.LogErrorThenContinue("GetEnv1", err)
		return nil
	}

	env := envFunc[key].(func() interface{})()
	return env
}

func mode() interface{} {
	mode := os.Getenv(envKey[EnvEnum.Mode])
	if len(mode) == 0 {
		mode = "debug"
	}
	return mode
}

func port() interface{} {
	port := ":" + os.Getenv(envKey[EnvEnum.Port])
	if len(port) == 1 {
		port = ":8080"
	}
	return port
}

func postgresHost() interface{} {
	host := os.Getenv(envKey[EnvEnum.PostgresHost])
	if len(host) == 0 {
		host = "localhost"
	}
	return host
}

func postgresPort() interface{} {
	port := os.Getenv(envKey[EnvEnum.PostgresPort])
	if len(port) == 0 {
		port = "5432"
	}
	return port
}

func postgresUser() interface{} {
	user := os.Getenv(envKey[EnvEnum.PostgresUser])
	if len(user) == 0 {
		user = "postgres"
	}
	return user
}

func postgresPassword() interface{} {
	password := os.Getenv(envKey[EnvEnum.PostgresPassword])
	if len(password) == 0 {
		password = ""
	}
	return password
}

func postgresDb() interface{} {
	db := os.Getenv(envKey[EnvEnum.PostgresDb])
	if len(db) == 0 {
		db = "postgres"
	}
	return db
}

func redisHost() interface{} {
	host := os.Getenv(envKey[EnvEnum.RedisHost])
	if len(host) == 0 {
		host = "localhost"
	}
	return host
}

func redisPort() interface{} {
	port := os.Getenv(envKey[EnvEnum.RedisPort])
	if len(port) == 0 {
		port = "6379"
	}
	return port
}

func redisPassword() interface{} {
	password := os.Getenv(envKey[EnvEnum.RedisPassword])
	if len(password) == 0 {
		password = ""
	}
	return password
}

func redisDb() interface{} {
	db := os.Getenv(envKey[EnvEnum.RedisDb])
	if len(db) == 0 {
		db = "0"
	}
	return db
}

func mqttBrokerHost() interface{} {
	host := os.Getenv(envKey[EnvEnum.MQTTBrokerHost])
	if len(host) == 0 {
		host = "localhost"
	}
	return host
}

func mqttBrokerPort() interface{} {
	port := os.Getenv(envKey[EnvEnum.MQTTBrokerPort])
	if len(port) == 0 {
		port = "1883"
	}
	return port
}

func mqttClientID() interface{} {
	id := os.Getenv(envKey[EnvEnum.MQTTClientID])
	if len(id) == 0 {
		id = "server"
	}
	return id
}

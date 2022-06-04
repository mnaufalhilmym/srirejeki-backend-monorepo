package mqttclient

import (
	"errors"
	"greenhouse-monitoring-iot/config"
	"greenhouse-monitoring-iot/internal/pkg/errorHandler"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Sub(num, qos *int, topic *string) (*[]*string, error) {
	brokerHost := config.GetEnv(config.EnvEnum.MQTTBrokerHost).(string)
	brokerPort := config.GetEnv(config.EnvEnum.MQTTBrokerPort).(string)
	clientID := config.GetEnv(config.EnvEnum.MQTTClientID).(string)

	choke := make(chan *string)

	opts := mqtt.NewClientOptions().AddBroker(brokerHost + ":" + brokerPort).SetClientID(clientID)
	opts.SetAutoReconnect(false)
	opts.SetConnectRetry(false)
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		payload := string(msg.Payload())
		choke <- &payload
	})

	if num == nil {
		*num = 1
	}
	if qos == nil {
		*qos = 0
	}
	if topic == nil {
		*topic = ""
	}

	receiveCount := 0

	client := mqtt.NewClient(opts)

	timeLimit := time.NewTimer(8 * time.Second)

	if token := client.Connect(); !token.WaitTimeout(7*time.Second) && token.Error() != nil {
		errorHandler.LogErrorThenContinue("mqttClient/Sub2", token.Error())
		return nil, token.Error()
	}

	if token := client.Subscribe(*topic, byte(*qos), nil); !token.WaitTimeout(7*time.Second) && token.Error() != nil {
		errorHandler.LogErrorThenContinue("mqttClient/Sub3", token.Error())
		return nil, token.Error()
	}

	msg := []*string{}

	for receiveCount < *num {
		select {
		case incoming := <-choke:
			msg = append(msg, incoming)
			receiveCount++
		case <-timeLimit.C:
			client.Disconnect(0)
			err := errors.New("subscribe timeout")
			errorHandler.LogErrorThenContinue("mqttClient/Sub1", err.Error())
			return nil, err
		}
	}

	client.Disconnect(0)

	return &msg, nil
}

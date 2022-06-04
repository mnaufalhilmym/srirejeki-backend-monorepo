package mqttclient

import (
	"greenhouse-monitoring-iot/config"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// num: The number of messages to publish or subscribe (default 1)
// qos: The Quality of Service 0,1,2 (default 0)
// topic: The topic name to/from which to publish/subscribe (default empty)
// message: The message text to publish (default empty)
func Pub(num, qos *int, topic, payload *string) error {
	brokerHost := config.GetEnv(config.EnvEnum.MQTTBrokerHost).(string)
	brokerPort := config.GetEnv(config.EnvEnum.MQTTBrokerPort).(string)
	clientID := config.GetEnv(config.EnvEnum.MQTTClientID).(string)

	opts := mqtt.NewClientOptions().AddBroker(brokerHost + ":" + brokerPort).SetClientID(clientID)
	opts.SetWriteTimeout(3 * time.Second)

	if num == nil {
		*num = 1
	}
	if qos == nil {
		*qos = 0
	}
	if topic == nil {
		*topic = ""
	}
	if payload == nil {
		*payload = ""
	}

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	for i := 0; i < *num; i++ {
		token := client.Publish(*topic, byte(*qos), false, *payload)
		token.Wait()
	}

	client.Disconnect(250)

	return nil
}

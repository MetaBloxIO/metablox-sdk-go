package iot

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"strings"
)

func PublishData(mc *MqttClient, topic string, data interface{}) (msgId string, err error) {
	req := NewMessage(data)
	jb, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	return req.Id, mc.Publish(topic, jb)
}

func SubscribeData[T any](mc *MqttClient, topic string, handler func(res *Message[T], err error)) (err error) {

	onMessage := func(client mqtt.Client, message mqtt.Message) {
		msg := &Message[T]{}
		if err = json.Unmarshal(message.Payload(), &msg); err != nil {
			handler(&Message[T]{rawMsg: message}, err)
			return
		}
		msg.rawMsg = message
		handler(msg, nil)
	}

	if tc := mc.Client.Subscribe(topic, mc.qos, onMessage); tc.Wait() && tc.Error() != nil {
		return tc.Error()
	}

	mc.topics[topic] = onMessage
	log.Println(fmt.Sprintf("Subscribe[%s] success", topic))
	return nil
}

func ToTopic(str ...string) string {
	return strings.Join(str, "/")
}

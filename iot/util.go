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

func SubscribeData(mc *MqttClient, topic string, handler func(res *Message, err error)) (err error) {

	onMessage := func(client mqtt.Client, message mqtt.Message) {

		msg := &Message{}
		split := strings.Split(message.Topic(), "/")
		var productKey = ""
		var clientId = ""
		if (strings.HasPrefix(message.Topic(), Sys) ||
			strings.HasPrefix(message.Topic(), Biz) ||
			strings.HasPrefix(message.Topic(), Ota) ||
			strings.HasPrefix(message.Topic(), Ext)) && len(split) >= 3 {
			productKey = split[1]
			clientId = split[2]
		}

		if err = json.Unmarshal(message.Payload(), &msg); err != nil {
			handler(&Message{rawMsg: message, productKey: productKey, clientId: clientId}, err)
			return
		}
		msg.rawMsg = message
		msg.productKey = productKey
		msg.clientId = clientId
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

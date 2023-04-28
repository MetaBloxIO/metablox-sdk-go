package iot

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"strings"
)

func PublishData(mc *MqttClient, bizType TopicType, data interface{}) (msgId string, err error) {
	req := NewMessageReq(data)
	jb, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	opt := mc.Client.OptionsReader()
	return req.Id, mc.Publish(bizType.Topic(ProductKey, opt.ClientID()), jb)
}

func SubscribeData[T any](mc *MqttClient, topicType TopicType, handler func(res *MessageRes[T], err error)) (err error) {
	opt := mc.Client.OptionsReader()
	topic := topicType.Topic(ProductKey, opt.ClientID())

	onMessage := func(client mqtt.Client, message mqtt.Message) {
		res := &MessageRes[T]{}
		if err = json.Unmarshal(message.Payload(), &res); err != nil {
			handler(nil, err)
			return
		}
		handler(res, nil)
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

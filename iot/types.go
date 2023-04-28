package iot

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"time"
)

type Message struct {
	rawMsg  mqtt.Message
	Id      string      `json:"id"`
	Ts      int64       `json:"ts"`
	Data    interface{} `json:"data"`
	Version string      `json:"version"`
}

func (m Message) RawMessage() mqtt.Message {
	return m.rawMsg
}

func NewMessage(data interface{}) *Message {
	return &Message{
		Id:      uuid.NewString(),
		Ts:      time.Now().UnixMilli(),
		Data:    data,
		Version: Version,
	}
}

func NewMessageWithId(msgId string, data interface{}) *Message {
	return &Message{
		Id:      msgId,
		Ts:      time.Now().UnixMilli(),
		Data:    data,
		Version: Version,
	}
}

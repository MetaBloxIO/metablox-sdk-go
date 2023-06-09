package iot

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"time"
)

type Message struct {
	rawMsg     mqtt.Message
	productKey string
	clientId   string
	Id         string      `json:"id"`
	Ts         int64       `json:"ts"`
	Data       interface{} `json:"data"`
	Version    string      `json:"version"`
}

func (m Message) ProductKey() string {
	return m.productKey
}

func (m Message) ClientId() string {
	return m.clientId
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

type ReplyData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

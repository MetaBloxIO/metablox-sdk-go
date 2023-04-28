package iot

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"time"
)

type Message[T any] struct {
	rawMsg  mqtt.Message
	Id      string `json:"id"`
	Ts      int64  `json:"ts"`
	Data    T      `json:"data"`
	Version string `json:"version"`
}

func (m Message[T]) RawMessage() mqtt.Message {
	return m.rawMsg
}

func NewMessage[T any](data T) *Message[T] {
	return &Message[T]{
		Id:      uuid.NewString(),
		Ts:      time.Now().UnixMilli(),
		Data:    data,
		Version: Version,
	}
}

func NewMessageWithId[T any](msgId string, data T) *Message[T] {
	return &Message[T]{
		Id:      msgId,
		Ts:      time.Now().UnixMilli(),
		Data:    data,
		Version: Version,
	}
}

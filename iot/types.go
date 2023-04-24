package iot

import (
	"github.com/google/uuid"
	"time"
)

type MessageReq[T any] struct {
	Id      string `json:"id"`
	Ts      int64  `json:"ts"`
	Params  T      `json:"params"`
	Version string `json:"version"`
	Timeout int64  `json:"timeout"`
}

func NewMessageReq[T any](params T) *MessageReq[T] {
	return &MessageReq[T]{
		Id:      uuid.NewString(),
		Ts:      time.Now().UnixMilli(),
		Params:  params,
		Version: Version,
	}
}

type MessageRes[T any] struct {
	Id   string `json:"id"`
	Ts   int64  `json:"ts"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func NewMessageRes[T any](id string, code int, msg string, data T) *MessageRes[T] {
	return &MessageRes[T]{
		Id:   id,
		Ts:   time.Now().UnixMilli(),
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

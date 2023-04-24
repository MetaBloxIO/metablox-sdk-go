package iot

type MessageReq[T any] struct {
	Id      string `json:"id"`
	Ts      int64  `json:"ts"`
	Method  string `json:"method"`
	Params  T      `json:"params"`
	Version string `json:"version"`
}

type MessageRes[T any] struct {
	Id   string `json:"id"`
	Ts   int64  `json:"ts"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

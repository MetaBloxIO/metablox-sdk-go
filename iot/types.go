package iot

type MessageReq struct {
	Id      string      `json:"id"`
	Ts      int64       `json:"ts"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Version string      `json:"version"`
}

type MessageRes struct {
	Id   string      `json:"id"`
	Ts   int64       `json:"ts"`
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

package iot

const ProductKey = "metablox"
const DeviceKey = "miner"

const SLW = "+" // single level wildcard
const MLW = "#" // multi level wildcard

const Sys = "sys"
const Biz = "biz"
const Ota = "ota"

const SysTopicPrefix = Sys + "/" + ProductKey + "/" + DeviceKey
const BizTopicPrefix = Biz + "/" + ProductKey + "/" + DeviceKey
const OtaTopicPrefix = Ota + "/" + ProductKey + "/" + DeviceKey

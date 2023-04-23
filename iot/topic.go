package iot

const ProductKey = "metablox"
const DeviceKey = "miner"

const SingleLevelWildcard = "+" // single level wildcard
const MultiLevelWildcard = "#"  // multi level wildcard

const Sys = "sys"
const Biz = "biz"
const Ota = "ota"

const SysTopicPrefix = Sys + "/" + ProductKey + "/" + DeviceKey
const BizTopicPrefix = Biz + "/" + ProductKey + "/" + DeviceKey
const OtaTopicPrefix = Ota + "/" + ProductKey + "/" + DeviceKey

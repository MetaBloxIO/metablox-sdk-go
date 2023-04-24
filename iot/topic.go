package iot

import "fmt"

const ProductKey = "miner" // Temporarily fixed as “miner”
const DeviceKey = "${clientid}"

const SLW = "+" // single level wildcard
const MLW = "#" // multi level wildcard
const _REPLY = "_reply"
const _ACK = "_ack"

const SSys = "$SYS"
const Sys = "/sys"
const Biz = "/biz"
const Ota = "/ota"
const Ext = "/ext"
const Shadow = "/shadow"
const Broadcast = "/broadcast"

type TopicType string

func (t TopicType) Topic(productKey, deviceKey string) string {
	return fmt.Sprintf(string(t), productKey, deviceKey)
}

const (
	SysWillStatus TopicType = Sys + "/%s/%s/will/status" // client Pub

	SysHeartbeatUpdate      TopicType = Sys + "/%s/%s/heartbeat/update" // client Pub
	SysHeartbeatUpdateReply TopicType = SysHeartbeatUpdate + _REPLY     // client Sub

	SysSettingsGet    TopicType = Sys + "/%s/%s/settings/get"    // client Pub
	SysSettingsUpdate TopicType = Sys + "/%s/%s/settings/update" // client Sub

	BizWorkloadValidate      TopicType = Biz + "/%s/%s/workload/validate" // client Pub
	BizWorkloadValidateReply TopicType = BizWorkloadValidate + _REPLY     // client Sub

	OtaFirmwareCheck   TopicType = Ota + "/%s/%s/firmware/check"   // client Pub
	OtaFirmwareUpgrade TopicType = Ota + "/%s/%s/firmware/upgrade" // client Sub
)

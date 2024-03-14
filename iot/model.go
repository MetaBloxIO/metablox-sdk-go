package iot

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Device lifecycle changes: device online, device offline, device deleted, device renamed, device binding, device unbinding, device information changes, disable, enable.
// Historical data reporting of physical model: historical data reported by devices for attributes, events, and services.
// Firmware upgrade status notification: firmware upgrade status reported by devices.
// Device message reporting: real-time data reported by devices for attributes, events, and services.
// Device status change notification: real-time data reported by devices for attributes, events, and services.

var defaultMessageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	log.Printf("TOPIC: %s\n", msg.Topic())
	log.Printf("MSG: %s\n", string(msg.Payload()))
}

type MinerWorkloadData struct {
	Identity *Identity `json:"identity" v:"required" binding:"required"`
	Qos      *Qos      `json:"qos" v:"required" binding:"required"`
	Tracks   string    `json:"tracks"`
}

// MinerGuestData is the golang structure for table miner_guest_record.
type MinerGuestData struct {
	MinerSn       string `json:"minerSn"       description:""`
	ClientId      string `json:"clientId"      description:"Device ID"`
	Hostname      string `json:"hostname"      description:"guest phone brand"`
	Mac           string `json:"mac"           description:"guest device mac"`
	DisplayName   string `json:"displayName"   description:"guest display name"`
	PhoneNumber   string `json:"phoneNumber"   description:"Phone number"`
	BirthDate     string `json:"birthDate"     description:"Birthday"`
	Gender        int    `json:"gender"        description:"Gender(0-Unknown,1-Male,2-Female)"`
	Email         string `json:"email"         description:"Email address"`
	EmailVerified bool   `json:"emailVerified" description:"Whether the email has been verified"`
	CountryCode   string `json:"countryCode"   description:"Country Code"`
	LanguageCode  string `json:"languageCode"  description:"Language code"`
}

type MinerSettingsData struct {
	CreatedAt    int64  `json:"createdAt"`
	UpdatedAt    int64  `json:"updatedAt"`
	Banner       string `json:"banner"`
	ForwardLink  string `json:"forwardLink"`
	Timeout      int    `json:"timeout"`
	UploadRate   int    `json:"uploadRate"`
	DownloadRate int    `json:"downloadRate"`
}

type SystemCommandData struct {
	Id      int      `json:"id" binding:"required" v:"required"`
	BinName string   `json:"binName" binding:"required" v:"required"`
	Args    []string `json:"args" binding:"required" v:"required"`
}

type SystemCommandReplyData struct {
	Id      int    `json:"id" binding:"required" v:"required"`
	Output  string `json:"output" binding:"required" v:"required"`
	Success bool   `json:"success" binding:"required" v:"required"`
}

type MinerHeartbeatData struct {
	Sn            string       `json:"sn" binding:"required" v:"required"`
	Did           string       `json:"did" binding:"required" v:"required"`
	Mac           string       `json:"mac" binding:"required" v:"required"`
	PublicIP      string       `json:"publicIP"`
	WalletAddress string       `json:"walletAddress" `
	RadioStatus   []WiFiStatus `json:"radioStatus" binding:"required"`
	SystemStatus  SystemInfo   `json:"systemStatus" binding:"required"`
	DeviceStatus  DeviceInfo   `json:"deviceStatus" binding:"required"`
}

type OtaFirmwareUpgradeData struct {
	Id           int    `json:"id" binding:"required" v:"required"`
	Distribution string `json:"distribution" binding:"required" v:"required"`
	Version      string `json:"version" binding:"required" v:"required"`
	Target       string `json:"target" binding:"required" v:"required"`
	Description  string `json:"description" binding:"required" v:"required"`
	ImageUrl     string `json:"imageUrl" binding:"required" v:"required"`
	IsForced     bool   `json:"isForced" binding:"required" v:"required"`
	Sha256       string `json:"sha256" binding:"required" v:"required"`
	Model        string `json:"model" binding:"required" v:"required"`
	BoardName    string `json:"boardName" binding:"required" v:"required"`
}

type OtaFirmwareUpgradeReplyData struct {
	UpgradeId int    `json:"upgradeId" binding:"required" v:"required"`
	Message   string `json:"message" binding:"required" v:"required"`
	Success   bool   `json:"success" binding:"required" v:"required"`
}

type OtaFirmwareCheckData struct {
	Distribution string `json:"distribution" binding:"required" v:"required"`
	Version      string `json:"version" binding:"required" v:"required"`
	Target       string `json:"target" binding:"required" v:"required"`
	Description  string `json:"description" binding:"required" v:"required"`
	Model        string `json:"model" binding:"required" v:"required"`
	BoardName    string `json:"boardName" binding:"required" v:"required"`
	Sn           string `json:"sn" binding:"required" v:"required"`
}

// =================================================================================

type DeviceInfo struct {
	Kernel    string `json:"kernel"`
	Hostname  string `json:"hostname"`
	System    string `json:"system"`
	Model     string `json:"model"`
	BoardName string `json:"board_name"`
	Release   struct {
		Distribution string `json:"distribution"`
		Version      string `json:"version"`
		Revision     string `json:"revision"`
		Target       string `json:"target"`
		Description  string `json:"description"`
	} `json:"release"`
}

type SystemInfo struct {
	Localtime int   `json:"localtime"`
	Uptime    int   `json:"uptime"`
	Load      []int `json:"load"`
	Memory    struct {
		Total     int `json:"total"`
		Free      int `json:"free"`
		Shared    int `json:"shared"`
		Buffered  int `json:"buffered"`
		Available int `json:"available"`
		Cached    int `json:"cached"`
	} `json:"memory"`
	Swap struct {
		Total int `json:"total"`
		Free  int `json:"free"`
	} `json:"swap"`
}

type WiFiStatus struct {
	Disabled        bool   `json:"disabled"`
	Type            string `json:"type"` // openroaming | free
	Phy             string `json:"phy"`
	Ssid            string `json:"ssid"`
	Bssid           string `json:"bssid"`
	Country         string `json:"country"`
	Mode            string `json:"mode"`
	Channel         int    `json:"channel"`
	CenterChan1     int    `json:"center_chan1"`
	Frequency       int    `json:"frequency"`
	FrequencyOffset int    `json:"frequency_offset"`
	Txpower         int    `json:"txpower"`
	TxpowerOffset   int    `json:"txpower_offset"`
	Quality         int    `json:"quality"`
	QualityMax      int    `json:"quality_max"`
	Signal          int    `json:"signal"`
	Noise           int    `json:"noise"`
	Bitrate         int    `json:"bitrate"`
	Encryption      struct {
		Enabled        bool     `json:"enabled"`
		Wpa            []int    `json:"wpa,omitempty"`
		Authentication []string `json:"authentication,omitempty"`
		Ciphers        []string `json:"ciphers,omitempty"`
	} `json:"encryption"`
	Htmodes []string         `json:"htmodes"`
	Hwmodes []string         `json:"hwmodes"`
	Hwmode  string           `json:"hwmode"`
	Htmode  string           `json:"htmode"`
	Station []WifiClientInfo `json:"station"`
}

type WifiClientInfo struct {
	Hostname      string `json:"hostname"`
	Mac           string `json:"mac"`
	Signal        int    `json:"signal"`
	SignalAvg     int    `json:"signal_avg"`
	Noise         int    `json:"noise"`
	Inactive      int    `json:"inactive"`
	ConnectedTime int    `json:"connected_time"`
}

// =================================================================================

type Qos struct {
	Bandwidth  string `json:"bandwidth"`
	Rssi       string `json:"rssi"`
	PacketLose string `json:"packetLose"`
	Latency    string `json:"latency"`
	Nonce      string `json:"nonce" binding:"required"`
	Signature  string `json:"signature" binding:"required"`
}

type Identity struct {
	Validator interface{} `json:"validator" v:"required" binding:"required"`
	Miner     interface{} `json:"miner" v:"required" binding:"required"`
}

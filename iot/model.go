package iot

// Device lifecycle changes: device online, device offline, device deleted, device renamed, device binding, device unbinding, device information changes, disable, enable.
// Historical data reporting of physical model: historical data reported by devices for attributes, events, and services.
// Firmware upgrade status notification: firmware upgrade status reported by devices.
// Device message reporting: real-time data reported by devices for attributes, events, and services.
// Device status change notification: real-time data reported by devices for attributes, events, and services.

var eventMap = map[string]interface{}{
	"online":  nil,
	"offline": nil,
	"bind":    nil,
	"unbind":  nil,
	"disable": nil,
	"enable":  nil,
}

var actionMap = map[string]interface{}{
	"heartbeat": nil,
	"settings":  nil,
	"workload":  nil,
	"guest":     nil,
}

type WorkloadData struct {
	Identity *Identity `json:"identity" v:"required"`
	Qos      *Qos      `json:"qos" v:"required"`
	Tracks   string    `json:"tracks"`
	Sn       string    `json:"sn"`
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

type SettingsData struct {
	CreatedAt    int64  `json:"createdAt"`
	UpdatedAt    int64  `json:"updatedAt"`
	Banner       string `json:"banner"`
	ForwardLink  string `json:"forwardLink"`
	Timeout      int    `json:"timeout"`
	UploadRate   int    `json:"uploadRate"`
	DownloadRate int    `json:"downloadRate"`
}

type HeartbeatData struct {
	Sn            string       `json:"sn" binding:"required" v:"required"`
	Mac           string       `json:"mac" binding:"required" v:"required"`
	Did           string       `json:"did" binding:"required" v:"required"`
	WalletAddress string       `json:"walletAddress" binding:"required" v:"required"`
	RadioStatus   []WiFiStatus `json:"radioStatus" binding:"required"`
	SystemStatus  SystemInfo   `json:"SystemStatus" binding:"required"`
	DeviceStatus  DeviceInfo   `json:"deviceStatus" binding:"required"`
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

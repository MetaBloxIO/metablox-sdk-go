package iot

import (
	"crypto/tls"
	"errors"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

// MqttConfig  MQTT Config
type MqttConfig struct {
	Broker                string `json:"broker"`
	ClientId              string `json:"clientId"` //SN
	Username              string `json:"username"` //did (optional)
	Password              string `json:"password"`
	CertFile              string `json:"certFile"`
	CertPrivateKey        string `json:"certPrivateKey"`
	WillEnabled           bool   `json:"willEnabled"`
	WillPayload           string `json:"willPayload"`
	WillQos               byte   `json:"willQos"`
	Qos                   byte   `json:"qos"`
	Retained              bool   `json:"retained"`
	ConnectionLostHandler mqtt.ConnectionLostHandler
	OnConnectHandler      mqtt.OnConnectHandler
}

// MqttClient MQTT Client
type MqttClient struct {
	qos      byte
	retained bool
	Client   mqtt.Client
	topics   map[string]mqtt.MessageHandler
}

func (mc *MqttClient) ClientId() string {
	reader := mc.Client.OptionsReader()
	return reader.ClientID()
}

func (mc *MqttClient) SubscribeDefault(handler ...mqtt.MessageHandler) error {
	if !mc.Client.IsConnected() {
		return errors.New("mqtt client is not connected")
	}
	opt := mc.Client.OptionsReader()
	defaultSubTopics := []TopicType{
		SysHeartbeatUpdateReply,
		SysSettingsUpdate,
		BizWorkloadValidateReply,
		OtaFirmwareUpgrade,
	}
	var defaultHandler mqtt.MessageHandler
	if len(handler) > 0 {
		defaultHandler = handler[0]
	} else {
		defaultHandler = defaultMessageHandler
	}

	for _, topicType := range defaultSubTopics {
		topic := topicType.Topic(ProductKey, opt.ClientID())
		if err := mc.Subscribe(topic, defaultHandler); err != nil {
			return err
		}
		mc.topics[topic] = defaultHandler
	}
	return nil
}

func NewMqttClient(cfg MqttConfig) (*MqttClient, error) {
	var c MqttClient

	if cfg.OnConnectHandler == nil {
		cfg.OnConnectHandler = func(client mqtt.Client) {
			log.Println("mqtt client connected")
		}
	}
	if cfg.ConnectionLostHandler == nil {
		cfg.ConnectionLostHandler = func(client mqtt.Client, err error) {
			log.Println("mqtt client disconnected", err)
		}
	}

	opts := mqtt.NewClientOptions().
		AddBroker(cfg.Broker).
		SetClientID(cfg.ClientId).
		SetUsername(cfg.Username).
		SetPassword(cfg.Password).
		SetOnConnectHandler(cfg.OnConnectHandler).
		SetConnectionLostHandler(cfg.ConnectionLostHandler).
		SetProtocolVersion(5)

	//Determine whether to set up a will
	if cfg.WillEnabled {
		opts.SetWill(SysWillStatus.Topic(ProductKey, cfg.ClientId), cfg.WillPayload, cfg.WillQos, cfg.Retained)
	}
	//Determine whether to set up a certificate
	if cfg.CertFile != "" {
		tlsConfig, err := NewTLSConfig(cfg.CertFile, cfg.CertPrivateKey)
		if err != nil {
			return nil, err
		}
		opts.SetTLSConfig(tlsConfig)
	}
	c.Client = mqtt.NewClient(opts)
	c.qos = cfg.Qos
	c.retained = cfg.Retained
	c.topics = make(map[string]mqtt.MessageHandler)
	// Judging by the status of the token
	if tc := c.Client.Connect(); tc.Wait() && tc.Error() != nil {
		return nil, tc.Error()
	}
	return &c, nil
}

// Publish  Mqtt message.
func (mc *MqttClient) Publish(topic string, payload []byte) error {
	if mc != nil && mc.Client.IsConnected() {
		if tc := mc.Client.Publish(topic, mc.qos, mc.retained, payload); tc.Wait() && tc.Error() != nil {
			return tc.Error()
		}
		return nil
	}
	return errors.New("mqttClient is nil or disconnected")
}

// Subscribe subscribe a Mqtt topic.
func (mc *MqttClient) Subscribe(topic string, onMessage mqtt.MessageHandler) error {
	if tc := mc.Client.Subscribe(topic, mc.qos, onMessage); tc.Wait() && tc.Error() != nil {
		return tc.Error()
	}
	mc.topics[topic] = onMessage
	log.Println(fmt.Sprintf("Subscribe[%s] success", topic))
	return nil
}

// Unsubscribe unsubscribe a Mqtt topic.
func (mc *MqttClient) Unsubscribe(topics ...string) error {
	if tc := mc.Client.Unsubscribe(topics...); tc.Wait() && tc.Error() != nil {
		return tc.Error()
	}
	for _, topic := range topics {
		delete(mc.topics, topic)
	}
	return nil
}

func (mc *MqttClient) Close() {
	mc.Client.Disconnect(250) //Millisecond
}

func (mc *MqttClient) PublishData(topicType TopicType, data interface{}) (string, error) {
	return PublishData(mc, topicType.Topic(ProductKey, mc.ClientId()), data)
}

func (mc *MqttClient) PublishDataTo(clientId string, topicType TopicType, data interface{}) (string, error) {
	return PublishData(mc, topicType.Topic(ProductKey, clientId), data)
}

func (mc *MqttClient) SubscribeData(topicType TopicType, handler func(res *Message, err error)) (err error) {
	return SubscribeData(mc, topicType.Topic(ProductKey, mc.ClientId()), handler)
}

func (mc *MqttClient) SubscribeDataFrom(clientId string, topicType TopicType, handler func(res *Message, err error)) (err error) {
	return SubscribeData(mc, topicType.Topic(ProductKey, clientId), handler)
}

// NewTLSConfig New TLS Config
func NewTLSConfig(certFile string, certPrivateKey string) (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(certFile, certPrivateKey)
	if err != nil {
		return nil, err
	}
	return &tls.Config{
		ClientAuth:         tls.NoClientCert, //no certificate required
		ClientCAs:          nil,              //do not verify certificate
		InsecureSkipVerify: true,             //accept any certificate presented by the server and any hostname in that certificate
		Certificates:       []tls.Certificate{cert},
	}, nil
}

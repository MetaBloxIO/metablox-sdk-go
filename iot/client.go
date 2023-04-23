package iot

import (
	"crypto/tls"
	"errors"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"time"
)

// MqttConfig  MQTT Config
type MqttConfig struct {
	Broker         string
	ClientId       string
	Username       string
	Password       string
	CertFilePath   string
	CertPrivateKey string
	WillEnabled    bool
	WillTopic      string
	WillPayload    string
	WillQos        byte
	Qos            byte
	Retained       bool
	onConnect      mqtt.OnConnectHandler
	connectionLost mqtt.ConnectionLostHandler
}

// MqttClient MQTT Client
type MqttClient struct {
	qos      byte
	retained bool
	Client   mqtt.Client
	topics   map[string]mqtt.MessageHandler
}

func NewMqttClient(cfg MqttConfig) (*MqttClient, error) {
	var c MqttClient
	opts := mqtt.NewClientOptions().
		AddBroker(cfg.Broker).
		SetClientID(cfg.ClientId).
		SetMaxReconnectInterval(time.Second * 5)
	//Determine whether to set up a will
	if cfg.WillEnabled {
		opts.SetWill(cfg.WillTopic, cfg.WillPayload, cfg.WillQos, cfg.Retained)
	}
	//Determine whether to set up a certificate
	if cfg.CertFilePath != "" {
		tlsConfig, err := newTLSConfig(cfg.CertFilePath, cfg.CertPrivateKey)
		if err != nil {
			return nil, err
		}
		opts.SetTLSConfig(tlsConfig)
	} else {
		opts.SetUsername(cfg.Username).SetPassword(cfg.Password)
	}

	if cfg.onConnect == nil {
		cfg.onConnect = func(c mqtt.Client) {}
	}
	if cfg.connectionLost == nil {
		cfg.connectionLost = func(c mqtt.Client, err error) {}
	}
	opts.SetOnConnectHandler(c.OnConnectHandler(cfg.onConnect)).
		SetConnectionLostHandler(c.ConnectionLostHandler(cfg.connectionLost))
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
func (mc *MqttClient) Subscribe(topics []string, onMessage mqtt.MessageHandler) error {
	for _, topic := range topics {
		if tc := mc.Client.Subscribe(topic, mc.qos, onMessage); tc.Wait() && tc.Error() != nil {
			return tc.Error()
		}
		mc.topics[topic] = onMessage
		log.Println(fmt.Sprintf("Subscribe[%s] success", topic))
	}
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

func (mc *MqttClient) OnConnectHandler(handler mqtt.OnConnectHandler) mqtt.OnConnectHandler {
	return func(c mqtt.Client) {
		for topic, onMessage := range mc.topics {
			mc.Client.Subscribe(topic, mc.qos, onMessage)
		}
		handler(c)
	}
}

func (mc *MqttClient) ConnectionLostHandler(handler mqtt.ConnectionLostHandler) mqtt.ConnectionLostHandler {
	return func(c mqtt.Client, e error) {
		handler(c, e)
	}
}

// newTLSConfig  new TLS Config
func newTLSConfig(certFile string, privateKey string) (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(certFile, privateKey)
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

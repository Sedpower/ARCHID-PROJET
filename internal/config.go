package internal

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	BrokerUrl           string `yaml:"brokerUrl"`
	BrokerPort          string `yaml:"brokerPort"`
	BrokerBaseTopicPath string `yaml:"brokerBaseTopicPath"`
	RedisHost           string `yaml:"redisHost"`
	RedisPort           string `yaml:"redisPort"`
	RedisProtocol       string `yaml:"redisProtocol"`
	Publisher           struct {
		ClientId   string   `yaml:"clientId"`
		QoS        int      `yaml:"QoS"`
		ApiUrl     string   `yaml:"apiUrl"`
		ListOfIATA []string `yaml:"listOfIATA"`
	} `yaml:"publisher"`
	Subscriber struct {
		ClientId string `yaml:"clientId"`
		QoS      int    `yaml:"QoS"`
		Topic    string `yaml:"topic"`
	} `yaml:"subscriber"`
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func LoadConfig() Config {
	f, err := os.Open("config.yaml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		processError(err)
	}
	return cfg
}

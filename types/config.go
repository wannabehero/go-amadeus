package types

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type AmadeusConfig struct {
	Username       string `yaml:"username"`
	Password       string `yaml:"password"`
	PseudoCityCode string `yaml:"pseudoCityCode"`
	AgentDutyCode  string `yaml:"agentDutyCode"`
	RequestorType  string `yaml:"requestorType"`
	POSType        int    `yaml:"posType"`
	URL            string `yaml:"url"`
	WSAP           string `yaml:"wsap"`
}

func LoadConfigFromFile(filePath string) AmadeusConfig {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Panicf("Cannot read config: %v", err)
	}

	var config AmadeusConfig

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Panicf("Cannot parse yaml: %v", err)
	}

	return config
}

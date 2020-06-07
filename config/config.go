package config

import (
	"gopkg.in/yaml.v2" 
	"log"
)

type Config struct {
	config map[string]interface{}
}

func (c *Config) SetFromBytes(data []byte) *Config {
	err := yaml.Unmarshal(data, data)
	if err != nil {
		log.Fatalf("Unmarshal %v",err)
	}

	return c
}

func (c *Config) Get(serviceName string) (map[string]interface{}, error) {
	return nil, nil

}

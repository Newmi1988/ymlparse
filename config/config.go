package config

import (
	"gopkg.in/yaml.v2" 
	"log"
	"fmt"

	// "github.com/Newmi1988/ymlparse/tools"
)

type Config struct {
	config map[string]interface{}
}

func (c *Config) SetFromBytes(data []byte) *Config {
	var rawConfig interface{}

	err := yaml.Unmarshal(data, &rawConfig)
	if err != nil {
		log.Fatalf("Unmarshal %v",err)
	}

	untypedConfig, ok := rawConfig.(map[interface{}]interface{})
	if !ok {
		log.Fatalf("Config is not a map")
	}

	config, err := convertKeysToString(untypedConfig)
	if err != nil {
		log.Fatalf("%v",err)
	}

	c.config = config

	return nil
}


func convertKeysToString(m map[interface{}]interface{}) (map[string]interface{}, error) {
	n := make(map[string]interface{})

	for key, value := range m {
		str, ok := key.(string)
		if !ok {
			return nil, fmt.Errorf("config key is not a string")
		}

		if vMap, ok := value.(map[interface{}]interface{}); ok {
			var err error
			value, err = convertKeysToString(vMap)
			if err != nil {
				return nil, err
			}
		}

		n[str] = value
	}

	return n, nil
}


func (c *Config) is_yml_valid() (bool,error) {
	l := len(c.config)
	keys := make([]string, l)

	i := 0
	for k,_ := range c.config {
		keys[i] = k
		i++
	}

	// if every key has key values bellow yaml is corroct
	for _, k := range keys {
		_, ok := c.config[k].(map[string]interface{})
		if !ok {
			
			return false, fmt.Errorf("yaml is invalid. Keyerror : %v ", k)
			
		}
	}

	return true, nil
}



func (c *Config) Get(serviceName string) (map[string]interface{}, error) {
	
	// is this the base the config
	ok, err := c.is_yml_valid()
	if !ok {
		fmt.Println(err)
	}

	config := make(map[string]interface{})
	serviceConfig, ok := c.config[serviceName].(map[string]interface{})
	if !ok {
		log.Fatalf("wrong format")
	}

	for k,v := range serviceConfig {
		config[k] = v
	}

	return config, nil
}

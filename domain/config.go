package domain

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

type Config struct {
	config map[string]interface{}
}

func (c *Config) SetFromBytes(data []byte) error {
	var untypedConfig map[interface{}]interface{}
	if err := yaml.Unmarshal(data, &untypedConfig); err != nil {
		return err
	}

	config, err := convertKeys2Strings(untypedConfig)
	if err != nil {
		return err
	}
	c.config = config
	return nil
}

func (c *Config) Get(serviceName string) (map[string]interface{}, error) {
	b, ok := c.config["base"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("base config should be a map")
	}

	if _, ok = c.config[serviceName]; !ok {
		return b, nil
	}

	a, ok := c.config[serviceName].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("service %q is not valid", serviceName)
	}

	config := make(map[string]interface{})
	for k, v := range b {
		config[k] = v
	}
	for k, v := range a {
		config[k] = v
	}
	return config, nil
}

func convertKeys2Strings(from map[interface{}]interface{}) (map[string]interface{}, error) {
	n := make(map[string]interface{})

	for k, v := range from {
		str, ok := k.(string)
		if !ok {
			return nil, fmt.Errorf("config keys should be strings totally")
		}
		if vMap, ok := v.(map[interface{}]interface{}); ok {
			var err error
			v, err = convertKeys2Strings(vMap)
			if err != nil {
				return nil, err
			}
		}
		n[str] = v
	}
	return n, nil
}

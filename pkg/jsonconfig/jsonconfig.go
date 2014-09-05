/*
JSON config parser.

Inspires by github.com/stathat/jconfig.
*/

package jsonconfig

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"
)

type Config struct {
	data map[string]interface{}
}

func newConfig() *Config {
	rv := new(Config)
	rv.data = make(map[string]interface{})

	return rv
}

func LoadFromFile(filename string) (*Config, error) {
	config := newConfig()
	err := config.parseFromFile(filename)

	return config, err
}

func (c *Config) parseFromMap(obj map[string]interface{}) {
	for k, v := range obj {
		c.data[k] = v
	}
}

func (c *Config) parseFromFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	b := new(bytes.Buffer)
	_, err = b.ReadFrom(f)
	if err != nil {
		return err
	}
	return c.parseFromBytes(b.Bytes())
}

func (c *Config) parseFromBytes(b []byte) error {
	err := json.Unmarshal(b, &c.data)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) getNestedMap(keys []string) map[string]interface{} {
	obj := c.data

	for _, key := range keys {
		next, present := obj[key]
		if !present {
			return nil
		}
		obj = next.(map[string]interface{})
	}

	return obj
}

func (c *Config) getNested(key string) interface{} {
	keys := strings.Split(key, ".")

	obj := c.data
	for _, key := range keys[0 : len(keys)-1] {
		next, present := obj[key]
		if !present {
			return nil
		}
		obj = next.(map[string]interface{})
	}

	return obj[keys[len(keys)-1]]
}

func (c *Config) GetString(key string) string {
	value := c.getNested(key)
	if value == nil {
		return ""
	}
	return value.(string)
}

func (c *Config) GetInt(key string) int {
	return int(c.GetFloat(key))
}

func (c *Config) GetFloat(key string) float64 {
	value := c.getNested(key)
	if value == nil {
		return 0
	}
	return value.(float64)
}

func (c *Config) GetBool(key string) bool {
	value := c.getNested(key)
	if value == nil {
		return false
	}
	return value.(bool)
}

func (c *Config) GetArray(key string) []interface{} {
	value := c.getNested(key)
	if value == nil {
		return []interface{}{nil}
	}
	return value.([]interface{})
}

func (c *Config) GetConfigObj(key string) *Config {
	value := c.getNested(key)
	if value == nil {
		return nil
	}

	subConfig := newConfig()
	subConfig.parseFromMap(value.(map[string]interface{}))

	return subConfig
}

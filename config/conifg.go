package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	DBConfig    DBConfig    `yaml:"database"`
	SrvConfig   SrvConfig   `yaml:"server"`
	CacheConfig CacheConfig `yaml:"cache"`
}

func New() *Config {
	return &Config{}
}

func (c *Config) LoadFile(filename string) error {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(f, c)
	if err != nil {
		return err
	}

	return nil
}

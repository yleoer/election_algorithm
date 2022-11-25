package config

import "github.com/spf13/viper"

type Config struct {
	SvcCfg   SvcConfig      `json:"svcCfg"`
	RedGroup RedGroupConfig `json:"redGroup"`
}

type SvcConfig struct {
	Name   string `json:"name"`
	DbFqdn string `json:"dbFqdn"`
	DbPort uint16 `json:"dbPort"`
}

type RedGroupConfig struct {
	KeepAliveInterval  uint32       `json:"keepAliveInterval,omitempty"`
	KeepAliveRespTime  uint32       `json:"keepAliveRespTime,omitempty"`
	KeepAliveConseFail uint8        `json:"keepAliveConseFail,omitempty"`
	AutoReversionDelay uint32       `json:"autoReversionDelay,omitempty"`
	BaseRouteCost      uint32       `json:"baseRouteCost"`
	Nodes              []NodeConfig `json:"nodes"`
}

type NodeConfig struct {
	Name     string `json:"name"`
	Priority uint32 `json:"priority"`
	Fqdn     string `json:"fqdn"`
	SbiPort  uint16 `json:"sbiPort"`
}

func NewConfig() (Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath("/etc/config")

	var config Config

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}

func (c *Config) SbiPort() uint16 {
	var port uint16 = 9000
	for _, node := range c.RedGroup.Nodes {
		if node.Name == c.SvcCfg.Name {
			port = node.SbiPort
		}
	}
	return port
}

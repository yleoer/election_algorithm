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

func NewConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath("/etc/config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
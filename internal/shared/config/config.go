package config

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Name, Version, Address, Env string
	HttpPort                    int
	Database                    Database
}

type Database struct {
	Username, Password, Name, Schema, Host       string
	Port, MaxIdleConnections, MaxOpenConnections int
	DebugMode                                    bool
}

func NewConfig(path string) *Config {
	fmt.Println("Try new config..")
	viper.SetConfigFile(path + "/config.json")
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	conf := Config{}
	err := viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}

	configuration, _ := json.Marshal(conf)
	fmt.Println(string(configuration))
	return &conf
}

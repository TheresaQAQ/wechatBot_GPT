package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

func Loading(config *Config) (*Config, error) {
	vip := viper.New()
	vip.AddConfigPath("./")
	vip.SetConfigName("configuration")
	vip.SetConfigType("yaml")

	if err := vip.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件")
			return nil, errors.New("找不到配置文件")
		} else {
			fmt.Println("配置文件出错..")
			return nil, errors.New("配置文件出错")
		}
	}

	if err := vip.ReadInConfig(); err != nil {
		panic(err)
	}

	vip.Unmarshal(&config)
	return config, nil
}

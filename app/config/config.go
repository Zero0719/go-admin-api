package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Cfg *Config

type Config struct {
	Server ServerConfig `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Jwt JwtConfig `yaml:"jwt"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
	Mode string `yaml:"mode"`
}

type DatabaseConfig struct {
	Driver string `yaml:"driver"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset string `yaml:"charset"`
}

type JwtConfig struct {
	Secret string `yaml:"secret"`
	Expire int `yaml:"expire"`
}


func Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		return fmt.Errorf("解析配置失败: %w", err)
	}

	return nil
}
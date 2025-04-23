package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig      `yaml:"app" mapstructure:"app"`
	Database DatabaseConfig `yaml:"database" mapstructure:"database"`
}

type AppConfig struct {
	Port int `mapstructure:"port"`
}

type DatabaseConfig struct {
	Driver string `yaml:"driver" mapstructure:"driver"`
	Source string `yaml:"source" mapstructure:"source"`
}

var Conf *Config

// LoadConfig 加载配置文件
func LoadConfig() error {

	// 设置配置文件路径和名称
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("读取配置文件失败: %v", err))
	}

	// 将配置文件内容解析到 Conf 变量中
	Conf = &Config{}
	err = viper.Unmarshal(Conf)
	if err != nil {
		panic(fmt.Sprintf("解析配置文件失败: %v", err))
	}

	return nil
}

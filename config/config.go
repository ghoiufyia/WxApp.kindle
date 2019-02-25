package config

import "github.com/jinzhu/configor"

type DatabaseConfig struct {
	Type         string `default:"mysql"`
	Host         string `default:"localhost"`
	Port         int    `default:"3306"`
	User         string `default:"gravitee"`
	Password     string `default:"gravitee"`
	DatabaseName string `default:"gravitee"`
	MaxIdleConns int    `default:"5"`
	MaxOpenConns int    `default:"5"`
}

type Config struct {
	Database	DatabaseConfig
	ServerPort    int  `default:"8080"`
	IsDevelopment bool `default:"True"`
}

var DefaultConfig = &Config{
	Database: DatabaseConfig{
		Type:         "mysql",
		Host:         "101.132.37.148",
		Port:         33061,
		User:         "root",
		Password:     "Sn93007997",
		DatabaseName: "kindle",
		MaxIdleConns: 5,
		MaxOpenConns: 5,
	},
	ServerPort:	8000,

}

func NewDefaultConfig() *Config {
	return DefaultConfig
}

func NewConfig(configFile string) * Config {
	if configFile != "" {
		config := &Config{}
		configor.Load(config, configFile)
		return config
	}

	return NewDefaultConfig()
}
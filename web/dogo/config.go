package dogo

import (
	// "github.com/jinzhu/configor"
	"errors"
	"io/ioutil"
	"strings"
)

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

type ServerConfig struct {
	ServerPort    int  `default:"8080"`
}

type Config struct {
	Database	DatabaseConfig
	Server    	ServerConfig
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
	Server:	ServerConfig {
		ServerPort:		8080,
	},
}

func NewDefaultConfig() *Config {
	return DefaultConfig
}

func NewConfig(configFile string) *Config {
	if configFile != "" {
		config := &Config{}
		err := parseFile(config, configFile)
		if err == nil {
			return config
		}
	}
	return NewDefaultConfig()
}

func parseFile(config interface{},file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	switch {
	case strings.HasSuffix(file,".json"):
		Log.Info("%s,type is %s",file,"json")
		unmarshalJSON(data,config)
	default :
		Log.Info("There is no valid config file,json,yml or tml is needed.")
		Log.Info("Now is using the default config...")
		return errors.New("invalid config file")
	}
	return nil
}

func unmarshalJSON(data,config interface{}) {

}
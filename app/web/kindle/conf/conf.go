package conf

import (
	"errors"
	// "io/ioutil"
	// "strings"
	// "encoding/json"
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/ghoiufyia/kindle/library/net/http/klove"
)

type Config struct {
	AppConfig	*AppConfig
	KloveConf	*klove.ServerConfig
	SessionConf	*klove.SessionConfig
}

type AppConfig struct {
	RootDir	string
}

var (
	confPath string
	Conf = &Config{}
)

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}
func Init() error {
	if confPath == "" {
		return errors.New("no conf file path")
	}
	_, err := toml.DecodeFile(confPath, &Conf)
		if err != nil {
			return err
		}
	return nil
}

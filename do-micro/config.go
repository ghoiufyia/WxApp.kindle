package do-micro

import (
	"errors"
	"io/ioutil"
	"strings"
	"encoding/json"
	// "fmt"
)

// func NewDefaultConfig() *AccountConfig {
// 	return DefaultConfig
// }

func NewConfig(config interface{}, file string) error {
	if file != "" {
		err := parseFile(config, file)
		// fmt.Printf("%+v",config)
		if err != nil {
			return err
		}
	}
	return nil
}

func parseFile(config interface{},file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	switch {
	case strings.HasSuffix(file,".json"):
		err = unmarshalJSON(config,data)
		if err != nil {
			return err
		}
	default :
		return errors.New("invalid config file")
	}
	return nil
}
//解析json数据
func unmarshalJSON(config interface{},data []byte) error{
	if json.Valid(data) != true {
		return errors.New("invalid json string")
	}
	err := json.Unmarshal(data,config)
	// fmt.Printf("%+v",config)
	if err != nil {
		return err
	}
	return nil
}
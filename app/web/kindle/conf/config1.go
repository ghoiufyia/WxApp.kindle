package conf

import (
	"errors"
	"io/ioutil"
	"strings"
	"encoding/json"
	// "fmt"
)

//解析配置文件config  *Config
func ParseFile(config interface{},file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	switch {
	case strings.HasSuffix(file,".json"):
		err = unmarshalJSON(data,config)
		if err != nil {
			return err
		}
	default :
		return errors.New("invalid config file")
	}
	return nil
}
//解析json数据
func unmarshalJSON(data []byte,config interface{}) error{
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
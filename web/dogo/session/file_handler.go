package session

import (
	"os"
	"github.com/ghoiufyia/WxApp.kindle/web/dogo/log"
)

func Open(savePath string,sessionName string) (interface{}) {
	return true
}
func Close() (interface{}) {
	return true
}
func Read(sessionId string) interface{} {
	path := "./storage/session"+"/"+sessionId
	fileInfo,err := os.Stat(path)
	fmt.Printf("%+v",fileInfo)
	fmt.Printf("%+v",err)
	return true
}
func Write(sessionId string,data interface{}) (interface{}) {
	path := "./storage/session"+"/"+sessionId
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		
	}
}
// func Destroy(sessionId string)(interface{}) {}
// func Gc(lifetime int32) {}
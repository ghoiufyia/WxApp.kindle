package sessionhandler

import (

)
//session接口handler标准
type Handler interface {
	Open(savePath string,sessionName string) (interface{})
	Close() (interface{})
	Read(sessionId string) (string,error)
	Write(sessionId string,data string) (error)
	Destroy(sessionId string)(error)
	Gc() (error)
}
package session

import (

)

type handler interface {
	Open(savePath string,sessionName string) (interface{})
	Close() (interface{})
	Read(sessionId string) interface{}
	Write(sessionId string,data interface{}) (interface{})
	Destroy(sessionId string)(interface{})
	Gc(lifetime int32)
}
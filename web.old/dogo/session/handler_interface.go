package session

import (

)

type handler interface {
	Open(savePath string,sessionName string) (interface{})
	Close() (interface{})
	Read(sessionId string) (string,error)
	Write(sessionId string,data string) (error)
	Destroy(sessionId string)(error)
	Gc() (error)
}
package dogo

import (
	"log"
	"os"
)

var (
	var DoLog *log.Logger
)

func init() {
	var DoLog = log.New(os.Stdout,"dolog:",log.Ldate|log.Ltime)
}

func Info(info string,v ...interface{}) {
	DoLog.SetPrefix("[info]")
	if v == nil || len(v) == 0 {
		d.log.Printf(info)
	} else {
		d.log.Printf(info,v...)
	}
}
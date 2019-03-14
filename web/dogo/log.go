package dogo

import (
	"log"
	"os"
)

var (
	Log *DoLog
)

func init() {
	Log = NewDoLog()
}

type DoLog struct {
	log *log.Logger
}
func NewDoLog() *DoLog {
	return &DoLog{
		log:log.New(os.Stdout,"dolog:",log.Ldate|log.Ltime),
	}
}

func (d *DoLog)Info(info string,v ...interface{}) {
	d.log.SetPrefix("[info]")
	if v == nil || len(v) == 0 {
		d.log.Printf(info)
	} else {
		d.log.Printf(info,v...)
	}
}
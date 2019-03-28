package dogo

import (
	"log"
	"os"
	"time"
	"fmt"
)

const (
	LevelDebug = 1
	LevelInfo = 2
	LevelNotice = 3
	LevelWarning = 4
	LevelError = 5
	LevelCritical = 6
	LevelAlert = 7
	LevelEmergency = 8
)

var (
	level = LevelDebug
	DoLog *log.Logger
)
func InitLog(cfg LogConfig) {
	SetLogLevel(cfg.Level)
	DoLog = log.New(os.Stdout,"",log.Ldate|log.Ltime|log.Lmicroseconds)
	if cfg.Driver == "file" {
		path := cfg.Path
		t := time.Now()
		fileName := fmt.Sprintf("%d%02d%02d.file",t.Year(),t.Month(),t.Day())
		filePath := path+fileName
		f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
		if err != nil {
			Error(err)
		}
		DoLog.SetOutput(f)
	}
}

func SetLogLevel(l int) {
	level = l
}

func Debug(v ...interface{}) {
	if level <= LevelDebug {
		DoLog.SetPrefix("[Debug]")
		DoLog.Printf("%v\n",v)
	}
}
func Info(v ...interface{}) {
	if level <= LevelInfo {
		DoLog.SetPrefix("[Info]")
		DoLog.Printf("%v\n",v)
	}
}
func Notice(v ...interface{}) {
	if level <= LevelNotice {
		DoLog.SetPrefix("[Notice]")
		DoLog.Printf("%v\n",v)
	}
}
func Warning(v ...interface{}) {
	if level <= LevelWarning {
		DoLog.SetPrefix("[Warning]")
		DoLog.Printf("%v\n",v)
	}
}
func Error(v ...interface{}) {
	if level <= LevelError {
		DoLog.SetPrefix("[Error]")
		DoLog.Printf("%v\n",v)
	}
}
func Critical(v ...interface{}) {
	if level <= LevelCritical {
		DoLog.SetPrefix("[Critical]")
		DoLog.Printf("%v\n",v)
	}
}
func Alert(v ...interface{}) {
	if level <= LevelAlert {
		DoLog.SetPrefix("[Alert]")
		DoLog.Printf("%v\n",v)
	}
}
func Emergency(v ...interface{}) {
	if level <= LevelEmergency {
		DoLog.SetPrefix("[Emergency]")
		DoLog.Printf("%v\n",v)
	}
}
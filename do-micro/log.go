package do-micro

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
	level = LevelError
	Log *log.Logger
)
func InitLog(cfg LogConfig) {
	SetLogLevel(cfg.Level)
	Log = log.New(os.Stdout,"",log.Ldate|log.Ltime|log.Lmicroseconds)
	if cfg.Driver == "file" {
		path := cfg.Path
		t := time.Now()
		fileName := fmt.Sprintf("%d%02d%02d.file",t.Year(),t.Month(),t.Day())
		filePath := path+fileName
		f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
		if err != nil {
			Error(err)
		}
		Log.SetOutput(f)
	}
}

func SetLogLevel(l int) {
	level = l
}

func Debug(v ...interface{}) {
	if level <= LevelDebug {
		Log.SetPrefix("[Debug]")
		Log.Printf("%v\n",v)
	}
}
func Info(v ...interface{}) {
	if level <= LevelInfo {
		Log.SetPrefix("[Info]")
		Log.Printf("%v\n",v)
	}
}
func Notice(v ...interface{}) {
	if level <= LevelNotice {
		Log.SetPrefix("[Notice]")
		Log.Printf("%v\n",v)
	}
}
func Warning(v ...interface{}) {
	if level <= LevelWarning {
		Log.SetPrefix("[Warning]")
		Log.Printf("%v\n",v)
	}
}
func Error(v ...interface{}) {
	if level <= LevelError {
		Log.SetPrefix("[Error]")
		Log.Printf("%v\n",v)
	}
}
func Critical(v ...interface{}) {
	if level <= LevelCritical {
		Log.SetPrefix("[Critical]")
		Log.Printf("%v\n",v)
	}
}
func Alert(v ...interface{}) {
	if level <= LevelAlert {
		Log.SetPrefix("[Alert]")
		Log.Printf("%v\n",v)
	}
}
func Emergency(v ...interface{}) {
	if level <= LevelEmergency {
		Log.SetPrefix("[Emergency]")
		Log.Printf("%v\n",v)
	}
}
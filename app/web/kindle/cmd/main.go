package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"github.com/ghoiufyia/kindle/app/web/kindle/conf"
	"github.com/ghoiufyia/kindle/app/web/kindle/internal/routes"
	"github.com/ghoiufyia/kindle/app/web/kindle/internal/template"
	"github.com/ghoiufyia/kindle/library/net/http/klove"
	"github.com/ghoiufyia/kindle/app/web/kindle/internal/resource"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Printf("conf.Init() error(%v)", err)
		panic(err)
	}
	engine := klove.NewWebEngine(conf.Conf.KloveConf)
	// engine.Ping(ping)

	sessionManager := klove.NewSessionManager(conf.Conf.SessionConf)
	
	//set route and http.Handler
	rg := routes.Init()

	//绑定sessionmanager
	rg.SetSessionmManager(sessionManager)
	engine.SetHandler(rg)

	//载入模板
	template.Init()

	//资源和数据库操作
	resource.New(conf.Conf)


	//开启运行服务
	if err := engine.Start(); err != nil {
		panic(err)
	}
	// listen signal
	c := make(chan os.Signal,1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-c
		log.Printf("engine get a signal %s", si.String())
		switch si {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}

}
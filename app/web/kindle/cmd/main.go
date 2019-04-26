package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"fmt"
	"github.com/ghoiufyia/kindle/app/web/kindle/conf"
	"github.com/ghoiufyia/kindle/app/web/kindle/internal/routes"
	"github.com/ghoiufyia/kindle/library/net/http/klove"
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
	fmt.Printf("%+v\n",rg)

	//绑定sessionmanager
	rg.SetSessionmManager(sessionManager)
	engine.SetHandler(rg)

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
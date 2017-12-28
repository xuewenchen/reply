package main

import (
	"flag"
	"kit/log"
	"os"
	"os/signal"
	"reply/config"
	"reply/service"
	"syscall"
)

func main() {
	// config
	flag.Parse()

	// log
	log.Init(config.Conf.Log)
	defer log.Close()

	// run server
	log.Info("service start")
	if err := service.Run(config.Conf); err != nil {
		log.Info("service fail")
		return
	}

	// exit signal
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	for s := range c {
		switch s {
		case syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT:
			log.Info("service end receive signal %v", s)
			service.UnRegisterEtcd() // 从etcd注摘除这个节点
			service.EndTracing()     // 关闭trace
			service.CloseService()   // 关闭service，channel等
			return
		default:
			log.Info("other")
		}
	}
	return
}

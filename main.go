package main

import (
	"os"
	"runtime"
	"os/signal"
	"syscall"
	"tsstream/config"
	"tsstream/controller"
	"tsstream/monitor"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	runtime.GOMAXPROCS(runtime.NumCPU())

	isRun := false
	if "normal" == config.AppConf.RunMode {
		isRun = true
		controller.RunNormal()
	}

	if "extra" == config.AppConf.RunMode {
		isRun = true
		controller.RunExtra()
	}

	if "monitor" == config.AppConf.RunMode {
		isRun = true
		controller.RunMonitor()
	}

	go monitor.RunMonitorTimer()
	if isRun {
		chSig := make(chan os.Signal)
		signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
		<-chSig
	} else {
		log.Error("Didn't support run style,must is (normal or extra or monitor)")
	}

	return 
}
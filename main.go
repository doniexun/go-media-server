package main

import (
	"github.com/songshenyi/go-media-server/server"
	"os"
	"os/signal"
	"syscall"
	"runtime"
	"github.com/songshenyi/go-media-server/logger"
	"github.com/songshenyi/go-media-server/application"
	"github.com/songshenyi/go-media-server/core"
	agent_manager "github.com/songshenyi/go-media-server/agent/manager"
)

func signalHandle(){
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan)
	for{
		s:= <- signalChan
		logger.Infof("recv signal %d", s)
		switch s{
		case syscall.SIGTERM:
			fallthrough
		case syscall.SIGQUIT:
			buf :=make([]byte, 1<<20)
			runtime.Stack(buf, true)
			logger.Infof("killed by signal %d", s)
			logger.Infof("goroutine stack \n%s", buf)
			return
		}
	}
}

func main(){
	logger.InitLaunchLog()
	logger.Info("Server Start")
	logger.InitAccessLog("config/access.xml")
	httpServer := server.NewHttpServer(8888)
	application.AddHandle(httpServer)
	httpServer.Start()
	ctx := core.NewContext()
	agent_manager.Manager= agent_manager.NewManager(ctx)
	signalHandle()
}

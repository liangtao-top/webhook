package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"webhook/src/controller"
	"webhook/src/global"
	"webhook/src/global/enum"
	"webhook/src/global/http"
	"webhook/src/logger"
	"webhook/src/util"
)

func main() {
	// 初始化指令
	global.InitCommand()
	// 解析指令
	flag.Parse()
	// 初始化配置
	global.InitConfig()
	// 初始化日志
	logger.InitLogger()
	// 启动Banner
	global.Welcome()
	// 打印指令
	logger.Debug("\n", util.ToJsonString(enum.CMD, true))
	// Http 服务
	go http.Start()
	// 任务定时器
	go controller.Cron()

	// 监听退出序号
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	// 设置要接收的信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		logger.Info(sig)
		done <- true
	}()
	<-done
	logger.Infof("exiting webhook.")
}

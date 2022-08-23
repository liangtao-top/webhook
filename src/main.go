package main

import (
	"flag"
	"webhook/src/controller"
	"webhook/src/global"
	"webhook/src/global/enum"
	"webhook/src/global/http"
	"webhook/src/logger"
	"webhook/src/util"
)

func main() {
	// 初始化配置
	global.InitConfig()
	// 初始化日志
	logger.InitLogger()
	// 启动Banner
	global.Welcome()
	// 初始化指令
	global.InitCommand()
	// 解析指令
	flag.Parse()
	logger.Debug("\n", util.ToJsonString(enum.CMD, true))
	// 任务定时器
	go controller.Cron()
	// Http 服务
	http.Start()
}

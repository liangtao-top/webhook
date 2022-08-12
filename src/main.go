package main

import (
	"flag"
	"webhook/src/global"
	"webhook/src/global/enum"
	"webhook/src/global/http"
)

func main() {
	global.Welcome()
	// 初始化配置
	global.IniConfigFromYaml()
	// 解析命令行
	initCommand()
	flag.Parse()
	// Http 服务
	http.Start()
}

func initCommand() {
	flag.Uint64Var(&enum.CMD.Port, "p", 0, "Http服务端口")
	flag.Uint64Var(&enum.CMD.Port, "port", 0, "Http服务端口")
	flag.StringVar(&enum.CMD.Sh, "sh", "", "WebHook预执行指令")
	flag.StringVar(&enum.CMD.Sh, "cmd", "", "WebHook预执行指令")
	flag.StringVar(&enum.CMD.File, "f", "", "WebHook预执行文件")
	flag.StringVar(&enum.CMD.File, "file", "", "WebHook预执行文件")
}

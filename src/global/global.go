package global

import (
	"flag"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"runtime"
	"webhook/src/global/enum"
	"webhook/src/logger"
)

func InitConfig() {
	file, err := ioutil.ReadFile(enum.RootPath + string(os.PathSeparator) + "config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(file, enum.CONFIG)
	if err != nil {
		panic(err)
	}
}

func InitCommand() {
	flag.BoolVar(&enum.CMD.Daemon, "d", false, "进程守护")
	flag.BoolVar(&enum.CMD.Daemon, "daemon", false, "进程守护")
	flag.Uint64Var(&enum.CMD.Port, "p", 0, "Http服务端口")
	flag.Uint64Var(&enum.CMD.Port, "port", 0, "Http服务端口")
	flag.StringVar(&enum.CMD.Context, "context", "", "Http服务上下文路径")
	flag.StringVar(&enum.CMD.Sh, "sh", "", "WebHook预执行指令")
	flag.StringVar(&enum.CMD.Sh, "cmd", "", "WebHook预执行指令")
	flag.StringVar(&enum.CMD.File, "f", "", "WebHook预执行文件")
	flag.StringVar(&enum.CMD.File, "file", "", "WebHook预执行文件")
	flag.StringVar(&enum.CMD.Cron, "c", "", "定时器预执行文件")
	flag.StringVar(&enum.CMD.Cron, "cron", "", "定时器预执行文件")
	flag.StringVar(&enum.CMD.Token, "token", "", "用户的 WebHook 密钥")
}

func Welcome() {
	println("███╗   ███╗███████╗ █████╗\n████╗ ████║██╔════╝██╔══██╗\n██╔████╔██║███████╗███████║\n██║╚██╔╝██║╚════██║██╔══██║\n██║ ╚═╝ ██║███████║██║  ██║\n╚═╝     ╚═╝╚══════╝╚═╝  ╚═╝\n")
	println("\033[91;1mMSA 自动部署工具 v" + enum.VERSION + "\033[0m\r\n")
	// 当前版本
	logger.Infof("start webhook v%s ", enum.VERSION)
	logger.Infof("go version %s ", runtime.Version())
}

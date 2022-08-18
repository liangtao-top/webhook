package global

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"runtime"
	"webhook/src/global/enum"
	"webhook/src/logger"
)

func IniConfigFromYaml() {
	file, err := ioutil.ReadFile(enum.RootPath + string(os.PathSeparator) + "config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(file, enum.CONFIG)
	if err != nil {
		panic(err)
	}
}

func Welcome() {
	println("███╗   ███╗███████╗ █████╗\n████╗ ████║██╔════╝██╔══██╗\n██╔████╔██║███████╗███████║\n██║╚██╔╝██║╚════██║██╔══██║\n██║ ╚═╝ ██║███████║██║  ██║\n╚═╝     ╚═╝╚══════╝╚═╝  ╚═╝\n")
	println("\033[91;1mMSA 自动部署工具 v" + enum.VERSION + "\033[0m\r\n")
	// 当前版本
	logger.Infof("start webhook v%s ", enum.VERSION)
	logger.Infof("go version %s ", runtime.Version())
}

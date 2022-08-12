package global

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"runtime"
	"webhook/src/global/enum"
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
	// 当前版本
	fmt.Printf("start webhook v%s \n", enum.VERSION)
	fmt.Printf("go version %s \n", runtime.Version())
}

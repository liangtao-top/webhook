package controller

import (
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"webhook/src/global/enum"
	"webhook/src/logger"
	"webhook/src/util"
)

func GetItems(writer http.ResponseWriter, request *http.Request) {
	var defaultConfig = enum.CONFIG.WebHook
	result := util.Result{Writer: writer}
	ContentType := request.Header.Get("Content-Type")
	UserAgent := request.Header.Get("User-Agent")
	Token := request.Header.Get("X-Gitee-Token")
	res := "success"
	if defaultConfig.ContentType == ContentType && defaultConfig.UserAgent == UserAgent && defaultConfig.Token == Token {
		if len(enum.CMD.Sh) > 0 {
			logger.Info("sh:", enum.CMD.Sh)
			shell := strings.Fields(enum.CMD.Sh)
			bytes, err := Cmd(shell[0], shell[1:]...)
			if err != nil {
				logger.Errorf("sh exec err:%v", err)
			}
			res = string(bytes)
			logger.Info("output:\n", res)
		}
		if len(enum.CMD.File) > 0 {
			logger.Info("file: ", enum.CMD.File)
			err := os.Chmod(enum.CMD.File, 0777)
			if err != nil {
				logger.Errorf("file chmod err:%v", err)
			}
			bytes, err := Cmd(enum.CMD.File)
			if err != nil {
				logger.Errorf("file exec err:%v", err)
			}
			res = string(bytes)
			logger.Info("output:\n", res)
		}
	}
	result.Success(res)
}

func Cmd(name string, arg ...string) (out []byte, err error) {
	//// 5秒超时
	//ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(15)*time.Second)
	//cmd := exec.CommandContext(ctx, name, arg...)
	cmd := exec.Command(name, arg...)
	stdout, err := cmd.StdoutPipe()
	if err != nil { //获取输出对象，可以从该对象中读取输出结果
		return nil, err
	}
	defer stdout.Close() // 保证关闭输出流
	//defer func() {
	//	cancelFunc()
	//	_ = stdout.Close()
	//	_ = cmd.Wait()
	//}()

	if err := cmd.Start(); err != nil { // 运行命令
		return nil, err
	}

	if opBytes, err := ioutil.ReadAll(stdout); err != nil { // 读取输出结果
		return nil, err
	} else {
		return opBytes, nil
	}
}

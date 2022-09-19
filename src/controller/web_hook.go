package controller

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"sync"
	"webhook/src/global/enum"
	"webhook/src/logger"
	"webhook/src/util"
)

func GetItems(writer http.ResponseWriter, request *http.Request) {
	var defaultConfig = enum.CONFIG.WebHook
	result := util.Result{Writer: writer}
	ContentType := request.Header.Get("Content-Type")
	UserAgent := request.Header.Get("User-Agent")
	Timestamp := request.Header.Get("X-Gitee-Timestamp")
	Token := request.Header.Get("X-Gitee-Token")
	if len(enum.CMD.Token) > 0 {
		enum.CONFIG.WebHook.Token = enum.CMD.Token
	}
	sign := util.GenHmacSha256(Timestamp+"\n"+enum.CONFIG.WebHook.Token, enum.CONFIG.WebHook.Token)
	logger.Debugf("%s|%s", enum.CMD.Token, Token, sign)
	b := defaultConfig.ContentType == ContentType && defaultConfig.UserAgent == UserAgent && Token == sign
	if b {
		result.Success("success")
	} else {
		result.Error(enum.CALL_ERROR, "fail")
	}
	//defer request.Body.Close()
	if b {
		if len(enum.CMD.Sh) > 0 {
			logger.Infof("exec %s", enum.CMD.Sh)
			ctx, cancel := context.WithCancel(context.Background())
			defer func() {
				cancel()
			}()
			err := Command(ctx, enum.CMD.Sh)
			if err != nil {
				logger.Error(err)
				return
			} else {
				logger.Infof("exec %s complete", enum.CMD.Sh)
			}
		}
		if len(enum.CMD.File) > 0 {
			logger.Infof("exec %s", enum.CMD.File)
			if err := os.Chmod(enum.CMD.File, 0777); err != nil {
				logger.Errorf("file chmod err:%v", err)
			}
			ctx, cancel := context.WithCancel(context.Background())
			defer func() {
				cancel()
			}()
			err := Command(ctx, enum.CMD.File)
			if err != nil {
				logger.Error(err)
				return
			}
			logger.Infof("exec %s complete", enum.CMD.File)
		}

	}
}

func Command(ctx context.Context, cmd string) error {
	//c := exec.CommandContext(ctx, "cmd", "/C", cmd) // windows
	c := exec.CommandContext(ctx, "bash", "-c", cmd) // mac linux
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := c.StderrPipe()
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	// 因为有2个任务, 一个需要读取stderr 另一个需要读取stdout
	wg.Add(2)
	go read(ctx, &wg, stderr)
	go read(ctx, &wg, stdout)
	// 这里一定要用start,而不是run 详情请看下面的图
	err = c.Start()
	// 等待任务结束
	wg.Wait()
	return err
}

func read(ctx context.Context, wg *sync.WaitGroup, std io.ReadCloser) {
	reader := bufio.NewReader(std)
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			logger.Infof("read 关闭")
			return
		default:
			readString, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				logger.Infof("read 关闭2")
				return
			}
			fmt.Print(readString)
		}
	}
}

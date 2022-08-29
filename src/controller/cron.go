package controller

import (
	"context"
	"os"
	"strconv"
	"strings"
	"time"
	"webhook/src/global/enum"
	"webhook/src/logger"
)

func Cron() {
	if enum.CMD.Cron != "" {
		for _, cron := range strings.Split(enum.CMD.Cron, ",") {
			cron = strings.Trim(cron, "")
			if cron != "" {
				arr := strings.Split(cron, ":")
				logger.Infof("定时任务执行间隔 %+v 秒 File:%s ", arr[1], arr[0])
				parseInt, err := strconv.ParseInt(arr[1], 10, 64)
				if err != nil {
					return
				}
				ticker := time.NewTicker(time.Duration(int64(time.Second) * parseInt))
				for range ticker.C {
					Affair(arr[0])
				}
			}
		}
	}
}

func Affair(file string) {
	logger.Infof("start exec %s ", file)
	if err := os.Chmod(file, 0777); err != nil {
		logger.Errorf("file chmod err:%v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()
	err := Command(ctx, file)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Infof("complete exec %s ", file)
}

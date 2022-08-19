package controller

import (
	"context"
	"os"
	"time"
	"webhook/src/global/enum"
	"webhook/src/logger"
)

func Cron() {
	logger.Infof("定时任务执行间隔 %+v 秒 File:%s ", enum.CMD.Ticker, enum.CMD.Cron)
	ticker := time.NewTicker(time.Duration(int64(time.Second) * enum.CMD.Ticker))
	defer ticker.Stop()
	for range ticker.C {
		Affair(enum.CMD.Cron)
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

/*
 * Copyright 1999-2020 Alibaba Group Holding Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strconv"
	"sync"
	"time"
	"webhook/src/config"
	"webhook/src/global/enum"
)

var (
	logger  Logger
	logLock sync.RWMutex
)

var levelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
}

type Config struct {
	Level            string
	Sampling         *SamplingConfig
	LogRollingConfig *lumberjack.Logger
}

type SamplingConfig struct {
	Initial    int
	Thereafter int
	Tick       time.Duration
}

type MsaLogger struct {
	Logger
}

// Logger is the interface for Logger types
type Logger interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Debug(args ...interface{})

	Infof(fmt string, args ...interface{})
	Warnf(fmt string, args ...interface{})
	Errorf(fmt string, args ...interface{})
	Debugf(fmt string, args ...interface{})
}

func init() {
	zapLoggerConfig := zap.NewDevelopmentConfig()
	zapLoggerEncoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	zapLoggerConfig.EncoderConfig = zapLoggerEncoderConfig
	zapLogger, _ := zapLoggerConfig.Build(zap.AddCaller(), zap.AddCallerSkip(1))
	SetLogger(&MsaLogger{zapLogger.Sugar()})
}

func BuildLoggerConfig(clientConfig *config.Logger) Config {
	loggerConfig := Config{
		Level: clientConfig.LogLevel,
	}
	currentTime := time.Now()
	if enum.CMD.Port > 0 {
		enum.CONFIG.Server.Port = enum.CMD.Port
	}
	loggerConfig.LogRollingConfig = &lumberjack.Logger{
		Filename: clientConfig.LogDir + string(os.PathSeparator) + strconv.FormatUint(enum.CONFIG.Server.Port, 10) + string(os.PathSeparator) + currentTime.Format("2006-01-02 15_04_05.999") + ".log",
	}
	return loggerConfig
}

func InitLogger() {
	logLock.Lock()
	defer logLock.Unlock()
	logger, _ = InitMsaLogger(BuildLoggerConfig(enum.CONFIG.Logger))
}

// InitMsaLogger is init msa default logger
func InitMsaLogger(config Config) (Logger, error) {
	logLevel := getLogLevel(config.Level)
	encoder := getEncoder()
	writer := config.getLogWriter()
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoder),
		zapcore.NewMultiWriteSyncer(writer, zapcore.AddSync(os.Stdout)), logLevel)
	zaplogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return &MsaLogger{zaplogger.Sugar()}, nil
}

func getLogLevel(level string) zapcore.Level {
	if zapLevel, ok := levelMap[level]; ok {
		return zapLevel
	}
	return zapcore.InfoLevel
}

func getEncoder() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

//SetLogger sets logger for sdk
func SetLogger(log Logger) {
	logLock.Lock()
	defer logLock.Unlock()
	logger = log
}

func GetLogger() Logger {
	logLock.RLock()
	defer logLock.RUnlock()
	return logger
}

// getLogWriter get Lumberjack writer by LumberjackConfig
func (c *Config) getLogWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(c.LogRollingConfig)
}

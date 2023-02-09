package logData

import (
	project "Notes-go-project/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

// 日志记录到文件
func WriterLog() *logrus.Logger {
	// 实例化
	logger := logrus.New()
	project.ReadConfig()
	//配置MySQL连接参数
	fileInfo := project.ConfigToml.LogFile
	logFilePath := fileInfo.LogRouterPath
	logFileName := fileInfo.FileName
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	// 设置日志输出到什么地方去
	// 将日志输出到标准输出，就是直接在控制台打印出来。
	logger.SetOutput(os.Stdout)
	// 设置为true则显示日志在代码什么位置打印的
	logger.SetReportCaller(true)
	// 写入文件
	src, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logrus.Info(err)
	}
	// 设置输出
	logger.Out = src
	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",
		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),
		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// 新增 Hook
	logger.AddHook(lfHook)
	return logger
}

package log

import (
	"io"
	"os"
	"path"
	"sync"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var logger *log.Logger = nil
var once sync.Once

// IsDirExist 文件夹是否存在
func IsDirExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// InitLogger 初始化logger
func InitLogger() {
	basePath := viper.GetString("log.path")

	exist, err := IsDirExist(basePath)
	if err != nil {
		panic(err)
	}
	if !exist {
		os.MkdirAll(basePath, os.ModePerm)
	}

	logger = log.New()

	writer, err := rotatelogs.New(
		basePath+"log_%Y%m%d.log",
		rotatelogs.WithLinkName(path.Join(basePath, "curLog.log")),
		rotatelogs.WithMaxAge(time.Duration(3*24)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Second),
	)
	if err != nil {
		panic(err)
	}

	logger.SetOutput(io.MultiWriter(writer, os.Stdout))
	logger.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
	logger.SetReportCaller(true)
}

// GetLogger 获取logger
func GetLogger() *log.Logger {
	if logger == nil {
		once.Do(func() {
			InitLogger()
		})
	}

	return logger
}

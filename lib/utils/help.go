package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"github.com/lestrrat-go/file-rotatelogs"
	"os"
	"path"
	"go_grpc_gorm_micro/lib/global"
	"time"
)


// 批量创建文件夹
func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			global.CURD_LOG.Debug("create directory" + v)
			err = os.MkdirAll(v, os.ModePerm)
			if err != nil {
				global.CURD_LOG.Error("create directory"+ v, zap.Any(" error:", err))
			}
		}
	}
	return err
}


// 字符串是否在数组里
func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// 路径是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// GetWriteSyncer zap logger中加入file-rotatelogs
func GetWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := rotatelogs.New(
		path.Join(global.CURD_CONFIG.Zap.Director, "%Y-%m-%d.log"),
		rotatelogs.WithLinkName(global.CURD_CONFIG.Zap.LinkName),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if global.CURD_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
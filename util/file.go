package util

import (
	"go.uber.org/zap"
	"os"
)

func CreatFile(name string) (*os.File, error) {
	return os.Create(name)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		// 创建文件夹
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			zap.L().Error("mkdir failed![]\n", zap.Error(err))
		} else {
			return true, nil
		}
	}
	return false, err
}

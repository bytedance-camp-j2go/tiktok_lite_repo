package util

import (
	"encoding/json"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"tiktok-lite/global"
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

// Exists determine whether the file exists
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// IsDir determine whether the file is dir
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// CreatNestedFile create nested file
func CreatNestedFile(path string) (*os.File, error) {
	basePath := filepath.Dir(path)
	if !Exists(basePath) {
		err := os.MkdirAll(basePath, 0700)
		if err != nil {
			global.Logf.Errorf("can't create foler，%s\n", err)
			return nil, err
		}
	}
	return os.Create(path)
}

// WriteToJson write struct to json file
func WriteToJson(src string, conf interface{}) bool {
	data, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		global.Logf.Errorf("failed convert Conf to []byte:%s\n", err.Error())
		return false
	}
	err = ioutil.WriteFile(src, data, 0777)
	if err != nil {
		global.Logf.Errorf("failed to write json file:%s\n", err.Error())
		return false
	}
	return true
}

func ParsePath(path string) string {
	path = strings.TrimRight(path, "/")
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return path
}

func RemoveLastSlash(path string) string {
	if len(path) > 1 {
		return strings.TrimSuffix(path, "/")
	}
	return path
}

func Dir(path string) string {
	idx := strings.LastIndex(path, "/")
	if idx == 0 {
		return "/"
	}
	if idx == -1 {
		return path
	}
	return path[:idx]
}

func Base(path string) string {
	idx := strings.LastIndex(path, "/")
	if idx == -1 {
		return path
	}
	return path[idx+1:]
}

func Join(elem ...string) string {
	res := path.Join(elem...)
	if res == "\\" {
		res = "/"
	}
	return res
}

func Split(p string) (string, string) {
	return path.Split(p)
}

func Ext(name string) string {
	return strings.TrimPrefix(path.Ext(name), ".")
}

// ParentPath2Suffix 去除 parent path 前缀 "/"
func ParentPath2Suffix(name string) string {
	if name[0] == '/' {
		return name[1:]
	}
	return name
}

package util

import (
	"Whisper_Record/util/file"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// GetCurrPath 获取程序运行路径
func GetCurrPath() string {
	f, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(f)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return ret
}

// GetCfgPath ...
func GetCfgPath() string {
	cfgPath := fmt.Sprintf("%s/configs", GetCurrPath())
	if !file.IsExist(cfgPath) {
		err := os.Mkdir(cfgPath, os.ModePerm)
		if err != nil {
			panic("创建应用配置目录失败: " + err.Error())
		}
	}
	return cfgPath
}

// GetLogPath ...
func GetLogPath() string {
	logPath := fmt.Sprintf("%s/logs", GetCurrPath())
	if !file.IsExist(logPath) {
		err := os.Mkdir(logPath, os.ModePerm)
		if err != nil {
			panic("创建应用日志目录失败: " + err.Error())
		}
	}
	return logPath
}

// MkDir ...
func MkDir(path string) string {
	// 如果是文件 并且文件存在 会获取文件路径 直接返回
	if file.IsFile(path) {
		path = file.GetFileDir(path)
		return path
	}
	//如果不是文件并且不存在 如果是含有带“.”的路径,一定要确保是路径而不是文件名,
	//所以确保传入的要么是有效的文件路径，要么是文件夹，不要传入不存在的文件路径
	if !file.IsExist(path) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			panic("创建目录失败: " + err.Error())
		}
	}
	return path
}

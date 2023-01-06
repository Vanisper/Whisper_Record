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
	cfgPath := fmt.Sprintf("%s/Configs", GetCurrPath())
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
	logPath := fmt.Sprintf("%s/Logs", GetCurrPath())
	if !file.IsExist(logPath) {
		err := os.Mkdir(logPath, os.ModePerm)
		if err != nil {
			panic("创建应用日志目录失败: " + err.Error())
		}
	}
	return logPath
}

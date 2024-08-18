package initialize

import (
	"path"
	"runtime"

	"github.com/zhaohaihang/user_api/config"
	"github.com/zhaohaihang/user_api/global"
)

// InitFilePath 初始化全局文件路径
func InitFilePath(nacosConfig string) {
	basePath := getCurrentAbPathByCaller()
	global.FileConfig = &config.FileConfig{
		ConfigFile: basePath + "/config-"+nacosConfig+".yaml",
		LogFile:    basePath + "/log",
	}
}

// getCurrentAbPathByCaller 返回当前路径
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(2)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

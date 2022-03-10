package env

import (
	"os"
	"strings"
)

const (
	DEV        = "dev"
	TEST       = "test"
	PRERELEASE = "prerelease"
	ONLINE     = "production"
)

var env Environment

// Environment 存储环境相关变量
type Environment struct {
	Env string
}

// InitializeEnv 初始化环境相关参数
func InitializeEnv() {
	cmdMap := ResolveCMDParams()

	val, ok := cmdMap["env"]
	if !ok {
		// panic(errors.New("there is no 'env' param in cmd, please give it"))
		env.Env = DEV
	} else {
		env.Env = val
	}
}

// GetEnv 获取 env
func GetEnv() string {
	return env.Env
}

// ResolveCMDParams 解析命令行参数
func ResolveCMDParams() map[string]string {

	argMap := make(map[string]string)
	for _, v := range os.Args {
		kvs := strings.Split(v, "=")
		if len(kvs) < 2 {
			continue
		}
		argMap[kvs[0]] = kvs[1]
	}

	return argMap
}

package conf

import (
	"github.com/spf13/viper"
	"github.com/zhyeah/goyeahsk/basic/env"
)

// InitializeConfig 初始化配置
func InitializeConfig() error {
	// 本地配置初始化
	env := env.GetEnv()
	viper.SetConfigName(env)
	viper.AddConfigPath("basic/conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// 远端配置中心初始化

	return nil
}

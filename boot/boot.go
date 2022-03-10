package boot

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	autoreg "github.com/zhyeah/gin-autoreg"
	"github.com/zhyeah/goyeahsk/basic/conf"
	"github.com/zhyeah/goyeahsk/basic/dao"
	"github.com/zhyeah/goyeahsk/basic/env"
	"github.com/zhyeah/goyeahsk/boot/http/intercepter"
	"github.com/zhyeah/goyeahsk/boot/http/plugin"
)

// Router router
var Router *autoreg.AutoRouter
var engine *gin.Engine

// Boot fddata goweb启动流程
func Boot() error {
	// 初始化环境信息
	env.InitializeEnv()

	// 初始化配置
	conf.InitializeConfig()

	// 初始化dao
	dao.InitializeDao()

	// 初始化任务调度
	// job.InitializeSchedule()

	intercepter.AddDefaultPreIntercepters()
	intercepter.AddDefaultPostIntercepters()

	// 注册路由
	engine = gin.Default()
	Router = autoreg.GetAutoRouter()
	plugin.Register(Router)
	err := Router.RegisterRoute(&autoreg.AutoRouteConfig{
		BaseUrl: viper.GetString("baseurl"),
		Engine:  engine,
	})
	if err != nil {
		return err
	}

	// 设置模式
	if env.GetEnv() != env.DEV && env.GetEnv() != env.TEST {
		gin.SetMode(gin.ReleaseMode)
	}
	return engine.Run(fmt.Sprintf(":%d", viper.GetInt("serverport")))
}

package plugin

import (
	"github.com/spf13/viper"
	autoRouter "github.com/zhyeah/gin-autoreg"
	autoreg "github.com/zhyeah/gin-autoreg"
	autoregData "github.com/zhyeah/gin-autoreg/data"
	"github.com/zhyeah/goyeahsk/basic/env"
)

// Plugin web启动事件插件
type Plugin interface {
	SetProjectInfo(projectInfo *ProjectInfo)
	GetOnStartAction() func(routerContext *autoregData.RouterContext)
	GetOnFinishedAction() func(routerContext *autoregData.RouterContext)
}

func getProjectInfo() *ProjectInfo {
	projectInfo := &ProjectInfo{}
	projectInfo.ProjectName = viper.GetString("modulename")
	projectInfo.Env = env.GetEnv()
	projectInfo.Owner = viper.GetString("owner")
	// projectInfo.ServiceMeshInfo = &ServiceMeshInfo{
	// }
	return projectInfo
}

// Register 注册plugins
func Register(autoRouter *autoreg.AutoRouter) {
	projectInfo := getProjectInfo()
	register(projectInfo)
}

// Register 注册
func register(projectInfo *ProjectInfo) {
	// 注册http注册路由, 结束注册路由事件
	collector := GetPluginCollector()
	for i := range collector.Plugins {
		collector.Plugins[i].SetProjectInfo(projectInfo)
		onstart := collector.Plugins[i].GetOnStartAction()
		if onstart != nil {
			autoRouter.GetAutoRouter().AddStartAction(onstart)
		}
		onfinished := collector.Plugins[i].GetOnFinishedAction()
		if onfinished != nil {
			autoRouter.GetAutoRouter().AddFinishedAction(onfinished)
		}
	}
}

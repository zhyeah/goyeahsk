package intercepter

import (
	"github.com/gin-gonic/gin"
	"github.com/zhyeah/gin-autoreg/intercepter"
	"github.com/zhyeah/goyeahsk/basic/log"
)

// AddDefaultPreIntercepters 注册trace前后置拦截器
func AddDefaultPreIntercepters() {
	mgr := intercepter.GetIntercepterManager()
	// before user code
	mgr.AddPreIntercepters(SetTrace)

	// panic global handler
	gpi := intercepter.GetGlobalPanicIntercepter()
	gpi.AppendPanicIntercepter(func(ctx *gin.Context, panicInfo *intercepter.PanicInfo) {
		// log biz log
		log.GetLogger().Error(panicInfo)
	})

}

// AddDefaultPostIntercepters 注册默认后置拦截器
func AddDefaultPostIntercepters() {
	mgr := intercepter.GetIntercepterManager()
	// after user code
	mgr.AddPostIntercepters(ClearTrace)
}

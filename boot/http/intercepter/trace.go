package intercepter

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/tylerb/gls"
	"github.com/zhyeah/goyeahsk/basic/env"
	"github.com/zhyeah/goyeahsk/basic/log"
	"github.com/zhyeah/goyeahsk/basic/util"
)

// SetTrace 设置trace信息
func SetTrace(ctx *gin.Context) {
	// set trace
	traceId := ctx.GetHeader("traceid")
	u, _ := uuid.NewV4()
	spanId := u.String()
	lastSpanId := ctx.GetHeader("spanid")
	lastParentId := ctx.GetHeader("parentid")
	parentId := lastSpanId
	apiName := ctx.FullPath()
	staffName := ctx.GetHeader("staff")
	e := ctx.GetHeader("env")
	if e == "" {
		switch env.GetEnv() {
		case env.DEV:
			e = "Test"
		case env.PRERELEASE:
			e = "Prerelease"
		case env.ONLINE:
			e = "Production"
		case env.TEST:
			e = "Test"
		}
	}

	log.GetLogger().Debugf("[boot][%d] lastSpanId: %v, lastParentId: %v", util.CurGoroutineID(), lastSpanId, lastParentId)

	dataMap := map[string]interface{}{
		"traceId":      traceId,
		"apiName":      apiName,
		"parentId":     parentId,
		"spanId":       spanId,
		"lastSpanId":   lastSpanId,
		"lastParentId": lastParentId,
		"staff":        staffName,
		"env":          e,
	}
	util.SetGLSInfoWithData(dataMap)

	log.GetLogger().Debugf("[boot][%d] traceId: %v, apiName: %v", util.CurGoroutineID(), gls.Get("traceId"), gls.Get("apiName"))
}

// ClearTrace 清理trace信息
func ClearTrace(ctx *gin.Context) {
	gls.Cleanup()
}

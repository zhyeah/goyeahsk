package util

import (
	"github.com/gofrs/uuid"
	"github.com/tylerb/gls"
	"github.com/zhyeah/goyeahsk/basic/log"
)

// SetGLSInfoWithData 设置trace需要的gls数据
func SetGLSInfoWithData(dataMap map[string]interface{}) {
	// set traceId, spanId
	traceId := setUUID("traceId", &dataMap)
	spanId := setUUID("spanId", &dataMap)
	log.GetLogger().Info("dataMap: ", dataMap)

	log.GetLogger().Info("traceId loaded: ", traceId, " spanId: ", spanId)
	for k, v := range dataMap {
		gls.Set(k, v)
	}
}

func setUUID(field string, dataMap *map[string]interface{}) string {
	// set traceId
	fv := ""
	fvInter, ok := (*dataMap)[field]
	if !ok {
		u, _ := uuid.NewV4()
		fv = u.String()
		(*dataMap)[field] = fv
	} else {
		fv = fvInter.(string)
		if fv == "" {
			u, _ := uuid.NewV4()
			fv = u.String()
			(*dataMap)[field] = fv
		}
	}
	return fv
}

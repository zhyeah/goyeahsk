package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

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

// IsDirExist 文件夹是否存在
func IsDirExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// GetRandom 返回 [lower, upper) 区间的随机数
func GetRandom(lower, upper int) int {
	rand.Seed(int64(time.Now().UnixNano()))
	return lower + rand.Intn(upper-lower)
}

// GetIP 获取ip
func GetIP() (string, error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}

		}
	}
	return "", errors.New("cannot find legal ip")
}

// ConvertObjectToMap 将对象转换成map
func ConvertObjectToMap(data interface{}) (map[string]interface{}, error) {
	bts, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var ret map[string]interface{}
	err = json.Unmarshal(bts, &ret)
	return ret, err
}

// ConvertBoolToInt 转换bool为int
func ConvertBoolToInt(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

// 将json 格式化indent 方便查看
func IndenStruct(value interface{}) *bytes.Buffer {
	vals, err := json.MarshalIndent(value, "", "\t")
	if err != nil {
		return bytes.NewBufferString(err.Error())
	}
	return bytes.NewBuffer(vals)
}

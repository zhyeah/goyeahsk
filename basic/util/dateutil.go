package util

import (
	"regexp"
	"strings"
	"sync"
	"time"
)

const (
	stdLongMonth      = "January"
	stdMonth          = "Jan"
	stdNumMonth       = "1"
	stdZeroMonth      = "01"
	stdLongWeekDay    = "Monday"
	stdWeekDay        = "Mon"
	stdDay            = "2"
	stdUnderDay       = "_2"
	stdZeroDay        = "02"
	stdHour           = "15"
	stdHour12         = "3"
	stdZeroHour12     = "03"
	stdMinute         = "4"
	stdZeroMinute     = "04"
	stdSecond         = "5"
	stdZeroSecond     = "05"
	stdLongYear       = "2006"
	stdYear           = "06"
	stdPM             = "PM"
	stdpm             = "pm"
	stdTZ             = "MST"
	stdISO8601TZ      = "Z0700"  // prints Z for UTC
	stdISO8601ColonTZ = "Z07:00" // prints Z for UTC
	stdNumTZ          = "-0700"  // always numeric
	stdNumShortTZ     = "-07"    // always numeric
	stdNumColonTZ     = "-07:00" // always numeric
)

var dateConstantHelperOnce sync.Once
var dateConstantHelper *DateConstantHelper

// DateConstantHelper 日期长量助手
type DateConstantHelper struct {
	TimeUnitMap map[string]int64
	FormatMap   map[string]string
}

// GetDateConstantHelper 获取日期格式相关常量帮助类单例
func GetDateConstantHelper() *DateConstantHelper {
	if dateConstantHelper == nil {
		dateConstantHelperOnce.Do(func() {
			dateConstantHelper = &DateConstantHelper{
				TimeUnitMap: make(map[string]int64),
				FormatMap:   make(map[string]string),
			}

			// time unit (ms)
			dateConstantHelper.TimeUnitMap["min"] = 60 * 1000
			dateConstantHelper.TimeUnitMap["10min"] = 10 * dateConstantHelper.TimeUnitMap["min"]
			dateConstantHelper.TimeUnitMap["hour"] = 60 * dateConstantHelper.TimeUnitMap["min"]
			dateConstantHelper.TimeUnitMap["day"] = 24 * dateConstantHelper.TimeUnitMap["hour"]
			dateConstantHelper.TimeUnitMap["week"] = 7 * dateConstantHelper.TimeUnitMap["day"]
			dateConstantHelper.TimeUnitMap["month"] = 30 * dateConstantHelper.TimeUnitMap["day"]
			dateConstantHelper.TimeUnitMap["year"] = 365 * dateConstantHelper.TimeUnitMap["day"]

			// format map
			dateConstantHelper.FormatMap["yyyy"] = stdLongYear
			dateConstantHelper.FormatMap["yy"] = stdYear
			dateConstantHelper.FormatMap["MM"] = stdZeroMonth
			dateConstantHelper.FormatMap["M"] = stdNumMonth
			dateConstantHelper.FormatMap["dd"] = stdZeroDay
			dateConstantHelper.FormatMap["d"] = stdDay
			dateConstantHelper.FormatMap["HH"] = stdHour
			dateConstantHelper.FormatMap["m"] = stdMinute
			dateConstantHelper.FormatMap["mm"] = stdZeroMinute
			dateConstantHelper.FormatMap["s"] = stdSecond
			dateConstantHelper.FormatMap["ss"] = stdZeroSecond
		})
	}
	return dateConstantHelper
}

var dateFormatUtilOnce sync.Once
var dateFormatUtil *DateFormatUtil

// DateFormatUtil 日期格式工具
type DateFormatUtil struct {
}

func (util *DateFormatUtil) replace(input string, regx string) string {
	c := regexp.MustCompile(regx)
	extract := c.FindString(input)

	helper := GetDateConstantHelper()
	val, ok := helper.FormatMap[extract]
	if !ok {
		return input
	}
	return strings.Replace(input, extract, val, -1)
}

// ConvertDateFormat 将通用格式转换成go的日期格式
func (util *DateFormatUtil) ConvertDateFormat(str string) string {
	// 替换年份
	str = util.replace(str, "y+")
	// 替换月份
	str = util.replace(str, "M+")
	// 替换日
	str = util.replace(str, "d+")
	// 替换小时
	str = util.replace(str, "H+")
	// 替换分钟
	str = util.replace(str, "m+")
	// 替换秒
	str = util.replace(str, "s+")

	return str
}

// FormatDate 通用格式化日期
func (util *DateFormatUtil) FormatDate(ts int64, format string) string {
	// 转换时间格式
	dateFormatStr := util.ConvertDateFormat(format)

	// 返回格式化时间
	return time.Unix(ts/1000, ts%1000*1e6).Format(dateFormatStr)
}

// GetDateFormatUtil 获取日期格式工具
func GetDateFormatUtil() *DateFormatUtil {
	if dateFormatUtil == nil {
		dateFormatUtilOnce.Do(func() {
			dateFormatUtil = &DateFormatUtil{}
		})
	}
	return dateFormatUtil
}

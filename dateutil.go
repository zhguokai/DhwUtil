package goutil

import (
	"time"
	"errors"
)

const (
	YYYY_MM_DD__HI_MM_SS = "2006-01-02 15:04:05"
	YYYY_MM_DD = "2006-01-02"
	YYYY_MM = "2006-01"
	YYYY = "2006"
	YYYYMMDDHIMMSS = "20060102150405"
	YYYYMMDD = "20060102"
	YYYYMM = "200601"
)

// GetCurrentDate 获取当前月份,
// 返回格式为:YYYY-MM
// 示例:2016-06
func GetCurrentMonth() string {
	return time.Now().Format(YYYY_MM)
}

// GetCurrentDate 获取当前日期,
// 返回格式为:YYYY-MM-DD
// 示例:2016-06-05
func GetCurrentDate() string {
	return time.Now().Format(YYYY_MM_DD)
}

// GetCurrentTime 获取当前时间,
// 返回格式为:YYYY-MM-DD HI:MI:SS,
// 示例:2016-06-05 12:22:33
func GetCurrentTime() string {
	return time.Now().Format(YYYY_MM_DD__HI_MM_SS)
}

// GetPreDayTime 获取当前时间之前的N天的时间,
// 参数为int64位整型,
// 当参数小于0时,代表获取当前时间之后的N天时间
// 返回格式为:YYYY-MM-DD HI:MI:SS,
// 示例:2016-06-05 12:22:33
func GetPreDayTime(day int64) string {
	sec := time.Now().Unix() - day * 24 * 60 * 60
	timeStr, _ := UnixTimeToStr(YYYY_MM_DD__HI_MM_SS, sec)
	return timeStr
}

// GetPreHourTime 获取当前时间之前的N小时的时间,
// 参数为int64位整型,
// 当参数小于0时,获取当前时间之后的N小时的时间
// 返回格式为:YYYY-MM-DD HI:MI:SS,
// 示例:2016-06-05 12:22:33
func GetPreHourTime(hour int64) string {
	sec := time.Now().Unix() - hour * 60 * 60
	timeStr, _ := UnixTimeToStr(YYYY_MM_DD__HI_MM_SS, sec)
	return timeStr
}

// GetPreMinuteTime 获取当前时间之前的N分钟时间,
// 参数为int64位整型,
// 当参数小于0时,代表获取当前时间之后的N分钟时间
// 返回格式为:YYYY-MM-DD HI:MI:SS,
// 示例:2016-06-05 12:22:33
func GetPreMinuteTime(minute int64) string {
	sec := time.Now().Unix() - minute * 60
	timeStr, _ := UnixTimeToStr(YYYY_MM_DD__HI_MM_SS, sec)
	return timeStr
}

// GetPreSecondTime 获取当前时间之前的N秒钟时间,
// 参数为int64位整型,
// 当参数小于0时,代表获取当前时间之后的N秒钟时间
// 返回格式为:YYYY-MM-DD HI:MI:SS,
// 示例:2016-06-05 12:22:33
func GetPreSecondTime(second int64) string {
	sec := time.Now().Unix() - second
	timeStr, _ := UnixTimeToStr(YYYY_MM_DD__HI_MM_SS, sec)
	return timeStr
}
// UnixTimeToStr 将Unix时间转换为字符串格式(YYYY-MM-DD HI:mm:ss),
// 参数为int64位整型,
// 返回格式为:YYYY-MM-DD HI:MI:SS,
// 示例:2016-06-05 12:22:33
func UnixTimeToStr(format string, t1 int64) (timeStr string, err error) {
	if format == YYYY || format == YYYY_MM_DD__HI_MM_SS || format == YYYY_MM_DD ||
	format == YYYY_MM || format == YYYYMM || format == YYYYMMDDHIMMSS || format == YYYYMMDD {
		tm := time.Unix(t1, 0)
		return tm.Format(format), nil
	}
	return YYYY_MM_DD__HI_MM_SS, errors.New("暂时不支持的日期格式")

}

// StrToUnixTime 将Unix时间转换为字符串格式(format),
func StrToUnixTime(format, t1 string) (unixTime int64, err error) {
	if format == YYYY || format == YYYY_MM_DD__HI_MM_SS || format == YYYY_MM_DD ||
	format == YYYY_MM || format == YYYYMM || format == YYYYMMDDHIMMSS || format == YYYYMMDD {
		tm, err := time.Parse(format, t1)
		if err != nil {
			return 0, err
		}
		return tm.Unix(), nil
	}
	return 0, errors.New("暂时不支持的日期格式")

}

package goutil

import (
	"time"
	"errors"
	"log"
	"strconv"
	"strings"
)

//日期类型


const (
	YYYY_MM_DD__HI_MM_SS = "2006-01-02 15:04:05"
	YYYY_MM_DD = "2006-01-02"
	YYYY_MM = "2006-01"
	YYYY = "2006"
	YYYYMMDDHIMMSS = "20060102150405"
	YYYYMMDD = "20060102"
	YYYYMM = "200601"
	YYYYMMDDHIMMSSsss = "20060102150405.9999"
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
//获取当前时间
func GetFormatTime(format string) (string) {
	return strings.Replace(time.Now().Format(format), ".", "", -1)
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
		tm, err := time.ParseInLocation(format, t1, time.Local)
		if err != nil {
			return 0, err
		}
		return tm.Unix(), nil
	}
	return 0, errors.New("暂时不支持的日期格式")

}

//获取纳秒级时间差
func GetCurrentNano() int64 {
	return time.Now().UnixNano()
}

//获取指定时间前几天时间
func GetBeforeDayTime(format, nowTime string, day int64) string {

	uT, err := StrToUnixTime(format, nowTime)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	preUt := uT - day * 24 * 60 * 60
	sT, err := UnixTimeToStr(format, preUt)

	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return sT
}

func GetBeforeSecondTime(format, nowTime string, second int64) string {

	uT, err := StrToUnixTime(format, nowTime)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	preUt := uT - second
	sT, err := UnixTimeToStr(format, preUt)

	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return sT
}

//获取当月的开始时间,例:2016-06-08 00:00:00
func GetFirstDayTimeInMonth() string {
	year, mon, _ := time.Now().Date()
	forMon := ""
	if mon < 10 {
		forMon = "0" + strconv.Itoa(int(mon))
	} else {
		forMon = strconv.Itoa(int(mon))
	}
	strTime := strconv.Itoa(year) + "-" + forMon + "-01 00:00:00"
	return strTime
}

//取当月的结束时间,如: 2016-06-30 23:59:59
func GetLastDayTimeInMonth() string {
	firstTime := ""
	year, mon, _ := time.Now().Date()
	nextM := int(mon)
	if nextM < 12 {
		nextM = nextM + 1
	} else {
		nextM = 1
		year = year + 1
	}
	forMon := ""
	if nextM < 10 {
		forMon = "0" + strconv.Itoa(nextM)
	} else {
		forMon = strconv.Itoa(nextM)
	}
	firstTime = strconv.Itoa(year) + "-" + forMon + "-01 00:00:00"
	return GetBeforeSecondTime(YYYY_MM_DD__HI_MM_SS, firstTime, 1)
}

package utils

import (
	"github.com/golang-module/carbon/v2"
	"strings"
	"time"
)

const (
	FormatYYYYMMDDHHMMSS string = "2006-01-02 15:04:05"
)

var LocShanghai *time.Location

// GetUnixStamp2Time 将10位时间戳转换为time格式
func GetUnixStamp2Time(timestamp int64) time.Time {
	return carbon.CreateFromTimestamp(timestamp).Carbon2Time()
}

// GetDayOfYear 获取本年第几天
func GetDayOfYear(times time.Time) int {
	dayOfYear := carbon.Time2Carbon(times).DayOfYear()
	return dayOfYear
}

// GetWeekOfYear 获取本年第几周
func GetWeekOfYear(times time.Time) int {
	return carbon.Time2Carbon(times).WeekOfYear()
}

// GetMonthOfYear 获取本年第几月
func GetMonthOfYear(times time.Time) int {
	return carbon.Time2Carbon(times).MonthOfYear()
}

// GetYear 获取年份
func GetYear(times time.Time) int {
	return carbon.Time2Carbon(times).Year()
}

// GetYearMonthDay 返回所有年月日
func GetYearMonthDay(now time.Time) (year, month, week, day int) {
	// 获取获取本年第几天
	day = GetDayOfYear(now)
	// 获取本年第几周
	week = GetWeekOfYear(now)
	// 获取本年第几天
	month = GetMonthOfYear(now)
	// 获取年份
	year = GetYear(now)
	return year, month, week, day
}

// GetStartOfToday 获取今天的开始时间戳
func GetStartOfToday() int64 {
	return carbon.Now().StartOfDay().Timestamp()
}

// GetEndOfToday 获取今天的结束时间戳
func GetEndOfToday() int64 {
	return carbon.Now().EndOfDay().Timestamp()
}

// GetTimeStr2TimeStamp 根据传入的时间字符串转换为时间戳
func GetTimeStr2TimeStamp(timeStr string) int64 {
	return carbon.Parse(timeStr).Timestamp()
}

// GetTimeBeforeThirtyDays 获取最近30天开始时间
func GetTimeBeforeThirtyDays(days int) string {
	return carbon.Now().SubDays(days).StartOfDay().ToDateTimeString()
}

func GetTimeStr(time2 time.Time) string {
	return carbon.Time2Carbon(time2).ToDateTimeString()
}

func GetTimeStrForTimestamp(timeStr string) int64 {
	return carbon.Parse(timeStr).Timestamp()
}

// GetLunarZodiac 获取农历生肖
func GetLunarZodiac(time time.Time) string {
	return carbon.Time2Carbon(time).Lunar().Animal()
}

func GetWeekStartTime() int64 {
	return carbon.Now().StartOfWeek().Timestamp()
}
func GetMonthStartTime() int64 {
	return carbon.Now().StartOfMonth().Timestamp()
}

// Format 跟 PHP 中 date 类似的使用方式，如果 ts 没传递，则使用当前时间
func Format(format string, ts ...time.Time) string {
	patterns := []string{
		// 年
		"Y", "2006", // 4 位数字完整表示的年份
		"y", "06", // 2 位数字表示的年份

		// 月
		"m", "01", // 数字表示的月份，有前导零
		"n", "1", // 数字表示的月份，没有前导零
		"M", "Jan", // 三个字母缩写表示的月份
		"F", "January", // 月份，完整的文本格式，例如 January 或者 March

		// 日
		"d", "02", // 月份中的第几天，有前导零的 2 位数字
		"j", "2", // 月份中的第几天，没有前导零

		"D", "Mon", // 星期几，文本表示，3 个字母
		"l", "Monday", // 星期几，完整的文本格式;L的小写字母

		// 时间
		"g", "3", // 小时，12 小时格式，没有前导零
		"G", "15", // 小时，24 小时格式，没有前导零
		"h", "03", // 小时，12 小时格式，有前导零
		"H", "15", // 小时，24 小时格式，有前导零

		"a", "pm", // 小写的上午和下午值
		"A", "PM", // 小写的上午和下午值

		"i", "04", // 有前导零的分钟数
		"s", "05", // 秒数，有前导零
	}
	replacer := strings.NewReplacer(patterns...)
	format = replacer.Replace(format)

	t := time.Now()
	if len(ts) > 0 {
		t = ts[0]
	}
	return t.Format(format)
}

func ParseBeijingTime(sTime string) (t time.Time, err error) {
	return time.ParseInLocation(FormatYYYYMMDDHHMMSS, sTime, LocShanghai)
}

// IsToday 是否今天
func IsToday(dateStamp int64) bool {
	return carbon.CreateFromTimestamp(dateStamp).IsToday()
}

func GetDayStr() string {
	return carbon.Now().ToDateString()
}

// 获取时间范围内的小时列表
func GetHourBetweenDates(sdate, edate string) []string {
	list := []string{}
	timeFormatTpl := "2006-01-02 15:04:05"
	if len(timeFormatTpl) != len(sdate) {
		timeFormatTpl = timeFormatTpl[0:len(sdate)]
	}
	date, err := time.Parse(timeFormatTpl, sdate)
	if err != nil {
		// 时间解析，异常
		return list
	}
	date2, err := time.Parse(timeFormatTpl, edate)
	if err != nil {
		// 时间解析，异常
		return list
	}
	if date2.Before(date) {
		// 如果结束时间小于开始时间，异常
		return list
	}
	// 输出日期格式固定
	timeFormatTpl = "2006-01-02 15:04:05"
	list = append(list, date.Format(timeFormatTpl))
	for {
		date = date.Add(time.Hour * time.Duration(1))
		dateStr := date.Format(timeFormatTpl)
		// println(dateStr)
		if date.After(date2) {
			break
		}
		list = append(list, dateStr)
	}
	return list
}

// 获取时间范围内的天列表
func GetDateBetweenDates(sdate, edate string) []string {
	list := []string{}
	timeFormatTpl := "2006-01-02"
	if len(timeFormatTpl) != len(sdate) {
		timeFormatTpl = timeFormatTpl[0:len(sdate)]
	}
	date, err := time.Parse(timeFormatTpl, sdate)
	if err != nil {
		// 时间解析，异常
		return list
	}
	date2, err := time.Parse(timeFormatTpl, edate)
	if err != nil {
		// 时间解析，异常
		return list
	}
	if date2.Before(date) {
		// 如果结束时间小于开始时间，异常
		return list
	}
	// 输出日期格式固定
	timeFormatTpl = "2006-01-02"
	list = append(list, date.Format(timeFormatTpl))
	for {
		date = date.AddDate(0, 0, 1)
		dateStr := date.Format(timeFormatTpl)
		// println(dateStr)
		if date.After(date2) {
			break
		}
		list = append(list, dateStr)
	}
	return list
}

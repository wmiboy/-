package xTime

import (
	"xutil/xStr"
)

import "time"

//获取时间戳可根据参数取当天或指定天数内的时间戳
func GetUnix(days int) string {
	t := time.Now()
	return xStr.Int642str(t.AddDate(0, 0, days).Unix())
}

//对比时间戳，传入的时间戳比当前时间小为false
func CheckUnix(t1 string) bool {
	if xStr.Str2int64(t1) < time.Now().Unix() {
		return false
	}
	return true
}
func Str2Unix(t string) string {
	t2, _ := time.Parse("01/02/2006", t)
	return xStr.Int642str(t2.Unix())
}

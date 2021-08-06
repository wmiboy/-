package util

import (
	"math/rand"
	"time"
	"xutil/xEncrypt"
	"xutil/xStr"
	"xutil/xTime"
)

func GetCdk() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	rad := xStr.Int2str(rand.Intn(100000))
	tmp := xEncrypt.Md5(xTime.GetUnix(0) + rad)
	return tmp[0:16]
}

func GetTime() int {
	return xStr.Str2int(xTime.GetUnix(0))
}

func Page(Limit, Page int) (limit, offset int) {
	if Limit > 0 {
		limit = Limit
	} else {
		limit = 10
	}
	if Page > 0 {
		offset = (Page - 1) * limit
	} else {
		offset = -1
	}
	return limit, offset
}

// Sort 排序
// 默认 created_at desc
func Sort(Sort string) (sort string) {
	if Sort != "" {
		sort = Sort
	} else {
		sort = "create_at desc"
	}
	return sort
}

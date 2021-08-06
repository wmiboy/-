package xStr

import "strconv"

//类型转换
//64位整数转文本
func Int642str(i int64) string {
	return strconv.FormatInt(i, 10)
}

//文本转64位整数
func Str2int64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		i = 0
	}
	return i
}

//整数转文本
func Int2str(i int) string {
	return strconv.Itoa(i)
}

//文本转整数
func Str2int(s string) int {
	int, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return int
}

//float64转文本
func Float2Str(s float64) string {
	return strconv.FormatFloat(s, 'f', 2, 64)
}

//文本转float64
func Str2Float(s string) float64 {
	i, err := strconv.ParseFloat(s, 32/64)
	if err != nil {
		i = 0
	}
	return i
}

package xEncrypt

import (
	"crypto/md5"
	"fmt"
)

func Md5(s string) string {
	has := md5.Sum([]byte(s))
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str1
}
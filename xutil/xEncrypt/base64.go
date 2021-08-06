package xEncrypt

import (
	"encoding/base64"
)

func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
func Base64Decode(s string) string {
	sDec, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return ""
	}
	return string(sDec)
}

package util

import (
	"encoding/base64"
	glog "web/log"
)

func Decrypt_key(pKey string) (error, []byte) {
	key, err := base64.StdEncoding.DecodeString(pKey)
	if err != nil {
		return err, nil
	}
	key, err = RsaDecrypt(key)

	glog.GetLog().Debug("RSA解密结果----", "密文:", pKey, "原文:", string(key), "ERR:", err)
	return err, key
}
func Decrypt_data(pData string, pKey []byte) (error, []byte) {
	byte, err := base64.StdEncoding.DecodeString(pData)
	if err != nil {
		return err, nil
	}
	err, byte = AesDecryptCBC(byte, pKey)

	glog.GetLog().Debug("AES解密结果----", "密文:", pData, "原文:", string(byte), "ERR:", err)
	return err, byte
}

func Encrypt(pKey string, pData []byte) (error, string) {
	key := []byte(pKey)
	err, byte := AesEncryptCBC(pData, key)
	if err != nil {
		return err, ""
	}
	str := base64.StdEncoding.EncodeToString(byte)
	glog.GetLog().Debug("AES加密----原文:", pData, "秘钥:", pKey, "密文:", str, "ERR", err)
	return err, str
}

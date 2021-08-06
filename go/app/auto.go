package app

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"web/util"
	"xutil/xStr"
)

func sign(pAUTH *AUTH) string {
	str := xStr.Int2str(pAUTH.Time+pAUTH.Uid+pAUTH.Aid+pAUTH.Time_Login) + "WOrjME0LGpzZ4xi2" + pAUTH.Dev + pAUTH.Power
	has := md5.Sum([]byte(str))
	sign := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return sign
}

func SignToken(pAUTH *AUTH) (error, string) {
	pAUTH.Time = util.GetTime()
	pAUTH.Sign = sign(pAUTH)
	b, err := json.Marshal(pAUTH)
	if err != nil {
		return errors.New("Token发放失败"), ""
	}
	token := base64.StdEncoding.EncodeToString(b)
	return nil, token
}

func GetToken(pToken string) (error, *AUTH) {
	var a AUTH
	if len(pToken) == 0 {
		return errors.New("请先登录"), &a
	}
	byte, err := base64.StdEncoding.DecodeString(pToken)
	if err != nil {
		return errors.New("Token不合法-1000"), &a
	}
	err = json.Unmarshal(byte, &a)
	if err != nil {
		return errors.New("Token不合法-2000"), &a
	}
	sign := sign(&a)
	if sign != a.Sign {
		return errors.New("Token不合法-3000"), &a
	}
	return nil, &a
}

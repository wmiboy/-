package cmw

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
	"web/app"
	"web/conf"
	glog "web/log"
	"web/util"
	"xutil/xEncrypt"
	"xutil/xStr"
)

func CMW_Encrypt() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := Encrypt(c)
		if err != nil {
			c.Abort()
			app.Response(c, 800, err.Error())
			return
		}
		c.Next()
	}
}
func Encrypt(c *gin.Context) error {
	t := c.PostForm("t")
	s := c.PostForm("s")
	sign := xEncrypt.Md5(t + conf.Conf_Get("encrypt.salt").(string))
	if s != sign {
		return errors.New("请求异常-1006")
	}
	if util.GetTime()-xStr.Str2int(t) > 120 || xStr.Str2int(t)-util.GetTime() > 120 {
		return errors.New("请求异常-1007")
	}
	err, key := util.Decrypt_key(c.PostForm("k"))
	if err != nil {
		return errors.New("请求异常-1001")
	}
	err, buf := util.Decrypt_data(c.PostForm("d"), key)
	if err != nil {
		return errors.New("请求异常-1002")
	}
	c.Request.PostForm.Add("key", string(key))
	err = jieXi(string(buf), c)
	return err
}
func jieXi(pData string, c *gin.Context) error {
	arr := strings.Split(pData, "&")
	if len(arr) == 0 {
		return errors.New("请求异常-1005")
	}
	for _, v := range arr {
		tmp := strings.Split(v, `=`)
		if len(tmp) != 2 {
			continue
		}
		c.Request.PostForm.Add(tmp[0], tmp[1])
		glog.GetLog().Debug(tmp[0], ":", tmp[1])
	}
	return nil
}

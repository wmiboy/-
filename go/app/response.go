package app

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"web/util"
)

var isKuaYu = true

func kua_Yu(c *gin.Context) {
	if !isKuaYu {
		return
	}
	c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
}

func Response(c *gin.Context, pCode int, pMsg string) {
	kua_Yu(c)
	body := &response{Code: pCode, Msg: pMsg}
	c.JSON(http.StatusOK, &body)
}

func Response_data(c *gin.Context, pCode int, pMsg string, pData interface{}) {
	kua_Yu(c)
	body := &response_data{Code: pCode, Msg: pMsg, Data: pData}
	c.JSON(http.StatusOK, &body)
}
func Response_Encrypt(c *gin.Context, pCode int, pMsg string, pData interface{}) {
	kua_Yu(c)
	byte, err := json.Marshal(&pData)
	if err != nil {
		Response(c, 400, "请求异常-1003")
		return
	}
	err, str := util.Encrypt(c.PostForm("key"), byte)
	if err != nil {
		Response(c, 400, "请求异常-1004")
		return
	}
	body := &response_data{Code: 300, Msg: pMsg, Data: str}
	c.JSON(http.StatusOK, &body)
}
func Response_list(c *gin.Context, pCode int, pMsg string, pData interface{}) {
	kua_Yu(c)
	body := &response_list{Code: pCode, Msg: pMsg, Data: pData, Count: 100000}
	c.JSON(http.StatusOK, &body)
}

package controller

import (
	"github.com/gin-gonic/gin"
	"web/app"
	"web/app/server"
)

type c_power struct {
	srv *server.Power
}
// @Summary 充值
// @Description 详细描述
// @Tags 用户操作
// @Produce json
// @Param user formData string true "账号"
// @Param cdk formData string true "卡密"
// @Param aid formData string true "应用ID"
// @Success 200 {string} json "{"code":200,"msg":"充值成功"}"
// @Router /power/charge [post]
func (c2 *c_power) Charge(c *gin.Context) {
	err := c2.srv.Charge(c)
	if err != nil {
		app.Response(c, 400, err.Error())
	} else {
		app.Response(c, 200, "充值成功")
	}
}
// @Summary 获取权限列表
// @Description 详细描述
// @Tags 权限管理
// @Produce json
// @Param token query string true "登录返回的token"
// @Param page query string true "页数"
// @Param limit query string true "获取的数量"
// @Param aid query string true "应用ID"
// @Param state query string true "状态1=可用 2=不可用"
// @Param uid query string false "用户ID"
// @Success 200 {string} json "{"code":200,"msg":"查询成功","count":1,"data":[{"uid":1001,"aid":1002,"kid":0,"day":1625191303,"point":200,"ip":"","code":"","state":1,"time":1625018503}]}"
// @Router /power/getAll [get]
func (c2 *c_power) GetAll(c *gin.Context) {
	err, list := c2.srv.GetAll(c)
	if err != nil {
		app.Response(c, 400, err.Error())
	} else {
		app.Response_list(c, 200, "查询成功", list)
	}
}

// @Summary 编辑权限信息
// @Description
// @Tags 权限管理
// @Produce json
// @Param token query string true "登录返回的token"
// @Param uid formData string true "用户ID"
// @Param aid formData string true "应用ID"
// @Param state formData string false "1=可用 2=不可用"
// @Param day formData string false "到期时间的十位时间戳"
// @Param point formData string false "剩余点数"
// @Param ip formData string false "绑定IP"
// @Param code formData string false "绑定机器"
// @Success 200 {string} json "{"code":200,"msg":"修改成功"}"
// @Router /power/edit [post]
func (c2 *c_power) Edit(c *gin.Context) {
	err := c2.srv.Edit(c)
	if err != nil {
		app.Response(c, 400, err.Error())
	} else {
		app.Response(c, 200, "修改成功")
	}
}
// @Summary 心跳[数据已加密]
// @Description 数据已加密
// @Tags 验证相关
// @Produce json
// @Param token query string true "登录返回的token"
// @Param dev formData string true "机器码"
// @Success 200 {string} json "{"code":200,"msg":"效验成功","data":{"a":{"aid":3,"name":"JD注册","mold":2,"login":3,"bin":1,"msg":"关键参数222","state":1,"time":1625069421},"p":{"pid":3,"uid":1,"aid":3,"kid":0,"day":1625183400,"point":90,"ip":"192.168.1.7","code":"ca5df0fdae677510f937c256a7f5","state":1,"time":1625084599}}}"
// @Router /power/check [post]
func (c2 *c_power) Check(c *gin.Context) {
	err, apply, power := c2.srv.Check(c)
	if err != nil {
		app.Response(c, 400, err.Error())
	} else {
		app.Response_Encrypt(c, 200, "效验成功", gin.H{"a": apply, "p": power})
	}

}
// @Summary 用户登录[数据已加密]
// @Description 数据已加密
// @Tags 验证相关
// @Produce json
// @Param user formData string true "账号"
// @Param pwd formData string true "密码"
// @Param aid formData string true "应用ID"
// @Param dev formData string true "机器码"
// @Success 200 {string} json "{"code":200,"msg":"登陆成功","data":{"token":"111..."}}"
// @Router /power/login [post]
func (c2 *c_power) Login_User(c *gin.Context) {
	err, token := c2.srv.Login_User(c)
	if err != nil {
		app.Response(c, 400, err.Error())
	} else {
		app.Response_Encrypt(c, 200, "登陆成功", gin.H{"token": token})
	}
}

func (c2 *c_power) Login_Cdk(c *gin.Context) {

}
// @Summary 扣点[数据已加密]
// @Description 数据已加密
// @Tags 验证相关
// @Produce json
// @Param token query string true "登录返回的token"
// @Param point formData string true "扣除点数不可为负"
// @Success 200 {string} json "{"code":200,"msg":"扣除成功"}"
// @Router /power/sub [post]
func (c2 *c_power) Sub(c *gin.Context) {
	err := c2.srv.Sub(c)
	if err != nil {
		app.Response(c, 400, err.Error())
	} else {
		app.Response(c, 200, "扣除成功")
	}
}

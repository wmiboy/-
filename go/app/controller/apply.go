package controller

import (
	"github.com/gin-gonic/gin"
	"web/app"
	"web/app/dao"
	"web/app/enum"
	"web/app/mode"
	"xutil/xStr"
)

type c_apply struct {
	dao dao.APPLY
}
// @Summary 添加应用
// @Description 详细描述
// @Tags 应用管理
// @Produce json
// @Param token query string true "登录返回的token"
// @Param name formData string true "应用名称"
// @Param mold formData string true "收费模式 1=扣时 2=扣点"
// @Param bin formData string true "绑定模式 1=不绑定 2=绑定机器码 3=绑定IP"
// @Param login formData string true "登录模式 1=免费 2=单点登录 3=多点登录"
// @Param msg formData string false "远程数据"
// @Success 200 {string} json "{"code":200,"msg":"添加成功"}"
// @Router /apply/add [post]
func (c2 c_apply) Add(c *gin.Context) {
	var a mode.Apply
	err := c.ShouldBind(&a)
	if err != nil || len(a.Name) == 0 {
		app.Response(c, 400, "参数错误")
		return
	}
	if a.Mold <= 0 && a.Mold > 2 {
		a.Mold = enum.APPLY_MOLD_KouShi
	}
	if a.Login <= 0 && a.Login > 3 {
		a.Login = enum.APPLY_Login_MianFei
	}
	if a.Bin <= 0 && a.Bin > 3 {
		a.Bin = enum.APPLY_Bin_NoBin
	}

	err = c2.dao.ADD(&a)
	if err != nil {
		app.Response(c, 400, "添加失败")
		return
	}
	app.Response(c, 200, "添加成功")
}

// @Summary 获取应用信息[数据已加密]
// @Description 数据已加密处理
// @Tags 应用管理
// @Produce json
// @Param aid formData string true "应用ID"
// @Success 200 {string} json "{"code":200,"msg":"查询成功","data":{"aid":1002,"name":"QQ注册1","mold":2,"login":1,"bin":3,"msg":"关键参数2333","state":2,"time":1625000360}}"
// @Router /apply/get [post]
func (c2 c_apply) Get(c *gin.Context) {
	aid := c.PostForm("aid")
	err, a := c2.dao.Get(xStr.Str2int(aid))
	if err != nil || a.Aid == 0 {
		app.Response(c, 400, "软件不存在")
		return
	}
	if a.Login != enum.APPLY_Login_MianFei {
		app.Response(c, 400, "收费软件不可直接获取信息")
		return
	}
	app.Response_Encrypt(c, 200, "查询成功", a)
}
// @Summary 获取应用列表
// @Description 详细描述
// @Tags 应用管理
// @Produce json
// @Param token query string true "登录返回的token"
// @Param page query string true "页数"
// @Param limit query string true "获取的数量"
// @Param aid query string true "应用ID"
// @Param state query string true "应用状态1=可用 2=不可用"
// @Param name query string false "应用名称"
// @Success 200 {string} json "{"code":200,"msg":"查询成功","count":2,"data":[{"aid":1001,"name":"JD注册","mold":1,"login":3,"bin":2,"msg":"关键参数","state":1,"time":1625000292}]}"
// @Router /apply/getAll [get]
func (c2 c_apply) GetAll(c *gin.Context) {
	var a mode.Apply
	err := c.ShouldBind(&a)
	if err != nil {
		app.Response(c, 400, "参数错误")
		return
	}

	err, list := c2.dao.GetAll(&a)
	if err != nil {
		app.Response(c, 400, "查询失败")
		return
	}
	app.Response_list(c, 200, "查询成功", list)
}

// @Summary 编辑应用信息
// @Description
// @Tags 应用管理
// @Produce json
// @Param token query string true "登录返回的token"
// @Param aid formData string true "应用ID"
// @Param name formData string false "应用名称"
// @Param mold formData string false "收费模式 1=扣时 2=扣点"
// @Param bin formData string false "绑定模式 1=不绑定 2=绑定机器码 3=绑定IP"
// @Param login formData string false "登录模式 1=免费 2=单点登录 3=多点登录"
// @Param msg formData string false "远程数据"
// @Param state formData string false "1=可用 2=不可用"
// @Success 200 {string} json "{"code":200,"msg":"修改成功"}"
// @Router /apply/edit [post]
func (c2 c_apply) Edit(c *gin.Context) {
	var a mode.Apply
	err := c.ShouldBind(&a)
	if err != nil || a.Aid == 0 {
		app.Response(c, 400, "软件不存在")
		return
	}

	err = c2.dao.Edit(&a)
	if err != nil {
		app.Response(c, 400, "修改失败")
		return
	}
	app.Response(c, 200, "修改成功")
}

type i_apply interface {
	Add(c *gin.Context)
	Get(c *gin.Context)
	GetAll(c *gin.Context)
	Edit(c *gin.Context)
}

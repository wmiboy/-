package controller

import (
	"github.com/gin-gonic/gin"
	"web/app"
	"web/app/dao"
	"web/app/mode"
	"xutil/xStr"
)

type c_key struct {
	dao dao.CDK
}
// @Summary 生成卡密
// @Description 详细描述
// @Tags 卡密管理
// @Produce json
// @Param token query string true "登录返回的token"
// @Param num formData string true "生成数量"
// @Param mold formData string true "收费模式 1=扣时 2=扣点"
// @Param point formData string true "点数"
// @Param day formData string true "小时数"
// @Param aid formData string true "应用ID"
// @Success 200 {string} json "{"code":200,"msg":"生成成功","data":[{"kid":2001,"cdk":"acff954d-9160-4cc3-96c3-e9284543a68b","mold":1,"point":100,"day":24,"uid":0,"aid":1002,"cread_time":1625006496,"use_time":0,"state":1}]}"
// @Router /key/cread [post]
func (c2 c_key) cread(c *gin.Context) {
	var k mode.CDK
	err := c.ShouldBind(&k)
	if err != nil {
		app.Response(c, 400, "参数异常")
		return
	}
	num := xStr.Str2int(c.PostForm("num"))
	if num == 0 {
		num = 100
	} else if num > 1000 {
		num = 1000
	}

	err, list := c2.dao.Cread(num, &k)
	if err != nil {
		app.Response(c, 400, "生成失败")
		return
	}
	app.Response_data(c, 200, "生成成功", list)
}
// @Summary 获取卡密列表
// @Description 详细描述
// @Tags 卡密管理
// @Produce json
// @Param token query string true "登录返回的token"
// @Param page query string true "页数"
// @Param limit query string true "获取的数量"
// @Param aid query string false "应用ID"
// @Param state query string false "应用状态1=可用 2=不可用"
// @Param cdk query string false "卡密"
// @Success 200 {string} json "{"code":200,"msg":"查询成功","count":1004,"data":[{"kid":1005,"cdk":"3c262067-384b-4ba5-80ab-14ca31cbadf5","mold":1,"point":100,"day":24,"uid":0,"aid":1002,"cread_time":1625006471,"use_time":0,"state":1}]}"
// @Router /key/getAll [get]
func (c2 c_key) getAll(c *gin.Context) {
	var k mode.CDK
	err := c.ShouldBind(&k)
	if err != nil {
		app.Response(c, 400, "参数异常")
		return
	}
	err, list := c2.dao.GetAll(&k)
	if err != nil {
		app.Response(c, 400, "查询失败")
		return
	}
	app.Response_list(c, 200, "查询成功", list)
}

// @Summary 编辑卡密信息
// @Description
// @Tags 卡密管理
// @Produce json
// @Param token query string true "登录返回的token"
// @Param kid formData string true "卡密ID"
// @Param state formData string true "1=可用 2=不可用"
// @Success 200 {string} json "{"code":200,"msg":"修改成功"}}"
// @Router /key/edit [post]
func (c2 c_key) edit(c *gin.Context) {
	var k mode.CDK
	err := c.ShouldBind(&k)
	if err != nil {
		app.Response(c, 400, "参数异常")
		return
	}

	err = c2.dao.Edit(&k)
	if err != nil {
		app.Response(c, 400, "修改失败")
		return
	}
	app.Response(c, 200, "修改成功")
}

type i_key interface {
	cread(c *gin.Context)
	get(c *gin.Context)
	getAll(c *gin.Context)
	edit(c *gin.Context)
}

package controller

import (
	"github.com/gin-gonic/gin"
	"web/app"
	"web/app/cmw"
	"web/app/dao"
	"web/app/mode"
)

type c_user struct {
	dao dao.User
}
// @Summary 获取个人信息
// @Description 详细描述
// @Tags 用户操作
// @Produce json
// @Param token query string true "登录返回的token"
// @Success 200 {string} json "{"code":200,"msg":"更新成功"}
// @Router /user/getInfo [get]
func (c2 *c_user) getInfo(c *gin.Context) {
	f_uid := c.GetInt("uid")
	err, user := c2.dao.GetInfo(f_uid)
	if err != nil {
		app.Response(c, 400, "获取失败")
		return
	}
	app.Response_data(c, 200, "获取成功", user.Msg)
}

// @Summary 更新个人信息
// @Description 详细描述
// @Tags 用户操作
// @Produce json
// @Param token query string true "登录返回的token"
// @Param msg formData string true "更新内容"
// @Success 200 {string} json "{"code":200,"msg":"更新成功"}
// @Router /user/upInfo [post]
func (c2 *c_user) UpInfo(c *gin.Context) {
	f_msg := c.PostForm("msg")
	f_uid := c.GetInt("uid")
	err := c2.dao.UpInfo(f_uid, f_msg)
	if err != nil {
		app.Response(c, 400, "更新失败")
		return
	}
	app.Response(c, 200, "更新成功")
}

// @Summary 登录
// @Description 详细描述
// @Tags 用户操作
// @Produce json
// @Param user formData string true "账号"
// @Param pwd formData string true "密码"
// @Success 200 {string} json "{"code":200,"msg":"登录成功","data":"eyJ1aWQiOjEwMDEsImFpZCI6MCwidGltZSI6MTYyNDk4MjQ5OSwicG93ZXIiOjIsInRpbWVfdXBkYXRhIjowLCJ0aW1lX0xvZ2luIjowLCJzaWduIjoiNWQ3OGYyYTE4OGY1Y2NhNjQ2NDdhYzI4N2MyMmVlMzIifQ=="}"
// @Router /user/login [post]
func (c2 *c_user) login(c *gin.Context) {
	f_user := c.PostForm("user")
	f_pwd := c.PostForm("pwd")
	if len(f_user) == 0 || len(f_pwd) == 0 {
		app.Response(c, 400, "用户名或密码错误")
		return
	}

	err, user := c2.dao.Login(f_user, f_pwd)
	if err != nil {
		app.Response(c, 400, "用户名或密码错误")
		return
	}
	auth := cmw.AUTH_GET(user.Type)
	err, token := app.SignToken(&app.AUTH{
		Uid:         user.Uid,
		Power:       auth,
		Time_upData: user.Uptime,
	})
	if err != nil {
		app.Response(c, 400, "令牌发放失败,请重新登录")
		return
	}
	app.Response_data(c, 200, "登录成功", gin.H{"token": token, "auth": auth, "name": user.User})
}

// @Summary 注册
// @Description 详细描述
// @Tags 用户操作
// @Produce json
// @Param user formData string true "账号"
// @Param pwd formData string true "密码"
// @Success 200 {string} json "{"code":200,"msg":"注册成功,请前往登录"}"
// @Router /user/reg [post]
func (c2 *c_user) reg(c *gin.Context) {
	f_user := c.PostForm("user")
	f_pwd := c.PostForm("pwd")
	if len(f_user) == 0 || len(f_pwd) == 0 {
		app.Response(c, 400, "用户名或密码格式错误")
		return
	}

	if err := c2.dao.Reg(f_user, f_pwd); err != nil {
		app.Response(c, 400, "用户名已存在")
		return
	}
	app.Response(c, 200, "注册成功,请前往登录")
}

// @Summary 修改密码
// @Description 详细描述
// @Tags 用户操作
// @Produce json
// @Param token query string true "登录返回的token"
// @Param oldPwd formData string true "旧密码"
// @Param newPwd formData string true "新密码"
// @Success 200 {string} json "{"code":200,"msg":"修改成功,请重新登录"}"
// @Router /user/newPwd [post]
func (c2 *c_user) newPwd(c *gin.Context) {
	f_oldPwd := c.PostForm("oldPwd")
	f_newPwd := c.PostForm("newPwd")
	f_uid := c.GetInt("uid")
	if len(f_oldPwd) == 0 || len(f_newPwd) == 0 {
		app.Response(c, 400, "密码格式错误")
		return
	}

	if err := c2.dao.NewPwd(f_uid, f_oldPwd, f_newPwd); err != nil {
		app.Response(c, 400, "原密码错误")
		return
	}
	app.Response(c, 200, "修改成功,请重新登录")
}

// @Summary 退出
// @Description 将造成所有软件token失效,暂停接口
// @Tags 用户操作
// @Produce json
// @Param token query string true "登录返回的token"
// @Success 200 {string} json "{"code":200,"msg":"修改成功,请重新登录"}"
// @Router /user/out [get]
func (c2 *c_user) out(c *gin.Context) {
	f_uid := c.GetInt("uid")
	if err := c2.dao.Out(f_uid); err != nil {
		app.Response(c, 400, "退出失败")
		return
	}
	app.Response(c, 200, "退出成功")
}

func (c2 *c_user) get(c *gin.Context) {
	f_user := c.Query("user")
	if len(f_user) == 0 {
		app.Response(c, 400, "查询条件为空")
		return
	}
	err, user := c2.dao.Get(&mode.User{User: f_user});
	if err != nil {
		app.Response(c, 400, "用户不存在")
		return
	}
	app.Response_data(c, 200, "查询成功", user)
}

// @Summary 获取用户列表
// @Description 详细描述
// @Tags 用户管理
// @Produce json
// @Param token query string true "登录返回的token"
// @Param page query string true "页数"
// @Param limit query string true "获取的数量"
// @Param uid query string false "用户ID"
// @Param state query string false "用户状态1=可用 2=不可用"
// @Param user query string false "用户名"
// @Success 200 {string} json "{"code":200,"msg":"查询成功","count":2,"data":[{"uid":1001,"user":"99v66","state":2,"time":1624982434},{"uid":1002,"user":"529887136","state":1,"time":1624983043}]}"
// @Router /user/getAll [get]
func (c2 *c_user) getAll(c *gin.Context) {
	var user mode.User

	if err := c.ShouldBind(&user); err != nil {
		app.Response(c, 400, "查询条件为空")
		return
	}
	err, list := c2.dao.GetAll(&user)
	if err != nil {
		app.Response(c, 400, "查询失败")
		return
	}
	app.Response_list(c, 200, "查询成功", list)
}

// @Summary 编辑用户信息
// @Description
// @Tags 用户管理
// @Produce json
// @Param token query string true "登录返回的token"
// @Param uid formData string false "用户ID"
// @Param state formData string false "用户状态1=可用 2=不可用"
// @Success 200 {string} json "{"code":200,"msg":"修改成功"}"
// @Router /user/edit [post]
func (c2 *c_user) edit(c *gin.Context) {
	var user mode.User
	err := c.ShouldBind(&user)
	if err != nil {
		app.Response(c, 400, "参数错误")
		return
	}
	err = c2.dao.Edit(&user)
	if err != nil {
		app.Response(c, 400, "修改失败")
		return
	}
	app.Response(c, 200, "修改成功")
}

type i_user interface {
	login(c *gin.Context)
	reg(c *gin.Context)
	newPwd(c *gin.Context)
	out(c *gin.Context)
	get(c *gin.Context)
	getAll(c *gin.Context)
	edit(c *gin.Context)
	UpInfo(c *gin.Context)
	getInfo(c *gin.Context)
}

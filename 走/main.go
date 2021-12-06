package main

import (
	"strings"
	"web/app"
	"web/app/controller"
	"web/app/dao"
	"web/conf"
	_ "web/docs"
	glog "web/log"
	"web/mysql"
	"web/util"
	"xutil/xStr"
)

// @title 验证系统 API
// @version 1.0

// @license.name 易轩软件 http://www.99v66.com
// @license.url http://www.99v66.com
// @description 统一编码格式-UTF8
// @description
// @host 127.0.0.1:52777
// @BasePath /v1
func main() {
	conf.Conf_ini("conf")
	util.InitEncrypt()
	mysql.InitDB()
	创建管理员账户()
	controller.INI()
	app.Start(xStr.Int2str(conf.Conf_Get("port").(int)))
}
func 创建管理员账户() {
	err := dao.User{}.Cread("admin", "e10adc3949ba59abbe56e057f20f883e")
	if err != nil && !strings.Contains(err.Error(),"Duplicate entry") {
		glog.GetLog().INFO("管理员账户创建失败")
	}
}

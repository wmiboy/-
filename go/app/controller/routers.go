package controller

import (
	"web/app"
	"web/app/cmw"
	"web/app/dao"
	"web/app/enum"
	"web/app/server"
)

func INI() {

	h_user := c_user{dao: dao.User{}}
	g_user := app.HTTP_V1.Group("/user")
	{
		g_user.POST("/reg", h_user.reg)
		g_user.POST("/login", h_user.login)
		g_user.POST("/newPwd", cmw.AUTH(enum.AUTH_All), h_user.newPwd)
		g_user.POST("/upInfo", cmw.AUTH(enum.AUTH_All), h_user.UpInfo)
		g_user.GET("/getInfo", cmw.AUTH(enum.AUTH_All), h_user.getInfo)
		//g_user.POST("/out", cmv.Auto_User(), h_user.out)//将造成所有软件token失效,待考虑
		g_user.Use(cmw.AUTH(enum.AUTH_admin))
		g_user.GET("/getall", h_user.getAll)
		g_user.POST("/edit", h_user.edit)
	}

	h_apply := c_apply{dao: dao.APPLY{}}
	g_apply := app.HTTP_V1.Group("/apply")
	{
		g_apply.POST("/get", cmw.CMW_Encrypt(), h_apply.Get)
		g_apply.Use(cmw.AUTH(enum.AUTH_admin))
		g_apply.GET("/getAll", h_apply.GetAll)
		g_apply.POST("/add", h_apply.Add)
		g_apply.POST("/edit", h_apply.Edit)
	}

	h_key := c_key{dao: dao.CDK{}}
	g_key := app.HTTP_V1.Group("/key")
	{
		g_key.Use(cmw.AUTH(enum.AUTH_admin))
		g_key.GET("/getAll", h_key.getAll)
		g_key.POST("/cread", h_key.cread)
		g_key.POST("/edit", h_key.edit)
	}

	h_power := c_power{&server.Power{&dao.Power{}}}
	g_power := app.HTTP_V1.Group("/power")
	{
		g_power.POST("/login", cmw.CMW_Encrypt(), h_power.Login_User)
		g_power.POST("/charge", h_power.Charge)
		g_power.POST("/check", cmw.AUTH(enum.AUTH_user), cmw.CMW_Encrypt(), h_power.Check)
		g_power.POST("/sub", cmw.AUTH(enum.AUTH_user), cmw.CMW_Encrypt(), h_power.Sub)
		g_power.Use(cmw.AUTH(enum.AUTH_admin))
		g_power.POST("/edit", h_power.Edit)
		g_power.GET("/getAll", h_power.GetAll)
	}

}

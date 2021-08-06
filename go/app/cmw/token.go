package cmw

import (
	"errors"
	"github.com/gin-gonic/gin"
	"web/app"
	"web/app/dao"
	"web/app/enum"
	"web/app/mode"
)

func check(c *gin.Context, pPower string) error {
	token := c.Query("token")
	err, a := app.GetToken(token)
	if err != nil {
		return err
	} else if pPower != enum.AUTH_All && a.Power != pPower {
		return errors.New("权限不足")
	}

	err, user := dao.User{}.Get(&mode.User{Uid: a.Uid})
	if err != nil {
		return errors.New("系统繁忙")
	} else if user.State != enum.User_State_Yes {
		return errors.New("账号已被禁用")
	} else if a.Time_upData != user.Uptime {
		return errors.New("token已失效")
	}

	c.Set("uid", a.Uid)
	c.Set("aid", a.Aid)
	c.Set("dev", a.Dev)
	return nil
}

func AUTH(pPower string) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := check(c, pPower)
		if err != nil {
			c.Abort()
			app.Response(c, 600, err.Error())
			return
		}
		c.Next()
	}
}
func AUTH_GET(pType int) string {
	if pType == enum.USER_TYPE_ADMIN {
		return enum.AUTH_admin
	}
	return enum.AUTH_user
}

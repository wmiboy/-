package server

import (
	"errors"
	"github.com/gin-gonic/gin"
	"web/app"
	"web/app/dao"
	"web/app/enum"
	"web/app/mode"
	"web/util"
	"xutil/xStr"
)

type Power struct {
	Dao *dao.Power
}

func (s *Power) Charge(c *gin.Context) error {
	user := c.PostForm("user")
	cdk := c.PostForm("cdk")
	aid := xStr.Str2int(c.PostForm("aid"))
	//效验卡密
	err, vCdk := dao.CDK{}.Get(&mode.CDK{Cdk: cdk})
	if err != nil {
		return errors.New("卡密不存在")
	} else if vCdk.State != enum.Cdk_State_KeYong {
		return errors.New("卡密已使用")
	} else if vCdk.Aid != aid && vCdk.Aid != 0 {
		return errors.New("卡密与软件不符")
	}
	//效验账户
	err, vUser := dao.User{}.Get(&mode.User{User: user})
	if err != nil {
		return errors.New("账户名不存在")
	} else if vUser.State != enum.User_State_Yes {
		return errors.New("账户已被禁用")
	}
	//充值
	err = s.Dao.Charge(vCdk, vUser, aid)
	if err != nil {
		return errors.New("充值失败")
	}
	return nil
}

func (s *Power) GetAll(c *gin.Context) (error, *[]mode.Power) {
	var p mode.Power
	err := c.ShouldBind(&p)
	if err != nil {
		return errors.New("参数异常"), &[]mode.Power{}
	}
	err, list := s.Dao.GetAll(&p)
	if err != nil {
		return errors.New("查询失败"), list
	}
	return err, list
}
func (s *Power) Edit(c *gin.Context) error {
	var p mode.Power
	err := c.ShouldBind(&p)
	if err != nil {
		return errors.New("参数异常")
	}
	err = s.Dao.Edit(&p)
	if err != nil {
		return errors.New("修改失败")
	}
	return nil
}
func (s *Power) Sub(c *gin.Context) error {
	uid, _ := c.Get("uid")
	aid, _ := c.Get("aid")
	point := xStr.Str2int(c.PostForm("point"))
	if point <= 0 {
		return errors.New("不可扣除负数")
	}

	err := s.Dao.Sub(&mode.Power{Uid: uid.(int), Aid: aid.(int)}, point)
	if err != nil {
		return errors.New("扣除失败,剩余点数不足")
	}
	return nil
}

func (s *Power) Login_User(c *gin.Context) (error, string) {
	//效验参数
	code := c.PostForm("dev")
	user := c.PostForm("user")
	pwd := c.PostForm("pwd")
	aid := xStr.Str2int(c.PostForm("aid"))
	ip := c.ClientIP()
	if len(user) == 0 || aid == 0 {
		return errors.New("参数异常"), ""
	}
	//登录账号,检查账户状态
	err, vUser := dao.User{}.Login(user, pwd)
	if err != nil {
		return errors.New("账户名或密码错误"), ""
	} else if vUser.State != enum.User_State_Yes {
		return errors.New("账户已被禁用"), ""
	}
	//效验权限
	err, vAPPly, vPower := s._check(aid, vUser.Uid, ip, code)
	if err != nil {
		return err, ""
	}
	//更新权限信息
	err = s._updata(vAPPly, vPower, ip, code)
	if err != nil {
		return errors.New("系统繁忙"), ""
	}
	//签发token
	err, token := app.SignToken(&app.AUTH{
		Uid:         vUser.Uid,
		Aid:         aid,
		Power:       enum.AUTH_user,
		Time_upData: vUser.Uptime,
		Time_Login:  vPower.Time,
		Dev:         code,
	})
	return err, token
}

func (s *Power) Login_Cdk(c *gin.Context) (error, string) {
	panic("implement me")
}
func (s *Power) Check(c *gin.Context) (error, *mode.Apply, *mode.Power) {
	aid, _ := c.Get("aid")
	uid, _ := c.Get("uid")
	dev, _ := c.Get("dev")
	code := c.PostForm("dev")

	ip := c.ClientIP()
	if dev.(string) != code {
		return errors.New("当前设备与登录设备不一致"), &mode.Apply{}, &mode.Power{}
	}
	//效验权限
	err, vAPPly, vPower := s._check(aid.(int), uid.(int), ip, code)
	return err, vAPPly, vPower
}

func (s *Power) _check(pAid, pUid int, pIp, pCode string) (error, *mode.Apply, *mode.Power) {
	//检查应用状态
	err, pApply := dao.APPLY{}.Get(pAid)
	if err != nil || pApply.Aid == 0 {
		return errors.New("应用不存在"), pApply, &mode.Power{}
	} else if pApply.State != enum.APPLY_STATE_KeYong {
		return errors.New("应用暂不可用"), pApply, &mode.Power{}
	} else if pApply.Login == enum.APPLY_Login_MianFei {
		return errors.New("应用无需登录"), pApply, &mode.Power{}
	}

	//检查权限状态
	err, pPower := s.Dao.Get(&mode.Power{Uid: pUid, Aid: pAid})
	if err != nil || pPower.Pid == 0 {
		return errors.New("账户未充值,请充值后使用"), pApply, pPower
	} else if pPower.State != enum.User_State_Yes {
		return errors.New("账户已被禁止使用此应用"), pApply, pPower
	} else if pApply.Mold == enum.APPLY_MOLD_KouShi && pPower.Day <= util.GetTime() {
		return errors.New("账户已到期,请充值后使用"), pApply, pPower
	} else if pApply.Mold == enum.APPLY_MOLD_KouDian && pPower.Point <= 0 {
		return errors.New("账户剩余点数不足,请充值后使用"), pApply, pPower
	} else if pApply.Bin == enum.APPLY_Bin_NoBin { //无需绑定设备
		return nil, pApply, pPower
	} else if pApply.Bin == enum.APPLY_Bin_BinCode && len(pPower.Code) != 0 && pPower.Code != pCode {
		return errors.New("非绑定设备无法登录"), pApply, pPower
	} else if pApply.Bin == enum.APPLY_Bin_BinIP && len(pPower.Ip) != 0 && pPower.Ip != pIp {
		return errors.New("非绑定IP无法登录"), pApply, pPower
	}
	return nil, pApply, pPower
}
func (s *Power) _updata(pApply *mode.Apply, pPower *mode.Power, pIp, pCode string) error {
	//执行到这个函数则表明
	//1.应用无需绑定机器
	//2.应用首次登陆
	//3.登录机器为账户绑定的机器
	//判断是否首登,首登则保存IP与机器码,直接签发token
	if len(pPower.Ip) == 0 && len(pPower.Code) == 0 {
		err := s.Dao.Edit(&mode.Power{Pid: pPower.Pid, Ip: pIp, Code: pCode})
		return err
	}
	//允许多点登录直接签发token
	if pApply.Login == enum.APPLY_Login_DuoDian {
		return nil
	}
	//以下为单点登录的情况
	//机器码相同,直接签发token
	if pPower.Code == pCode {
		return nil
	}
	//机器码不同,更新设备信息,并更新time作废之前签发的token
	//执行到这里则需要登录机器码与绑定机器码相同,所以不会影响绑定模式的软件
	time := util.GetTime()
	pPower.Time = time
	err := s.Dao.Edit(&mode.Power{Pid: pPower.Pid, Ip: pIp, Code: pCode, Time: time})
	return err
}

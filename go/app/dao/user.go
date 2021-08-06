package dao

import (
	"errors"
	"web/app/enum"
	"web/app/mode"
	"web/mysql"
	"web/util"
)

var t_user = "s_user"

type User struct {
}

func (d User) GetInfo(pUid int) (error, *mode.User) {
	var user mode.User
	tx := mysql.DB.Table(t_user).Where("uid=?", pUid).Take(&user)
	return tx.Error, &user
}

func (d User) UpInfo(pUid int, pMsg string) error {
	tx := mysql.DB.Table(t_user).Where("uid=?", pUid).Update("i_msg", pMsg)
	if tx.Error != nil && tx.RowsAffected == 0 {
		return errors.New("更新失败")
	}
	return tx.Error
}

func (d User) Cread(pUser, pPwd string) error {
	err := mysql.DB.Table(t_user).Create(&mode.User{
		User:  pUser,
		Pwd:   pPwd,
		State: enum.User_State_Yes,
		Time:  util.GetTime(),
		Type:  enum.USER_TYPE_ADMIN,
	}).Error
	return err
}

func (d User) Reg(pUser, pPwd string) error {
	err := mysql.DB.Table(t_user).Create(&mode.User{
		User:  pUser,
		Pwd:   pPwd,
		State: enum.User_State_Yes,
		Time:  util.GetTime(),
		Type:  enum.USER_TYPE_USER,
	}).Error
	return err
}

func (d User) Login(pUser, pPwd string) (error, *mode.User) {
	err, user := d.Get(&mode.User{
		User: pUser,
		Pwd:  pPwd,
	})
	return err, user
}

func (d User) NewPwd(pUid int, oldPwd, newPwd string) error {
	tx := mysql.DB.Table(t_user).Where("uid=? AND i_pwd=?", pUid, oldPwd).Updates(&mode.User{
		Pwd:    newPwd,
		Uptime: util.GetTime(),
	})
	if tx.Error != nil || tx.RowsAffected == 0 {
		return errors.New("原密码错误")
	}
	return nil
}

func (d User) Out(pUid int) error {
	err := mysql.DB.Table(t_user).Where("uid=?", pUid).Update("i_uptime", 0).Error
	return err
}

func (d User) Get(pU *mode.User) (error, *mode.User) {
	var u mode.User
	err := mysql.DB.Table(t_user).Where(pU).Take(&u).Error
	return err, &u
}

func (d User) GetAll(pU *mode.User) (error, *[]mode.User) {
	var u []mode.User
	limit, offset := util.Page(pU.Limit, pU.Page)

	tx := mysql.DB.Table(t_user)
	if len(pU.User) != 0 {
		tx.Where("i_user LIKE ?", "%"+pU.User+"%")
		pU.User = ""
	}

	err := tx.Where("i_type > ?", enum.USER_TYPE_ADMIN).Where(&pU).Offset(offset).Limit(limit).Select("uid", "i_user", "i_state", "i_time", "i_msg").Find(&u).Error
	return err, &u
}

func (d User) Edit(pU *mode.User) error {
	err := mysql.DB.Table(t_user).Where("i_type>?", enum.USER_TYPE_ADMIN).Updates(pU).Error
	return err
}

type i_user interface {
	Reg(pUser, pPwd string) error                 //注册
	Login(pUser, pPwd string) (error, *mode.User) //登录
	NewPwd(pUid int, oldPwd, newPwd string) error //改密
	Out(pUid int) error                           //退出
	Get(pU *mode.User) (error, *mode.User)        //获取个人信息
	GetAll(pU *mode.User) (error, *[]mode.User)   //获取用户列表
	Edit(pU *mode.User) error                     //修改用户信息
	Cread(pUser, pPwd string) error               //创建管理员
	UpInfo(pUid int, pMsg string) error           //更新用户信息
	GetInfo(pUid int) (error, *mode.User)         //获取用户信息
}

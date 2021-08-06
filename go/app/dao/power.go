package dao

import (
	"errors"
	"gorm.io/gorm"
	"web/app/enum"
	"web/app/mode"
	"web/mysql"
	"web/util"
)

var t_power = "s_power"

type Power struct {
}

func (p Power) Get(pP *mode.Power) (error, *mode.Power) {
	var power mode.Power
	err := mysql.DB.Table(t_power).Where(pP).Take(&power).Error
	return err, &power
}

func (p Power) Cread(pP *mode.Power) error {
	err := mysql.DB.Table(t_power).Create(pP).Error
	return err
}

func (p Power) Charge(pCdk *mode.CDK, pUser *mode.User,aid int) error {
	err, vPower := p.Get(&mode.Power{Uid: pUser.Uid, Aid: aid}) //查询权限,不存在则创建
	if err != nil {
		vPower = &mode.Power{Uid: pUser.Uid, Aid: aid, Day: util.GetTime(), State: enum.Cdk_State_KeYong, Time: util.GetTime()}
		err = p.Cread(vPower)
		if err != nil {
			return err
		}
	}

	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		tx1 := tx.Table(t_cdk).Where("kid=? AND i_state=?", pCdk.Kid, enum.Cdk_State_KeYong).Updates(map[string]interface{}{"i_state": enum.Cdk_State_BuKeYong, "i_uid": pUser.Uid, "i_use_time": util.GetTime()})
		if tx1.Error != nil || tx1.RowsAffected == 0 { //将卡密设为已使用
			return errors.New("卡密已使用")
		}
		tx1 = tx.Table(t_power).Where("i_uid=? AND i_aid=?", pUser.Uid, aid).Updates(map[string]interface{}{"i_day": gorm.Expr("i_day + ?", pCdk.Day*3600), "i_point": gorm.Expr("i_point+?", pCdk.Point)})
		return tx1.Error
	})
	return err
}

func (p Power) GetAll(pP *mode.Power) (error, *[]mode.Power) {
	var list []mode.Power
	limit, page := util.Page(pP.Limit, pP.Page)
	err := mysql.DB.Table(t_power).Where(pP).Limit(limit).Offset(page).Find(&list).Error
	return err, &list
}

func (p Power) Edit(pP *mode.Power) error {
	err := mysql.DB.Table(t_power).Updates(pP).Error
	return err
}
func (p Power) Sub(pP *mode.Power, point int) error {
	tx := mysql.DB.Table(t_power).Where(pP).Where("i_point-?>=0", point).Update("i_point", gorm.Expr("i_point-?", point))
	if tx.Error != nil || tx.RowsAffected == 0 {
		return errors.New("扣除失败")
	}
	return nil
}

type i_power interface {
	Get(pP *mode.Power) (error, *mode.Power)
	Cread(pP *mode.Power) error
	Charge(pCdk *mode.CDK, pUser *mode.User) error
	GetAll(pP *mode.Power) (error, *[]mode.Power)
	Edit(pP *mode.Power) error
	Sub(pP *mode.Power, point int) error
}

package dao

import (
	"web/app/enum"
	"web/app/mode"
	"web/mysql"
	"web/util"
)

var t_apply = "s_apply"

type APPLY struct {

}

func (A APPLY) Get(pAid int) (error, *mode.Apply) {
	var a mode.Apply
	err := mysql.DB.Table(t_apply).Where("aid=?", pAid).Take(&a).Error
	return err, &a
}

func (A APPLY) ADD(pA *mode.Apply) error {
	pA.State = enum.APPLY_STATE_KeYong
	pA.Time = util.GetTime()
	err := mysql.DB.Table(t_apply).Create(pA).Error
	return err
}

func (A APPLY) Edit(pA *mode.Apply) error {
	err := mysql.DB.Table(t_apply).Updates(pA).Error
	return err
}

func (A APPLY) GetAll(pA *mode.Apply) (error, *[]mode.Apply) {
	var list []mode.Apply
	limit, offset := util.Page(pA.Limit, pA.Page)

	tx := mysql.DB.Table(t_apply)
	if len(pA.Name) != 0 {
		tx.Where("i_name LIKE ?", "%"+pA.Name+"%")
		pA.Name = ""
	}
	err := tx.Where(&pA).Offset(offset).Limit(limit).Find(&list).Error
	return err, &list
}

type i_apply interface {
	Get(pAid int) (error, *mode.Apply)
	ADD(pA *mode.Apply) error
	Edit(pA *mode.Apply) error
	GetAll(pA *mode.Apply) (error, *[]mode.Apply)
}

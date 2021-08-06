package dao

import (
	"github.com/google/uuid"
	"web/app/enum"
	"web/app/mode"
	"web/mysql"
	"web/util"
)

type CDK struct {

}

var t_cdk = "s_key"

func (C CDK) Cread(pNum int, pKey *mode.CDK) (error, *[]mode.CDK) {
	var list []mode.CDK
	for i := 0; i < pNum; i++ {
		list = append(list, mode.CDK{
			Cdk:        uuid.New().String(),
			Mold:       pKey.Mold,
			Point:      pKey.Point,
			Day:        pKey.Day,
			Uid:        0,
			Aid:        pKey.Aid,
			Cread_time: util.GetTime(),
			Use_time:   0,
			State:      enum.Cdk_State_KeYong,
		})
	}
	err := mysql.DB.Table(t_cdk).Create(&list).Error
	return err, &list
}

func (C CDK) GetAll(pKey *mode.CDK) (error, *[]mode.CDK) {
	var k []mode.CDK
	limit, page := util.Page(pKey.Limit, pKey.Page)
	err := mysql.DB.Table(t_cdk).Where(pKey).Limit(limit).Offset(page).Find(&k).Error
	return err, &k
}

func (C CDK) Get(pKey *mode.CDK) (error, *mode.CDK) {
	var k mode.CDK
	err := mysql.DB.Table(t_cdk).Where(pKey).Take(&k).Error
	return err, &k
}

func (C CDK) Edit(pKey *mode.CDK) error {
	err := mysql.DB.Table(t_cdk).Updates(pKey).Error
	return err
}

type i_key interface {
	Cread(pNum int, pKey *mode.CDK) (error, *[]mode.CDK)
	GetAll(pKey *mode.CDK) (error, *[]mode.CDK)
	Get(pKey *mode.CDK) (error, *mode.CDK)
	Edit(pKey *mode.CDK) error
}

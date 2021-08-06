package mode

type CDK struct {
	Kid        int    `gorm:"primary_key" json:"kid" form:"kid"`
	Cdk        string `gorm:"column:i_cdk" json:"cdk" form:"cdk"`
	Mold       int    `gorm:"column:i_mold" json:"mold" form:"mold"`
	Point      int    `gorm:"column:i_point" json:"point" form:"point"`
	Day        int    `gorm:"column:i_day" json:"day" form:"day"`
	Uid        int    `gorm:"column:i_uid" json:"uid" form:"uid"`
	Aid        int    `gorm:"column:i_aid" json:"aid" form:"aid"`
	Cread_time int    `gorm:"column:i_cread_time" json:"cread_time"`
	Use_time   int    `gorm:"column:i_use_time" json:"use_time"`
	State      int    `gorm:"column:i_state" json:"state" form:"state"`
	Page       int    `gorm:"-" form:"page" json:"-"`
	Limit      int    `gorm:"-" form:"limit" json:"-"`
}

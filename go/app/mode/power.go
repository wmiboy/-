package mode

type Power struct {
	Pid   int    `gorm:"primary_key" json:"pid" form:"pid"`
	Uid   int    `gorm:"column:i_uid" json:"uid" form:"uid"`
	Aid   int    `gorm:"column:i_aid" json:"aid" form:"aid"`
	Kid   int    `gorm:"column:i_kid" json:"kid" form:"kid"`
	Day   int    `gorm:"column:i_day" json:"day" form:"day"`
	Point int    `gorm:"column:i_point" json:"point" form:"point"`
	Ip    string `gorm:"column:i_ip" json:"ip" form:"ip"`
	Code  string `gorm:"column:i_code" json:"code" form:"code"`
	State int    `gorm:"column:i_state" json:"state" form:"state"`
	Time  int    `gorm:"column:i_time" json:"time"`
	Page  int    `gorm:"-" form:"page" json:"-"`
	Limit int    `gorm:"-" form:"limit" json:"-"`
}
package mode

type Apply struct {
	Aid   int    `gorm:"primary_key" json:"aid" form:"aid"`
	Name  string `gorm:"column:i_name" json:"name" form:"name"`
	Mold  int    `gorm:"column:i_mold" json:"mold" form:"mold"`
	Login int    `gorm:"column:i_login" json:"login" form:"login"`
	Bin   int    `gorm:"column:i_bin" json:"bin" form:"bin"`
	Msg   string `gorm:"column:i_msg" json:"msg" form:"msg"`
	State int    `gorm:"column:i_state" json:"state" form:"state"`
	Time  int    `gorm:"column:i_time" json:"time"`
	Page  int    `gorm:"-" form:"page" json:"-"`
	Limit int    `gorm:"-" form:"limit" json:"-"`
}

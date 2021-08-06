package mode

type User struct {
	Uid    int    `gorm:"primary_key" json:"uid" form:"uid"`
	User   string `gorm:"column:i_user" json:"user" form:"user"`
	Pwd    string `gorm:"column:i_pwd" json:"-"`
	Msg    string `gorm:"column:i_msg" json:"msg"`
	Type   int    `gorm:"column:i_type" json:"type" form:"type"`
	State  int    `gorm:"column:i_state" json:"state" form:"state"`
	Time   int    `gorm:"column:i_time" json:"time"`
	Uptime int    `gorm:"column:i_uptime" json:"-"`
	Page   int    `gorm:"-" form:"page" json:"-"`
	Limit  int    `gorm:"-" form:"limit" json:"-"`
}

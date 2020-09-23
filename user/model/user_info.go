package model

type UserInfo struct {
	ID         int64  `gorm:"column:id;primary_key" json:"id"`
	UID        int    `gorm:"column:uid" json:"uid"`
	Username   string `gorm:"column:username" json:"username"`
	Nickname   string `gorm:"column:nickname" json:"nickname"`
	Password   string `gorm:"column:password" json:"password"`
	Salt       string `gorm:"column:salt" json:"salt"`
	Type       int    `gorm:"column:type" json:"type"`
	PublicKey  string `gorm:"column:public_key" json:"public_key"`
	Status     int    `gorm:"column:status" json:"status"`
	CreateTime int    `gorm:"column:create_time" json:"create_time"`
	UpdateTime int    `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (u *UserInfo) TableName() string {
	return "t_user_info"
}

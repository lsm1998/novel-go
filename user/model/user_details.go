package model

type UserDetails struct {
	ID         int64  `gorm:"column:id;primary_key" json:"id"`
	UID        int    `gorm:"column:uid" json:"uid"`
	HeadImg    string `gorm:"column:head_img" json:"head_img"`
	Phone      string `gorm:"column:phone" json:"phone"`
	LoginType  int    `gorm:"column:login_type" json:"login_type"`
	Birthday   int    `gorm:"column:birthday" json:"birthday"`
	Region     string `gorm:"column:region" json:"region"`
	Nation     string `gorm:"column:nation" json:"nation"`
	Autograph  string `gorm:"column:autograph" json:"autograph"`
	Status     int    `gorm:"column:status" json:"status"`
	CreateTime int    `gorm:"column:create_time" json:"create_time"`
	UpdateTime int    `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (u *UserDetails) TableName() string {
	return "t_user_details"
}

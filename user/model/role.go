package model

type Role struct {
	ID         int64  `gorm:"column:id;primary_key" json:"id"`
	RoleName   string `gorm:"column:role_name" json:"role_name"`
	Remakes    string `gorm:"column:remakes" json:"remakes"`
	Status     int    `gorm:"column:status" json:"status"`
	CreateTime int    `gorm:"column:create_time" json:"create_time"`
	UpdateTime int    `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (r *Role) TableName() string {
	return "t_role"
}

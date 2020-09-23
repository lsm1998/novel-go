package model

type Auth struct {
	ID         int64 `gorm:"column:id;primary_key" json:"id"`
	RoleID     int64 `gorm:"column:role_id" json:"role_id"`
	UserID     int64 `gorm:"column:user_id" json:"user_id"`
	Status     int   `gorm:"column:status" json:"status"`
	CreateTime int   `gorm:"column:create_time" json:"create_time"`
	UpdateTime int   `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (a *Auth) TableName() string {
	return "t_auth"
}

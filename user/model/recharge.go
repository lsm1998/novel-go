package model

type Recharge struct {
	ID         int64 `gorm:"column:id;primary_key" json:"id"`
	UID        int   `gorm:"column:uid" json:"uid"`
	Amount     int   `gorm:"column:amount" json:"amount"`
	Type       int64 `gorm:"column:type" json:"type"`
	Status     int   `gorm:"column:status" json:"status"`
	CreateTime int   `gorm:"column:create_time" json:"create_time"`
	UpdateTime int   `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (r *Recharge) TableName() string {
	return "t_recharge"
}

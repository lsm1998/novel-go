package model

type CircleLike struct {
	ID         int64 `gorm:"column:id;primary_key" json:"id"`
	UID        int   `gorm:"column:uid" json:"uid"`
	Type       int   `gorm:"column:type" json:"type"`
	Status     int   `gorm:"column:status" json:"status"`
	CreateTime int   `gorm:"column:create_time" json:"create_time"`
	UpdateTime int   `gorm:"column:update_time" json:"update_time"`
	CircleID   int64 `gorm:"column:circle_id" json:"circle_id"`
}

// TableName sets the insert table name for this struct type
func (c *CircleLike) TableName() string {
	return "t_circle_like"
}

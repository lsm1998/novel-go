package model

type Follow struct {
	ID         int64 `gorm:"column:id;primary_key" json:"id"`
	FollowID   int   `gorm:"column:follow_id" json:"follow_id"`
	FansID     int   `gorm:"column:fans_id" json:"fans_id"`
	Status     int   `gorm:"column:status" json:"status"`
	CreateTime int   `gorm:"column:create_time" json:"create_time"`
	UpdateTime int   `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (f *Follow) TableName() string {
	return "t_follow"
}

package model

type CircleComment struct {
	ID         int64  `gorm:"column:id;primary_key" json:"id"`
	CircleID   int64  `gorm:"column:circle_id" json:"circle_id"`
	UID        int    `gorm:"column:uid" json:"uid"`
	Mast       int64  `gorm:"column:mast" json:"mast"`
	ReplyID    int64  `gorm:"column:reply_id" json:"reply_id"`
	Content    string `gorm:"column:content" json:"content"`
	Status     int    `gorm:"column:status" json:"status"`
	CreateTime int    `gorm:"column:create_time" json:"create_time"`
	UpdateTime int    `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (c *CircleComment) TableName() string {
	return "t_circle_comment"
}

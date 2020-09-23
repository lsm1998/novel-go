package model

type Circle struct {
	ID         int64  `gorm:"column:id;primary_key" json:"id"`
	UID        int    `gorm:"column:uid" json:"uid"`
	BookID     int64  `gorm:"column:book_id" json:"book_id"`
	Content    string `gorm:"column:content" json:"content"`
	Status     int    `gorm:"column:status" json:"status"`
	CreateTime int    `gorm:"column:create_time" json:"create_time"`
	UpdateTime int    `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (c *Circle) TableName() string {
	return "t_circle"
}

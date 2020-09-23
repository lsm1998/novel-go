package model

type Chapter struct {
	ID         int64  `gorm:"column:id;primary_key" json:"id"`
	BookID     int64  `gorm:"column:book_id" json:"book_id"`
	Sorted     int    `gorm:"column:sorted" json:"sorted"`
	Name       string `gorm:"column:name" json:"name"`
	Content    string `gorm:"column:content" json:"content"`
	Status     int    `gorm:"column:status" json:"status"`
	CreateTime int    `gorm:"column:create_time" json:"create_time"`
	UpdateTime int    `gorm:"column:update_time" json:"update_time"`
	ThirdID    int64  `gorm:"column:third_id" json:"third_id"`
	PlatformID int    `gorm:"column:platform_id" json:"platform_id"`
}

// TableName sets the insert table name for this struct type
func (c *Chapter) TableName() string {
	return "t_chapter"
}

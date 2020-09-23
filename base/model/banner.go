package model

// banner
type Banner struct {
	ID         int64  `gorm:"column:id;primary_key" json:"id"`
	Title      string `gorm:"column:title" json:"title"`
	Link       string `gorm:"column:link" json:"link"`
	Img        string `gorm:"column:img" json:"img"`
	Enable     int    `gorm:"column:enable" json:"enable"`
	Status     int    `gorm:"column:status" json:"status"`
	CreateTime int    `gorm:"column:create_time" json:"create_time"`
	UpdateTime int    `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (b *Banner) TableName() string {
	return "t_banner"
}

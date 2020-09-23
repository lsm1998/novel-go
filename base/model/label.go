package model

// 小说标签
type Label struct {
	ID         int64  `gorm:"column:id;primary_key" json:"id"`
	Name       string `gorm:"column:name" json:"name"`
	Status     int    `gorm:"column:status" json:"status"`
	CreateTime int    `gorm:"column:create_time" json:"create_time"`
	UpdateTime int    `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (l *Label) TableName() string {
	return "t_label"
}

package model

// 小说分类
type Type struct {
	ID         int64  `gorm:"column:id;primary_key" json:"id"`
	Pid        int64  `gorm:"column:pid" json:"pid"`
	Pic        string `gorm:"column:pic" json:"pic"`
	Name       string `gorm:"column:name" json:"name"`
	Status     int    `gorm:"column:status" json:"status"`
	CreateTime int    `gorm:"column:create_time" json:"create_time"`
	UpdateTime int    `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (t *Type) TableName() string {
	return "t_type"
}

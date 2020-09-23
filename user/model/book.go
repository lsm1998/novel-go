package model

type Book struct {
	ID          int64  `gorm:"column:id;primary_key" json:"id"`
	TypeID      int64  `gorm:"column:type_id" json:"type_id"`
	Title       string `gorm:"column:title" json:"title"`
	Synopsis    string `gorm:"column:synopsis" json:"synopsis"`
	Cover       string `gorm:"column:cover" json:"cover"`
	Recommend   int    `gorm:"column:recommend" json:"recommend"`
	Click       int    `gorm:"column:click" json:"click"`
	Collection  int    `gorm:"column:collection" json:"collection"`
	Instalments int    `gorm:"column:instalments" json:"instalments"`
	WordNum     int64  `gorm:"column:word_num" json:"word_num"`
	Status      int    `gorm:"column:status" json:"status"`
	CreateTime  int    `gorm:"column:create_time" json:"create_time"`
	UpdateTime  int    `gorm:"column:update_time" json:"update_time"`
	ThirdID     int64  `gorm:"column:third_id" json:"third_id"`
	PlatformID  int    `gorm:"column:platform_id" json:"platform_id"`
	IsSerial    int    `gorm:"column:is_serial" json:"is_serial"`
}

// TableName sets the insert table name for this struct type
func (b *Book) TableName() string {
	return "t_book"
}

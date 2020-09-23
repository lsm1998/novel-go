package model

type Message struct {
	ID         int64  `gorm:"column:id;primary_key" json:"id"`
	Cmd        int    `gorm:"column:cmd" json:"cmd"`
	Len        int    `gorm:"column:len" json:"len"`
	FormID     int    `gorm:"column:form_id" json:"form_id"`
	ToID       int    `gorm:"column:to_id" json:"to_id"`
	Content    string `gorm:"column:content" json:"content"`
	Status     int    `gorm:"column:status" json:"status"`
	CreateTime int    `gorm:"column:create_time" json:"create_time"`
	UpdateTime int    `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (m *Message) TableName() string {
	return "t_message"
}

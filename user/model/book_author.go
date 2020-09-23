package model

type BookAuthor struct {
	ID         int64 `gorm:"column:id;primary_key" json:"id"`
	UserID     int64 `gorm:"column:user_id" json:"user_id"`
	BookID     int64 `gorm:"column:book_id" json:"book_id"`
	Status     int   `gorm:"column:status" json:"status"`
	CreateTime int   `gorm:"column:create_time" json:"create_time"`
	UpdateTime int   `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (b *BookAuthor) TableName() string {
	return "t_book_author"
}

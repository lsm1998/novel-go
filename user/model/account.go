package model

type Account struct {
	ID               int64 `gorm:"column:id;primary_key" json:"id"`
	UID              int   `gorm:"column:uid" json:"uid"`
	AccountBalance   int64 `gorm:"column:account_balance" json:"account_balance"`
	MonthlyBalance   int64 `gorm:"column:monthly_balance" json:"monthly_balance"`
	RecommendBalance int64 `gorm:"column:recommend_balance" json:"recommend_balance"`
	Status           int   `gorm:"column:status" json:"status"`
	CreateTime       int   `gorm:"column:create_time" json:"create_time"`
	UpdateTime       int   `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (a *Account) TableName() string {
	return "t_account"
}

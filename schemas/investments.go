package schemas

type Investment struct {
	ID           uint    `gorm:"primarykey"`
	InvestmentId string  `gorm:"not null"`
	Name         string  `gorm:"not null"`
	Description  string  `gorm:"not null"`
	Type         string  `gorm:"not null"`
	Wallet       string  `gorm:"not null"`
	Amount       int32   `gorm:"not null"`
	Rate         float32 `gorm:"not null"`
	Liquidity    string  `gorm:"not null"`
	BoughtAt     string  `gorm:"not null"`
	ExpiresAt    string  `gorm:"not null"`
	IsActive     bool    `gorm:"not null"`
}

func (Investment) TableName() string {
	return "finance.investmens"
}

package schemas

type Investment struct {
	ID        uint `gorm:"primarykey"`
	Name      string
	Type      string
	Wallet    string
	Amount    int32
	Rate      float32
	BoughtAt  string
	ExpiresAt string
	IsActive  bool
}

func (Investment) TableName() string {
	return "finance.investmens"
}

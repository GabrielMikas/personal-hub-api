package schemas

type Card struct {
	ID           uint   `gorm:"primarykey"`
	CardId       string `gorm:"not null"`
	Name         string `gorm:"not null"`
	Description  string
	CollectionId string `gorm:"not null"`
	Type         string `gorm:"not null"`
	CardCode     string `gorm:"not null"`
	ImageUrl     string
	Amount       int32  `gorm:"not null"`
	Rarity       string `gorm:"not null"`
	BoughtAt     string `gorm:"not null"`
	IsSleeved    bool   `gorm:"not null"`
	IsInBinder   bool   `gorm:"not null"`
}

func (Card) TableName() string {
	return "hobbies.cards"
}

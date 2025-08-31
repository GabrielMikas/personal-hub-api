package schemas

type Collection struct {
	ID           uint   `gorm:"primarykey"`
	CollectionId string `gorm:"not null"`
	Name         string `gorm:"not null"`
	Type         string `gorm:"not null"`
	LaunchYear   string `gorm:"not null"`
}

func (Collection) TableName() string {
	return "hobbies.collections"
}

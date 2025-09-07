package cards_handler

type Card struct {
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description"`
	CollectionId string `json:"collectionId" binding:"required"`
	Type         string `json:"type" binding:"required"`
	CardCode     string `json:"cardCode" binding:"required"`
	ImageUrl     string `json:"imageUrl"`
	Amount       int32  `json:"amount" binding:"required"`
	Rarity       string `json:"rarity" binding:"required"`
	BoughtAt     string `json:"boughtAt" binding:"required"`
	IsSleeved    bool   `json:"isSleeved" binding:"required"`
	IsInBinder   bool   `json:"isInBinder" binding:"required"`
}

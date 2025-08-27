package cards_handler

type Card struct {
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description" binding:"required"`
	CollectionId string `json:"collectionId" binding:"required"`
	Type         string `json:"type" binding:"required"`
	CardCode     string `json:"cardCode" binding:"required"`
	Amount       int32  `json:"amount" binding:"required"`
	Rarity       string `json:"rarity" binding:"required"`
	BoughtAt     string `json:"boughtAt" binding:"required"`
	IsSleeved    bool   `json:"isSleeved" binding:"required"`
}

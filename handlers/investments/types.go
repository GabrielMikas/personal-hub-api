package investments_handler

type Investment struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Type        string  `json:"type" binding:"required"`
	Wallet      string  `json:"wallet" binding:"required"`
	Amount      int32   `json:"amount" binding:"required"`
	Rate        float32 `json:"rate" binding:"required"`
	BoughtAt    string  `json:"boughtAt" binding:"required"`
	ExpiresAt   string  `json:"expiresAt" binding:"required"`
	IsActive    bool    `json:"isActive" binding:"required"`
}

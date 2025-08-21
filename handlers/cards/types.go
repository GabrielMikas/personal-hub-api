package cards_handler

type Card struct {
	Name string `json:"name" binding:"required"`
}

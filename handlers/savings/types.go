package savings_handler

type Savings struct {
	Name string `json:"name" binding:"required"`
}

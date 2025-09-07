package collections_handlers

type Collection struct {
	Name       string `json:"name" binding:"required"`
	LaunchYear string `json:"launchYear" binding:"required"`
	Type       string `json:"type" binding:"required"`
}

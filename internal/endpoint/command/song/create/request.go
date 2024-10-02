package create

type Request struct {
	Group string `json:"group" binding:"required"`
	Song  string `json:"song" binding:"required"`
}

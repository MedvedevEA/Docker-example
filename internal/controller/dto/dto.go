package dto

type AddRecord struct {
	Text string `json:"text" binding:"required"`
}

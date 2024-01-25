package models

type ResUrl struct {
	Url string `json:"result" binding:"required"`
}

package models

type SaveUser struct {
	ID       int    `json:"ID" binding:"required"`
	UserName string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

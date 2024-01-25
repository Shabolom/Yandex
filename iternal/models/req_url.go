package models

type ReqUrl struct {
	Url string `json:"url" binding:"required"`
}

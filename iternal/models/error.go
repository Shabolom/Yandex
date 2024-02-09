package models

import "YandexPra/iternal/domain"

type Error struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type ErrorBatch struct {
	Code  int           `json:"code"`
	Error string        `json:"error"`
	Body  []domain.Urls `json:"duplicate"`
}

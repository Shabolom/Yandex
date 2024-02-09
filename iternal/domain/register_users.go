package domain

type RegisterUsers struct {
	Base
	Login    string `gorm:"colum:login; type:text"`
	Password string `gorm:"colum:password; type:text"`
}

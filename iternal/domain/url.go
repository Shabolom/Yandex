package domain

type Urls struct {
	Url   string `gorm:"column:url" gorm:"type:text"`
	Short string `gorm:"column:short" gorm:"type:text"`
}

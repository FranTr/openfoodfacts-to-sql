package product

import "gorm.io/gorm"

type LanguageCodes struct {
	gorm.Model
	Lc        string `gorm:"type:text"`
	ProductId string `gorm:"type:char(60)"`
}

func (l *LanguageCodes) UnmarshalJSON(b []byte) (err error) {
	l.Lc = string(b)
	return nil
}

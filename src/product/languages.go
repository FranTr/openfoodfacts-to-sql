package product

import "gorm.io/gorm"

type Languages struct {
	gorm.Model
	L         string `gorm:"type:text"`
	ProductId string `gorm:"type:char(60)"`
}

func (l *Languages) UnmarshalJSON(b []byte) (err error) {
	l.L = string(b)
	return nil
}

package product

import (
	"gorm.io/gorm"
)

type NutrientLevel struct {
	gorm.Model
	Id           string `gorm:"unique"`
	Sugars       string `json:"sugars"`
	SaturatedFat string `json:"saturated-fat"`
	Fat          string `json:"fat"`
	Salt         string `json:"salt"`
}

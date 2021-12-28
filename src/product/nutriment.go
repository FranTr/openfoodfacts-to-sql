package product

import (
	"gorm.io/gorm"
)

type Nutriment struct {
	gorm.Model
	Id                   string  `gorm:"unique"`
	Carbohydrates100G    float64 `json:"carbohydrates_100g"`
	ProteinsValue        float64 `json:"proteins_value"`
	CarbohydratesUnit    string  `json:"carbohydrates_unit"`
	Cholesterol100G      float64 `json:"cholesterol_100g"`
	Fat100G              float64 `json:"fat_100g"`
	Carbohydrates        float64 `json:"carbohydrates"`
	EnergyUnit           string  `json:"energy_unit"`
	SugarsUnit           string  `json:"sugars_unit"`
	ProteinsUnit         string  `json:"proteins_unit"`
	Sugars100G           float64 `json:"sugars_100g"`
	Proteins100G         float64 `json:"proteins_100g"`
	SodiumUnit           string  `json:"sodium_unit"`
	ProteinsServing      float64 `json:"proteins_serving"`
	EnergyServing        float64 `json:"energy_serving"`
	Proteins             float64 `json:"proteins"`
	FiberUnit            string  `json:"fiber_unit"`
	FatUnit              string  `json:"fat_unit"`
	Sodium100G           float64 `json:"sodium_100g"`
	Sodium               float64 `json:"sodium"`
	FiberValue           float64 `json:"fiber_value"`
	FatServing           float64 `json:"fat_serving"`
	CholesterolUnit      string  `json:"cholesterol_unit"`
	Cholesterol          float64 `json:"cholesterol"`
	EnergyKcalValue      float64 `json:"energy-kcal_value"`
	Fat                  float64 `json:"fat"`
	EnergyKcal           float64 `json:"energy-kcal"`
	SaturatedFat         float64 `json:"saturated-fat"`
	SaltUnit             string  `json:"salt_unit"`
	Salt                 float64 `json:"salt"`
	SugarsServing        float64 `json:"sugars_serving"`
	EnergyKcalUnit       string  `json:"energy-kcal_unit"`
	SaltValue            float64 `json:"salt_value"`
	SodiumServing        float64 `json:"sodium_serving"`
	EnergyValue          float64 `json:"energy_value"`
	SodiumValue          float64 `json:"sodium_value"`
	EnergyKcalServing    float64 `json:"energy-kcal_serving"`
	SaturatedFat100G     float64 `json:"saturated-fat_100g"`
	FatValue             float64 `json:"fat_value"`
	Salt100G             float64 `json:"salt_100g"`
	SaturatedFatUnit     string  `json:"saturated-fat_unit"`
	CarbohydratesServing float64 `json:"carbohydrates_serving"`
	Fiber100G            float64 `json:"fiber_100g"`
	SaltServing          float64 `json:"salt_serving"`
	Sugars               float64 `json:"sugars"`
	Energy               float64 `json:"energy"`
	FiberServing         float64 `json:"fiber_serving"`
	Energy100G           float64 `json:"energy_100g"`
	EnergyKcal100G       float64 `json:"energy-kcal_100g"`
	SugarsValue          float64 `json:"sugars_value"`
	CholesterolServing   float64 `json:"cholesterol_serving"`
	Fiber                float64 `json:"fiber"`
	SaturatedFatServing  float64 `json:"saturated-fat_serving"`
	CarbohydratesValue   float64 `json:"carbohydrates_value"`
	SaturatedFatValue    float64 `json:"saturated-fat_value"`
	CholesterolValue     float64 `json:"cholesterol_value"`
}

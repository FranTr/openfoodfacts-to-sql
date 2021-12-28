package product

import (
	"encoding/json"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Id                                          string         `json:"id" gorm:"type:char(60)"`
	Code                                        string         `json:"code" gorm:"type:text"`
	Brands                                      string         `json:"brands" gorm:"type:text"`
	BrandsTags                                  pq.StringArray `json:"brands_tags" gorm:"type:mediumtext"`
	GenericName                                 string         `json:"generic_name" gorm:"type:mediumtext"`
	GenericNameEn                               string         `json:"generic_name_en" gorm:"type:mediumtext"`
	CreatedTime                                 EpochTime      `json:"created_t" gorm:"foreignKey:CreatedTimeId"`
	LastImageTime                               EpochTime      `json:"last_image_t" gorm:"foreignKey:LastImageId"`
	LastModifiedTime                            EpochTime      `json:"last_modified_t" gorm:"foreignKey:LastModificationId"`
	CompletedTime                               EpochTime      `json:"completed_t" gorm:"foreignKey:CompletedId"`
	LastModifiedBy                              string         `json:"last_modified_by" gorm:"type:text"`
	Additives                                   string         `json:"additives" gorm:"type:mediumtext"`
	AdditivesDebugTags                          pq.StringArray `json:"additives_debug_tags" gorm:"type:mediumtext"`
	AdditivesNumber                             int            `json:"additives_n"`
	AdditivesOldNumber                          int            `json:"additives_old_n"`
	AdditivesOldTags                            pq.StringArray `json:"additives_old_tags" gorm:"type:mediumtext"`
	AdditivesPrev                               string         `json:"additives_prev" gorm:"type:mediumtext"`
	AdditivesPrevNumber                         int            `json:"additives_prev_n"  gorm:"type:mediumtext"`
	AdditivesPrevTags                           pq.StringArray `json:"additives_prev_tags" gorm:"type:mediumtext"`
	AdditivesTags                               pq.StringArray `json:"additives_tags" gorm:"type:mediumtext"`
	Allergens                                   string         `json:"allergens" gorm:"type:mediumtext"`
	AllergensHierarchy                          pq.StringArray `json:"allergens_hierarchy" gorm:"type:mediumtext"`
	AllergensTags                               pq.StringArray `json:"allergens_tags" gorm:"type:mediumtext"`
	Categories                                  string         `json:"categories"`
	CategoriesDebugTags                         pq.StringArray `json:"categories_debug_tags" gorm:"type:text"`
	CategoriesHierarchy                         pq.StringArray `json:"categories_hierarchy" gorm:"type:text"`
	CategoriesTags                              pq.StringArray `json:"categories_tags" gorm:"type:text"`
	Complete                                    int            `json:"complete"`
	Countries                                   string         `json:"countries"`
	CountriesHierarchy                          pq.StringArray `json:"countries_hierarchy" gorm:"type:text"`
	CountriesTags                               pq.StringArray `json:"countries_tags" gorm:"type:text"`
	EmbCodes                                    string         `json:"emb_codes"`
	EcoscoreGrade                               string         `json:"ecoscore_grade"`
	ExpirationDate                              string         `json:"expiration_date"`
	FruitsVegetablesNuts100GEstimate            json.Number    `json:"fruits-vegetables-nuts_100g_estimate"`
	Ingredients                                 []Ingredient   `json:"ingredients" gorm:"foreignKey:ProductId"`
	IngredientsFromOrThatMayBeFromPalmOilNumber int            `json:"ingredients_from_or_that_may_be_from_palm_oil_n"`
	IngredientsFromPalmOilNumber                int            `json:"ingredients_from_palm_oil_n"`
	IngredientsFromPalmOilTags                  pq.StringArray `json:"ingredients_from_palm_oil_tags" gorm:"type:mediumtext"`
	IngredientsIdsDebug                         pq.StringArray `json:"ingredients_ids_debug" gorm:"type:mediumtext"`
	IngredientsTags                             pq.StringArray `json:"ingredients_tags" gorm:"type:mediumtext"`
	IngredientsText                             string         `json:"ingredients_text" gorm:"type:mediumtext"`
	IngredientsTextEn                           string         `json:"ingredients_text_en" gorm:"type:mediumtext"`
	IngredientsTextWithAllergens                string         `json:"ingredients_text_with_allergens" gorm:"type:mediumtext"`
	IngredientsTextWithAllergensEn              string         `json:"ingredients_text_with_allergens_en" gorm:"type:mediumtext"`
	IngredientsThatMayBeFromPalmOilNumber       int            `json:"ingredients_that_may_be_from_palm_oil_n"`
	IngredientsThatMayBeFromPalmOilTags         pq.StringArray `json:"ingredients_that_may_be_from_palm_oil_tags" gorm:"type:text"`
	Keywords                                    pq.StringArray `json:"_keywords" gorm:"type:text"`
	Labels                                      string         `json:"labels"`
	LabelsDebugTags                             pq.StringArray `json:"labels_debug_tags" gorm:"type:text"`
	LabelsHierarchy                             pq.StringArray `json:"labels_hierarchy" gorm:"type:text"`
	LabelsPrevHierarchy                         pq.StringArray `json:"labels_prev_hierarchy" gorm:"type:text"`
	LabelsPrevTags                              pq.StringArray `json:"labels_prev_tags" gorm:"type:text"`
	LabelsTags                                  pq.StringArray `json:"labels_tags" gorm:"type:text"`
	Lang                                        string         `json:"lang"`
	Languages                                   Languages      `json:"languages" gorm:"foreignKey:ProductId"`
	LanguagesCodes                              LanguageCodes  `json:"languages_codes" gorm:"foreignKey:ProductId"`
	Locale                                      string         `json:"lc"`
	ManufacturingPlaces                         string         `json:"manufacturing_places"`
	NewAdditivesNumber                          int            `json:"new_additives_n"`
	NutrientLevels                              NutrientLevel  `json:"nutrient_levels" gorm:"foreignKey:Id"`
	NutrientLevelsTags                          pq.StringArray `json:"nutrient_levels_tags" gorm:"type:text"`
	Nutriments                                  Nutriment      `json:"nutriments" gorm:"foreignKey:Id"`
	NutritionDataPer                            string         `json:"nutrition_data_per"`
	NutritionGrades                             string         `json:"nutrition_grades"`
	NutritionScoreDebug                         string         `json:"nutrition_score_debug"`
	NutriscoreGrade                             string         `json:"nutriscore_grade"`
	NovaGroup                                   int            `json:"nova_group"`
	Origins                                     string         `json:"origins"`
	Packaging                                   string         `json:"packaging"`
	PackagingTags                               pq.StringArray `json:"packaging_tags" gorm:"type:text"`
	PnnsGroups1                                 string         `json:"pnns_groups_1"`
	PnnsGroups2                                 string         `json:"pnns_groups_2"`
	ProductName                                 string         `json:"product_name"`
	ProductNameEn                               string         `json:"product_name_en"`
	PurchasePlaces                              string         `json:"purchase_places"`
	PurchasePlacesTags                          pq.StringArray `json:"purchase_places_tags" gorm:"type:text"`
	Quantity                                    string         `json:"quantity"`
	Rev                                         json.Number    `json:"rev"`
	ServingQuantity                             json.Number    `json:"serving_quantity"`
	ServingSize                                 string         `json:"serving_size"`
	SortKey                                     int            `json:"sortkey"`
	Stores                                      string         `json:"stores"`
	StoresTags                                  pq.StringArray `json:"stores_tags" gorm:"type:text"`
	Traces                                      string         `json:"traces"`
	TracesHierarchy                             pq.StringArray `json:"traces_hierarchy" gorm:"type:text"`
	TracesTags                                  pq.StringArray `json:"traces_tags" gorm:"type:text"`
	UniqueScansNumber                           int            `json:"unique_scans_n"`
	UnknownNutrientsTags                        pq.StringArray `json:"unknown_nutrients_tags" gorm:"type:text"`
	UpdateKey                                   string         `json:"update_key"`
}

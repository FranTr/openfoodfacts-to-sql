package product

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"log"
	"strings"
)

type Ingredient struct {
	gorm.Model
	//Id            string  `gorm:"unique"`
	ProductIdText string  `json:"id" gorm:"type:mediumtext"`
	Rank          float64 `json:"rank,omitempty"`
	Text          string  `json:"text" gorm:"type:mediumtext"`
	Vegan         string  `json:"vegan,omitempty"`
	Vegetarian    string  `json:"vegetarian,omitempty"`
	Percent       string  `json:"percent,omitempty"`
	ProductId     string  `gorm:"type:char(60)"`
}

func (i *Ingredient) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	var arr map[string]interface{}
	err = json.Unmarshal([]byte(s), &arr)

	if err != nil {
		log.Fatalf("scan file error: %v", err)
	}

	i.ProductIdText = arr["id"].(string)
	i.Text = arr["text"].(string)
	if arr["rank"] != nil {
		i.Rank = arr["rank"].(float64)
	}
	val, ok := arr["percent"].(string)
	if ok == true {
		i.Percent = val
	}
	if ok == false {
		valfl, ok := arr["percent"].(float64)
		if ok == true {
			val := fmt.Sprintf("%f", valfl)
			i.Percent = val
		}
	}
	if arr["vegan"] != nil {
		i.Vegan = arr["vegan"].(string)
	}
	if arr["vegetarian"] != nil {
		i.Vegetarian = arr["vegetarian"].(string)
	}
	return nil
}

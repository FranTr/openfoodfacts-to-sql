package main

import (
	"bufio"
	"encoding/json"
	"github.com/FranTr/openfoodfacts-to-sql/src/product"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func scanLine(scanner *bufio.Scanner, db *gorm.DB) []string {
	s := make([]string, 2004000)
	counter := 0
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	for scanner.Scan() {
		text := scanner.Text()
		prod := &product.Product{}
		err := json.Unmarshal([]byte(text), prod)

		db.Create(prod)
		if err != nil {
			log.Fatalf("scan file error: %v", err)
		}
		counter++
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
	}
	return s
}

func main() {

	dsn := "user:userpw@tcp(127.0.0.1:3307)/openfoodfacts?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&product.EpochTime{})
	err = db.AutoMigrate(&product.Nutriment{})
	err = db.AutoMigrate(&product.NutrientLevel{})
	err = db.AutoMigrate(&product.Ingredient{})
	err = db.AutoMigrate(&product.URL{})
	err = db.AutoMigrate(&product.Languages{})
	err = db.AutoMigrate(&product.LanguageCodes{})
	err = db.AutoMigrate(&product.Product{})
	if err != nil {
		panic("Automigrate error")
	}

	f, err := os.Open("../openfoodfacts-products.jsonl")
	check(err)
	scanner := bufio.NewScanner(f)
	scanLine(scanner, db)
	log.Println("Migration completed")
}

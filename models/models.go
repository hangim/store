package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Author struct {
	gorm.Model

	Name string
}

type Book struct {
	gorm.Model

	Title    string
	Synopsis string
	Authors  []Author `gorm:"many2many:book_authors"`
	Price    float64
	Image    string
}

var (
	DB *gorm.DB
)

func init() {
	var err error
	DB, err = gorm.Open("sqlite3", "store.db")
	if err != nil {
		log.Panicln("Failed to connect database")
	}
	DB.AutoMigrate(&Author{})
	DB.AutoMigrate(&Book{})
}

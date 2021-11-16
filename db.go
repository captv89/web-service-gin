package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Book struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Title       string    `json:"title"`
	Authors     string    `json:"authors"`
	Rating      int       `json:"rating"`
	Language    string    `json:"language"`
	Pages       int       `json:"pages"`
	Publisher   string    `json:"publisher"`
	PublishedOn time.Time `json:"publishedon"`
	ISBN        string    `json:"isbn"`
}

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect to database!")
	}

	defer database.Close()
	database.AutoMigrate(&Book{})

	DB = database
}

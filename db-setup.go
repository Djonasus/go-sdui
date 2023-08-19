package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setup() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Todo{})
}

func addRecord(title string, desc string) {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Create(&Todo{Title: title, Description: desc})
}

func updateRecord(newRecord Todo, ch bool) {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Model(&newRecord).Update("checked", ch)
}

func deleteRecord(record Todo) {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Delete(&record)
}

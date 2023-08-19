package main

type Todo struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	Checked     bool
}

package models

import (
	"time"

	"gorm.io/gorm"
)

// Book represents a book in the database
type Book struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Title     string         `gorm:"not null;size:200" json:"title" binding:"required,min=1,max=200"`
	Author    string         `gorm:"not null;size:100" json:"author" binding:"required,min=1,max=100"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for Book
func (Book) TableName() string {
	return "books"
}

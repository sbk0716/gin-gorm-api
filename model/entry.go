package model

import (
	"diary_api/database"

	"gorm.io/gorm"
)

// ============================================================
// Declare Entry model
// ============================================================
type Entry struct {
	// GORM defined a gorm.Model struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	// FYI: https://gorm.io/docs/models.html#gorm-Model
	gorm.Model
	Content string `gorm:"type:text" json:"content"`
	UserID  uint
}

// ============================================================
// Save function
// ============================================================
func (entry *Entry) Save() (*Entry, error) {
	// Inserts value, returning the inserted data's primary key in value's id.
	// FYI: https://gorm.io/docs/create.html#Create-Record
	err := database.Database.Create(&entry).Error
	if err != nil {
		return &Entry{}, err
	}
	return entry, nil
}

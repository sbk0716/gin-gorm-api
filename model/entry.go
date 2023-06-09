package model

import (
	"diary_api/database"
	"fmt"

	"gorm.io/gorm"
)

/*
Entry struct:

1. Content

2. UserID
*/
type Entry struct {
	// GORM defined a gorm.Model struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	// FYI: https://gorm.io/docs/models.html#gorm-Model
	gorm.Model
	Content string `gorm:"type:text" json:"content"`
	UserID  uint
}

/*
Save function:

1. Passes the address of the pointer variable(entry) to (*gorm.DB).Create function.

2. If (*gorm.DB).Create function is successfully executed, it returns the address of the pointer variable(entry) and nil.

FYI: https://gorm.io/docs/create.html

Pointer receiver:

If you want to change the state of the receiver in a method, manipulating the value of it, use a pointer receiver.

FYI: https://go.dev/tour/methods/8
*/
func (entry *Entry) Save() (*Entry, error) {
	// Passes the address of the pointer variable(entry) to (*gorm.DB).Create function.
	// INSERT INTO "entries" ("created_at","updated_at","deleted_at","content","user_id") VALUES ($1$,$2$,$3$,$4$,$5$) RETURNING "id"
	result := database.Database.Create(&entry)
	fmt.Printf("result: %#v\n", result)
	// Returns inserted data's primary key.
	id := entry.ID
	fmt.Printf("id: %#v\n", id)
	// Returns error.
	err := result.Error
	fmt.Printf("err: %#v\n", err)
	// Returns inserted records count.
	rowsAffected := result.RowsAffected
	fmt.Printf("rowsAffected: %#v\n", rowsAffected)
	if err != nil {
		// If (*gorm.DB).Create function fails to execute,
		// it returns the address of empty struct and an error.
		return &Entry{}, err
	}
	// If (*gorm.DB).Create function is successfully executed,
	// it returns the address of the pointer variable(entry) and nil.
	return entry, nil
}

/*
FindEntryById function:

1. Queries the database to find the corresponding user.

2. If (*gorm.DB).Find function is successfully executed, it returns the user struct and nil.
*/
func FindEntryById(id uint) (Entry, error) {
	var entry Entry
	// SELECT * FROM "entries" WHERE ID=$1$ AND "entries"."deleted_at" IS NULL
	err := database.Database.Where("ID=?", id).Find(&entry).Error
	if err != nil {
		// If (*gorm.DB).Find function fails to execute,
		// it Entry the empty struct and an error.
		return Entry{}, err
	}
	// If (*gorm.DB).Find function is successfully executed,
	// it returns the entry struct and nil.
	return entry, nil
}

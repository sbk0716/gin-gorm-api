package model

import (
	"diary_api/database"
	"fmt"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

/*
User struct:

1. Username

2. Password

3. Entries
*/
type User struct {
	// GORM defined a gorm.Model struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	// FYI: https://gorm.io/docs/models.html#gorm-Model
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	// json:"-": This ensures that the user’s password is not returned in the JSON response.
	Password string `gorm:"size:255;not null;" json:"-"`
	// User has many Entries.
	// FYI: https://gorm.io/docs/has_many.html#Has-Many
	Entries []Entry
}

/*
Save function:

1. Passes the address of the pointer variable(user) to (*gorm.DB).Create function.

2. If (*gorm.DB).Create function is successfully executed, it returns the address of the pointer variable(user) and nil.

FYI: https://gorm.io/docs/create.html

Pointer receiver:

If you want to change the state of the receiver in a method, manipulating the value of it, use a pointer receiver.

FYI: https://go.dev/tour/methods/8
*/
func (user *User) Save() (*User, error) {
	// Passes the address of the pointer variable(user) to (*gorm.DB).Create function.
	result := database.Database.Create(&user)
	fmt.Printf("result: %#v\n", result)
	// Returns inserted data's primary key.
	id := user.ID
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
		return &User{}, err
	}
	// If (*gorm.DB).Create function is successfully executed,
	// it returns the address of the pointer variable(user) and nil.
	return user, nil
}

/*
BeforeSave function:

1. Change the state of the receiver.

Object Life Cycle:

FYI: https://gorm.io/docs/hooks.html#Object-Life-Cycle

Pointer receiver:

If you want to change the state of the receiver in a method, manipulating the value of it, use a pointer receiver.

FYI: https://go.dev/tour/methods/8
*/
func (user *User) BeforeSave(*gorm.DB) error {
	// Sets the password.
	pass := []byte(user.Password)
	// Sets the cost.
	cost := bcrypt.DefaultCost
	// Returns the bcrypt hash of the password at the given cost.
	passwordHash, err := bcrypt.GenerateFromPassword(pass, cost)
	if err != nil {
		// If bcrypt.GenerateFromPassword function fails to execute,
		// it returns an error.
		return err
	}
	// Change the state of the receiver.
	(*user).Password = string(passwordHash)
	(*user).Username = html.EscapeString(strings.TrimSpace(user.Username))
	// If bcrypt.GenerateFromPassword function is successfully executed,
	// it returns nil.
	return nil
}

/*
ValidatePassword function:

1. Sets the hash of the user’s password.

2. Sets the hash is generated for the provided plaintext password.

3. The values of hashedPassword and passwordFromString are compared. If they do not match, an error is returned.

Pointer receiver:

If you want to change the state of the receiver in a method, manipulating the value of it, use a pointer receiver.

FYI: https://go.dev/tour/methods/8
*/
func (user *User) ValidatePassword(password string) error {
	// Sets the hash of the user’s password.
	hashedPassword := []byte(user.Password)
	// Sets the hash is generated for the provided plaintext password.
	passwordFromString := []byte(password)
	// The values of hashedPassword and passwordFromString are compared.
	// If they do not match, an error is returned.
	return bcrypt.CompareHashAndPassword(hashedPassword, passwordFromString)
}

/*
FindUserByUsername function:

1. Queries the database to find the corresponding user.

2. If (*gorm.DB).Find function is successfully executed, it returns the user struct and nil.
*/
func FindUserByUsername(username string) (User, error) {
	var user User
	// Queries the database to find the corresponding user.
	err := database.Database.Where("username=?", username).Find(&user).Error
	if err != nil {
		// If (*gorm.DB).Find function fails to execute,
		// it returns the empty struct and an error.
		return User{}, err
	}
	// If (*gorm.DB).Find function is successfully executed,
	// it returns the user struct and nil.
	return user, nil
}

/*
FindUserById function:

1. Queries the database to find the corresponding user.

2. If (*gorm.DB).Find function is successfully executed, it returns the user struct and nil.
*/
func FindUserById(id uint) (User, error) {
	var user User
	err := database.Database.Preload("Entries").Where("ID=?", id).Find(&user).Error
	if err != nil {
		// If (*gorm.DB).Find function fails to execute,
		// it returns the empty struct and an error.
		return User{}, err
	}
	// If (*gorm.DB).Find function is successfully executed,
	// it returns the user struct and nil.
	return user, nil
}

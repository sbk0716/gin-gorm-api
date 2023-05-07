package model

import (
	"diary_api/database"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// ============================================================
// Declare User model
// ============================================================
type User struct {
	// GORM defined a gorm.Model struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	// FYI: https://gorm.io/docs/models.html#gorm-Model
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	// Notice that the JSON binding for the Password field is -. This ensures that the user’s password is not returned in the JSON response.
	Password string `gorm:"size:255;not null;" json:"-"`
	// User has many Entries.
	// FYI: https://gorm.io/docs/has_many.html#Has-Many
	Entries []Entry
}

// ============================================================
// Save function
// ============================================================
func (user *User) Save() (*User, error) {
	// Inserts value, returning the inserted data's primary key in value's id.
	// FYI: https://gorm.io/docs/create.html#Create-Record
	err := database.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

// ============================================================
// BeforeSave function
// FYI: https://gorm.io/docs/hooks.html#Object-Life-Cycle
// ============================================================
func (user *User) BeforeSave(*gorm.DB) error {
	pass := []byte(user.Password)
	// the cost that will actually be set if a cost below MinCost is passed into GenerateFromPassword
	cost := bcrypt.DefaultCost
	// Returns the bcrypt hash of the password at the given cost.
	passwordHash, err := bcrypt.GenerateFromPassword(pass, cost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

// ============================================================
// ValidatePassword function
// ============================================================
func (user *User) ValidatePassword(password string) error {
	// A: hash is generated for the provided plaintext password.
	// B: hash of the user’s password.
	// The values of A and B are compared. If they do not match, an error is returned.
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

// ============================================================
// FindUserByUsername function
// ============================================================
func FindUserByUsername(username string) (User, error) {
	var user User
	// FindUserByUsername takes a username and queries the database to find the corresponding user.
	err := database.Database.Where("username=?", username).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// ============================================================
// FindUserById function
// ============================================================
func FindUserById(id uint) (User, error) {
	var user User
	err := database.Database.Preload("Entries").Where("ID=?", id).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

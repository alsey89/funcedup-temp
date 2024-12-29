package schema

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordHash"`

	//all posts by this user
	Posts []Post `json:"posts" gorm:"foreignKey:UserID"`
}

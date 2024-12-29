package schema

import "gorm.io/gorm"

type Post struct {
	gorm.Model

	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID uint   `json:"userId"`
	Type   string `json:"type"` //forumPost, note, etc.

	// optional
	User *User `json:"user" gorm:"foreignKey:UserID"`
}

type User struct {
	gorm.Model

	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordHash"`
}

type PostTag struct {
	gorm.Model
	PostID uint `json:"postId"`
	TagID  uint `json:"tagId"`
}

type Tags struct {
	gorm.Model

	Name string `json:"name"`
}

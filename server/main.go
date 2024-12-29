package main

import (
	"fmt"
	"funcedup/internal/schema"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()

	// Connect to the database
	var dsn string
	dsn = "host=localhost user=postgres password=postgres dbname=postgres port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		e.Logger.Fatal(err)
	}
	if db != nil {
		fmt.Printf("Database connected")
	}

	// Migration
	err = db.AutoMigrate(
		&schema.Post{},
		&schema.User{},
		&schema.PostTag{},
		&schema.Tags{},
	)
	if err != nil {
		e.Logger.Fatal(err)
	}

	// // create test data in db
	// user := schema.User{
	// 	Name:         "John Doe",
	// 	Email:        "johndoe123@mailnesia.com",
	// 	PasswordHash: "password",
	// }
	// db.Create(&user)

	// // create post
	// post := schema.Post{
	// 	Title:  "Hello, World!",
	// 	Body:   "This is a test post",
	// 	UserID: uint(1),
	// 	Type:   "forumPost",
	// }
	// db.Create(&post)

	// query user and post
	var queriedUser schema.User
	db.
		Where("email = ?", "johndoe123@mailnesia.com").
		First(&queriedUser)

	var queriedPost schema.Post
	db.
		Where("title = ?", "Hello, World!").
		First(&queriedPost)

	fmt.Printf("Queried user: %v\n", queriedUser)
	fmt.Printf("Queried post: %v\n", queriedPost)

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/posts", func(c echo.Context) error {
		return c.String(http.StatusNotFound, "Here are your posts")
	})

	e.Logger.Fatal(e.Start(":1323"))
}

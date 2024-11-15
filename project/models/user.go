package models

import (
	"log/slog"

	"github.com/google/uuid"
	"github.com/peterszarvas94/goat/database"
	"github.com/peterszarvas94/goat/log"
	"gorm.io/gorm"
)

type User struct {
	ID    string `gorm:"primaryKey"`
	Name  string
	Email string
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}

func Seed() error {

	conn, err := database.Get()
	if err != nil {
		return err
	}

	err = conn.DB.Migrator().DropTable(&User{})
	if err != nil {
		return err
	}

	err = conn.DB.AutoMigrate(&User{})
	if err != nil {
		return err
	}

	newUser := User{Name: "John Doe", Email: "john@example.com"}
	conn.DB.Create(&newUser)

	var users []User
	conn.DB.Find(&users)

	for _, user := range users {
		log.Logger.Info(
			"User created with seed",
			slog.String("ID", user.ID),
			slog.String("name", user.Name),
			slog.String("email", user.Email),
		)
	}

	return nil
}

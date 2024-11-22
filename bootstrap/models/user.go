package models

import (
	"log/slog"

	"github.com/peterszarvas94/goat/database"
	l "github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID    string `gorm:"primaryKey"`
	Name  string
	Email string
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New("usr")
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

	newUsers := [3]User{
		{Name: "John Doe", Email: "john@example.com"},
		{Name: "John Deer", Email: "jdeer@example.com"},
		{Name: "George Hey", Email: "georg@example.com"},
	}

	conn.DB.Create(&newUsers)

	var users []User
	conn.DB.Find(&users)

	for _, user := range users {
		l.Logger.Debug(
			"User created with seed",
			slog.String("ID", user.ID),
			slog.String("name", user.Name),
			slog.String("email", user.Email),
		)
	}

	return nil
}

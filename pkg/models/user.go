package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	Email          string    `gorm:"uniqueIndex" json:"email"`
	HashedPassword string    `json:"-"`
	IsActive       bool      `json:"isActive"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
} //@name User

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

package models

import (
	"crypto/rand"
	"fmt"
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

func GeneratePassword(length int) string {
	if length == 0 || length < 8 {
		length = 8
	}
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	password := fmt.Sprintf("%X", bytes)

	return password
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

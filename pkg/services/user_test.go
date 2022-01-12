package services_test

import (
	"os"
	"testing"

	"github.com/chenjr0719/golang-boilerplate/pkg/db"
	"github.com/chenjr0719/golang-boilerplate/pkg/models"
	"github.com/chenjr0719/golang-boilerplate/pkg/services"
	"github.com/stretchr/testify/assert"
)

func TestUserService(t *testing.T) {

	databaseURI := "sqlite://file::memory:?cache=shared"
	os.Setenv("DATABASE_URI", databaseURI)
	dbConnection, _ := db.ConnectDatabase()
	userService := services.NewUserService(dbConnection)

	assert.NoError(t, nil)

	users, err := userService.List()
	assert.NoError(t, err)
	assert.Equal(t, len(users), 0)

	user := models.User{
		Name:  "test",
		Email: "test@example.com",
	}

	userInDB, err := userService.Create(user)
	assert.NoError(t, err)
	assert.Equal(t, userInDB.Name, user.Name)
	assert.Equal(t, userInDB.Email, user.Email)
	assert.NotEmpty(t, userInDB.CreatedAt)
}

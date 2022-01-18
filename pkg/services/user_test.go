package services_test

import (
	"testing"

	"github.com/chenjr0719/golang-boilerplate/pkg/models"
	"github.com/chenjr0719/golang-boilerplate/pkg/services"
	"github.com/stretchr/testify/assert"
)

func TestUserService(t *testing.T) {

	userService := services.NewUserService()

	users, err := userService.List()
	assert.NoError(t, err)
	assert.Equal(t, len(users), 0)

	user := models.User{
		Name:  "user-test",
		Email: "user-test@example.com",
	}

	userInDB, err := userService.Create(user)
	assert.NoError(t, err)
	assert.Equal(t, user.Name, userInDB.Name)
	assert.Equal(t, user.Email, userInDB.Email)
	assert.NotEmpty(t, userInDB.CreatedAt)
	assert.NotEmpty(t, userInDB.UpdatedAt)
	assert.Equal(t, userInDB.CreatedAt, userInDB.UpdatedAt)

	usersInDB, err := userService.List()
	assert.NoError(t, err)
	assert.Equal(t, len(usersInDB), 1)

	userInDB = usersInDB[0]
	assert.Equal(t, user.Name, userInDB.Name)
	assert.Equal(t, user.Email, userInDB.Email)

	userInDB, err = userService.Get("1")
	assert.NoError(t, err)
	assert.Equal(t, user.Name, userInDB.Name)
	assert.Equal(t, userInDB.Email, user.Email)

	updateUser := models.User{
		Name:  "test-updated",
		Email: "test-updated@example.com",
	}
	userInDB, err = userService.Update("1", updateUser)
	assert.NoError(t, err)
	assert.Equal(t, updateUser.Name, userInDB.Name)
	assert.Equal(t, updateUser.Email, userInDB.Email)
	assert.NotEqual(t, userInDB.CreatedAt, userInDB.UpdatedAt)

	err = userService.Delete("1")
	assert.NoError(t, err)

	userInDB, err = userService.Get("1")
	assert.Error(t, err)

}

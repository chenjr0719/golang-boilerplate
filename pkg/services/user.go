package services

import (
	"context"
	"time"

	"github.com/chenjr0719/golang-boilerplate/pkg/db"
	"github.com/chenjr0719/golang-boilerplate/pkg/log"
	"github.com/chenjr0719/golang-boilerplate/pkg/models"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService() *UserService {
	log.Info().Msg("Initialize UserService")

	userService := &UserService{db: db.DB}
	db.DB.AutoMigrate(&models.User{})

	return userService
}

func (service *UserService) newSession() *gorm.DB {
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	session := service.db.WithContext(ctx)

	return session
}

func (service *UserService) List() ([]models.User, error) {
	var users []models.User
	session := service.newSession()
	result := session.Find(&users)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("List users failed")
		return nil, result.Error
	}
	return users, nil
}

func (service *UserService) Create(user models.User) (models.User, error) {
	session := service.newSession()
	result := session.Create(&user)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Create user failed")
		return models.User{}, result.Error
	}
	session.First(&user, user.ID)
	return user, nil
}

func (service *UserService) Get(id string) (models.User, error) {
	var user models.User
	session := service.newSession()
	result := session.First(&user, id)
	if result.Error != nil {
		log.Error().Err(result.Error).Msgf("Get user %s failed", id)
		return models.User{}, result.Error
	}
	return user, nil
}

func (service *UserService) Update(id string, user models.User) (models.User, error) {
	userInDB, err := service.Get(id)
	if err != nil {
		log.Error().Err(err).Msgf("Update user %s failed", id)
		return models.User{}, err
	}
	session := service.newSession()
	result := session.Model(&userInDB).Updates(user)
	if result.Error != nil {
		log.Error().Err(result.Error).Msgf("Update user %s failed", id)
		return models.User{}, result.Error
	}
	return userInDB, nil
}

func (service *UserService) Delete(id string) error {
	userInDB, err := service.Get(id)
	if err != nil {
		log.Error().Err(err).Msgf("Delete user %s failed", id)
		return err
	}
	session := service.newSession()
	result := session.Delete(&userInDB)
	if result.Error != nil {
		log.Error().Err(result.Error).Msgf("Delete user %s failed", id)
		return result.Error
	}
	return nil
}

package services

import (
	"context"
	"time"

	"github.com/chenjr0719/golang-boilerplate/pkg/db"
	"github.com/chenjr0719/golang-boilerplate/pkg/log"
	"github.com/chenjr0719/golang-boilerplate/pkg/models"
	"gorm.io/gorm"
)

type JobService struct {
	db *gorm.DB
}

func NewJobService() *JobService {
	log.Info().Msg("Initialize JobService")

	jobService := &JobService{db: db.DB.Preload("User")}
	db.DB.Exec(models.CREATE_ENUM)
	db.DB.AutoMigrate(&models.Job{})

	return jobService
}

func (service *JobService) newSession() *gorm.DB {
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	session := service.db.WithContext(ctx)

	return session
}

func (service *JobService) List() ([]models.Job, error) {
	var jobs []models.Job
	session := service.newSession()
	result := session.Find(&jobs)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("List jobs failed")
		return nil, result.Error
	}
	return jobs, nil
}

func (service *JobService) Create(job models.Job) (models.Job, error) {
	session := service.newSession()
	result := session.Create(&job)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Create job failed")
		return models.Job{}, result.Error
	}
	session.First(&job, job.ID)
	return job, nil
}

func (service *JobService) Get(id string) (models.Job, error) {
	var job models.Job
	session := service.newSession()
	result := session.First(&job, id)
	if result.Error != nil {
		log.Error().Err(result.Error).Msgf("Get job %s failed", id)
		return models.Job{}, result.Error
	}
	return job, nil
}

func (service *JobService) Update(id string, job models.Job) (models.Job, error) {
	jobInDB, err := service.Get(id)
	if err != nil {
		log.Error().Err(err).Msgf("Update job %s failed", id)
		return models.Job{}, err
	}
	session := service.newSession()
	result := session.Model(&jobInDB).Updates(job)
	if result.Error != nil {
		log.Error().Err(result.Error).Msgf("Update job %s failed", id)
		return models.Job{}, result.Error
	}
	return jobInDB, nil
}

func (service *JobService) Delete(id string) error {
	jobInDB, err := service.Get(id)
	if err != nil {
		log.Error().Err(err).Msgf("Delete job %s failed", id)
		return err
	}
	session := service.newSession()
	result := session.Delete(&jobInDB)
	if result.Error != nil {
		log.Error().Err(result.Error).Msgf("Delete job %s failed", id)
		return result.Error
	}
	return nil
}

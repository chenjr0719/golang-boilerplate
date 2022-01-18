package services_test

import (
	"testing"

	"github.com/chenjr0719/golang-boilerplate/pkg/models"
	"github.com/chenjr0719/golang-boilerplate/pkg/services"
	"github.com/stretchr/testify/assert"
)

func TestJobService(t *testing.T) {

	userService := services.NewUserService()
	jobService := services.NewJobService()

	user := models.User{
		Name:  "job-test",
		Email: "job-test@example.com",
	}

	user, err := userService.Create(user)
	assert.NoError(t, err)

	job := models.Job{
		Name:   "sleep",
		Args:   []byte(`{"seconds": 3}`),
		UserID: user.ID,
	}
	jobInDB, err := jobService.Create(job)
	assert.NoError(t, err)
	assert.Equal(t, job.Name, jobInDB.Name)
	assert.Equal(t, job.Args, jobInDB.Args)
	assert.Equal(t, user.ID, jobInDB.UserID)
	assert.NotEmpty(t, jobInDB.CreatedAt)
	assert.NotEmpty(t, jobInDB.UpdatedAt)
	assert.Equal(t, jobInDB.CreatedAt, jobInDB.UpdatedAt)

	jobsInDB, err := jobService.List()
	assert.NoError(t, err)
	assert.Equal(t, len(jobsInDB), 1)

	jobInDB = jobsInDB[0]
	assert.Equal(t, job.Name, jobInDB.Name)
	assert.Equal(t, job.Args, jobInDB.Args)
	assert.Equal(t, user.ID, jobInDB.UserID)

	jobInDB, err = jobService.Get("1")
	assert.NoError(t, err)
	assert.Equal(t, job.Name, jobInDB.Name)
	assert.Equal(t, job.Args, jobInDB.Args)
	assert.Equal(t, user.ID, jobInDB.UserID)

	updateJob := models.Job{
		Args: []byte(`{"seconds": 5}`),
	}
	jobInDB, err = jobService.Update("1", updateJob)
	assert.NoError(t, err)
	assert.Equal(t, job.Name, jobInDB.Name)
	assert.Equal(t, updateJob.Args, jobInDB.Args)
	assert.NotEqual(t, jobInDB.CreatedAt, jobInDB.UpdatedAt)

	err = jobService.Delete("1")
	assert.NoError(t, err)

	jobInDB, err = jobService.Get("1")
	assert.Error(t, err)

	// Clean up
	userService.Delete("1")

}

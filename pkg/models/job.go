package models

import (
	"database/sql/driver"
	"time"

	"gorm.io/datatypes"
)

type JobStatus string

const (
	Pending     JobStatus = "Pending"
	Running     JobStatus = "Running"
	Finished    JobStatus = "Finished"
	Failed      JobStatus = "Failed"
	CREATE_ENUM string    = `DO $$
	BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'job_status') THEN
			CREATE TYPE job_status AS ENUM('Pending', 'Running', 'Finished', 'Failed');
		END IF;
	END$$;`
)

func (s *JobStatus) Scan(value interface{}) error {
	*s = JobStatus(value.(string))
	return nil
}

func (s JobStatus) Value() (driver.Value, error) {
	return string(s), nil
}

type Job struct {
	ID        uint           `json:"id"`
	Status    JobStatus      `json:"status" gorm:"type:job_status"`
	Name      string         `json:"name" binding:"required"`
	Args      datatypes.JSON `json:"args" binding:"required"`
	Result    datatypes.JSON `json:"result"`
	UserID    uint           `json:"userID" binding:"required"`
	User      *User          `json:"user"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
} //@name Job

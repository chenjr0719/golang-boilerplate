package tasks

import (
	"fmt"
)

func Error(jobID string, message string) error {
	return fmt.Errorf("error: %s", message)
}

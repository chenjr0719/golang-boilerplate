package tasks_test

import (
	"testing"

	"github.com/chenjr0719/golang-boilerplate/pkg/worker/tasks"
	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {

	jobID := "1"
	message := "MockError"
	result := tasks.Error(jobID, message)

	assert.Error(t, result)
}

package worker

import (
	"fmt"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/chenjr0719/golang-boilerplate/pkg/models"
)

type WorkerClient struct {
	Server *machinery.Server
}

func NewWorkerClient() *WorkerClient {

	server, err := startServer()
	if err != nil {
		panic(err)
	}
	workerClient := &WorkerClient{
		Server: server,
	}

	return workerClient
}

func (wc *WorkerClient) newTask(job models.Job) (tasks.Signature, error) {
	taskArgs := []tasks.Arg{{
		Name:  "jobID",
		Type:  "string",
		Value: fmt.Sprintf("%v", job.ID),
	}, {
		Name:  "input",
		Type:  "string",
		Value: string(job.Args),
	}}

	task := tasks.Signature{
		Name:                        job.Name,
		Args:                        taskArgs,
		IgnoreWhenTaskNotRegistered: true,
	}

	return task, nil
}

func (wc *WorkerClient) SendTask(job models.Job) error {
	task, err := wc.newTask(job)
	if err != nil {
		return fmt.Errorf("send task failed, cannot parse task args: %s", err.Error())
	}

	_, err = wc.Server.SendTask(&task)
	if err != nil {
		return fmt.Errorf("send task failed, cannot send task to worker: %s", err.Error())
	}

	return err
}

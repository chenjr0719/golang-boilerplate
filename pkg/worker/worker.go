package worker

import (
	"fmt"
	"time"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/backends/result"

	machineryconfig "github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/log"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/chenjr0719/golang-boilerplate/pkg/db"
	"github.com/chenjr0719/golang-boilerplate/pkg/models"
	"github.com/chenjr0719/golang-boilerplate/pkg/services"
	workertasks "github.com/chenjr0719/golang-boilerplate/pkg/worker/tasks"
)

func startServer() (*machinery.Server, error) {

	cnf, err := machineryconfig.NewFromEnvironment()
	if err != nil {
		return nil, err
	}

	server, err := machinery.NewServer(cnf)
	if err != nil {
		return nil, err
	}

	// Register tasks
	tasks := map[string]interface{}{
		"fibonacci":      workertasks.Fibonacci,
		"sleep":          workertasks.Sleep,
		"error":          workertasks.Error,
		"remoteHTTPCall": workertasks.RemoteHTTPCall,
	}

	return server, server.RegisterTasks(tasks)
}

func NewWorker() error {
	serviceName := "worker"

	server, err := startServer()
	if err != nil {
		return err
	}

	// The second argument is a consumer tag
	// Ideally, each worker should have a unique tag (worker1, worker2 etc)
	worker := server.NewWorker(serviceName, 0)

	// Here we inject some custom code for error handling,
	// start and end of task hooks, useful for metrics for example.
	errorhandler := func(err error) {
		log.ERROR.Println("Error:", err)
	}

	worker.SetPostTaskHandler(PostTaskHandler)
	worker.SetErrorHandler(errorhandler)
	worker.SetPreTaskHandler(PreTaskHandler)

	return worker.Launch()
}

func PreTaskHandler(signature *tasks.Signature) {
	log.INFO.Println("Started:", signature.Name)
	log.INFO.Println("UUID:", signature.UUID)
	log.INFO.Println("Args:", signature.Args)

	dbConnection, err := db.ConnectDatabase()
	if err != nil {
		log.FATAL.Println(err)
	}

	jobService := services.NewJobService(dbConnection)
	jobID := fmt.Sprintf("%s", signature.Args[0].Value)
	job, err := jobService.Get(jobID)
	if err != nil {
		log.FATAL.Println(err)
	}
	job.Status = models.Running
	_, err = jobService.Update(jobID, job)
	if err != nil {
		log.FATAL.Println(err)
	}
}

func PostTaskHandler(signature *tasks.Signature) {
	log.INFO.Println("Finished:", signature.Name)
	log.INFO.Println("UUID:", signature.UUID)

	dbConnection, err := db.ConnectDatabase()
	if err != nil {
		log.FATAL.Println(err)
	}

	jobService := services.NewJobService(dbConnection)
	jobID := fmt.Sprintf("%s", signature.Args[0].Value)
	job, err := jobService.Get(jobID)
	if err != nil {
		log.FATAL.Println(err)
	}

	server, _ := startServer()
	asyncResult := result.NewAsyncResult(signature, server.GetBackend())
	result, err := asyncResult.Get(time.Duration(time.Millisecond * 5))

	var jobStatus models.JobStatus
	var jobResult []byte
	if err != nil {
		jobStatus = models.Failed
		jobResult = []byte(fmt.Sprintf(`{"error": "%s"}`, err.Error()))
	} else {
		jobStatus = models.Finished
		jobResult = []byte(fmt.Sprintf("%s", result[0]))
	}

	job.Status = jobStatus
	job.Result = jobResult
	_, err = jobService.Update(jobID, job)
	if err != nil {
		log.FATAL.Println(err)
	}
}

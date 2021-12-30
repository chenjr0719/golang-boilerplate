package v1

import (
	"github.com/chenjr0719/golang-boilerplate/pkg/apiserver/error"
	"github.com/chenjr0719/golang-boilerplate/pkg/models"
	"github.com/chenjr0719/golang-boilerplate/pkg/services"
	"github.com/chenjr0719/golang-boilerplate/pkg/worker"
	"github.com/gin-gonic/gin"
)

type V1JobAPI struct {
	JobService   services.JobService
	WorkerClient worker.WorkerClient
}

type JobInput struct {
	Name   string   `json:"name" binding:"required"`
	Args   struct{} `json:"args"  binding:"required"`
	UserID uint     `json:"userID" binding:"required"`
} //@name JobInput

var v1JobAPI *V1JobAPI

func NewV1JobGroup(apiGroup *gin.RouterGroup) *gin.RouterGroup {
	v1JobAPI = &V1JobAPI{
		JobService:   *services.NewJobService(),
		WorkerClient: *worker.NewWorkerClient(),
	}

	v1JobGroup := apiGroup.Group("/jobs")
	v1JobGroup.GET("/", v1JobAPI.ListJobs)
	v1JobGroup.POST("/", v1JobAPI.CreateJob)
	v1JobGroup.GET("/:id", v1JobAPI.GetJob)
	v1JobGroup.PUT("/:id", v1JobAPI.UpdateJob)
	v1JobGroup.DELETE("/:id", v1JobAPI.DeleteJob)
	v1JobGroup.POST("/:id/run", v1JobAPI.RunJob)

	return v1JobGroup
}

// ListJob godoc
// @Summary List Job
// @Schemes
// @Tags Jobs
// @Accept json
// @Produce json
// @Success 200 {array} models.Job
// @Failure 500 {object} error.HTTPError
// @Router /v1/jobs [get]
func (api *V1JobAPI) ListJobs(ctx *gin.Context) {
	jobs, err := api.JobService.List()
	if err != nil {
		error.NewError(ctx, 500, err)
		return
	}

	ctx.JSON(200, gin.H{
		"items": jobs,
	})
}

// CreateJob godoc
// @Summary Create Job
// @Schemes
// @Tags Jobs
// @Accept json
// @Produce json
// @Param job body JobInput true "Job Input"
// @Success 200 {object} models.Job
// @Failure 400 {object} error.HTTPError
// @Failure 500 {object} error.HTTPError
// @Router /v1/jobs [post]
func (api *V1JobAPI) CreateJob(ctx *gin.Context) {
	var job models.Job
	if err := ctx.ShouldBindJSON(&job); err != nil {
		error.NewError(ctx, 400, err)
		return
	}

	job.Status = models.Pending
	job, err := api.JobService.Create(job)
	if err != nil {
		error.NewError(ctx, 400, err)
		return
	}

	err = api.WorkerClient.SendTask(job)
	// err = api.WorkerClient.SendTask(ctx.Request.Context(), job)
	if err != nil {
		error.NewError(ctx, 500, err)
	}

	ctx.JSON(200, job)
}

// GetJob godoc
// @Summary Get Job
// @Schemes
// @Tags Jobs
// @Accept json
// @Produce json
// @Param id path int true "Job ID"
// @Success 200 {object} models.Job
// @Failure 404 {object} error.HTTPError
// @Failure 500 {object} error.HTTPError
// @Router /v1/jobs/{id} [get]
func (api *V1JobAPI) GetJob(ctx *gin.Context) {
	id := ctx.Param("id")
	job, err := api.JobService.Get(id)
	if err != nil {
		error.NewError(ctx, 404, err)
		return
	}

	ctx.JSON(200, job)
}

// UpdateJob godoc
// @Summary Update Job
// @Schemes
// @Tags Jobs
// @Accept json
// @Produce json
// @Param id path int true "Job ID"
// @Param job body JobInput true "Job Input"
// @Success 200 {object} models.Job
// @Failure 400 {object} error.HTTPError
// @Failure 404 {object} error.HTTPError
// @Failure 500 {object} error.HTTPError
// @Router /v1/jobs/{id} [put]
func (api *V1JobAPI) UpdateJob(ctx *gin.Context) {
	id := ctx.Param("id")
	var job models.Job
	if err := ctx.ShouldBindJSON(&job); err != nil {
		error.NewError(ctx, 400, err)
		return
	}

	job, err := api.JobService.Update(id, job)
	if err != nil {
		error.NewError(ctx, 500, err)
		return
	}

	ctx.JSON(200, job)
}

// DeleteJob godoc
// @Summary Delete Job
// @Schemes
// @Tags Jobs
// @Accept json
// @Produce json
// @Param id path int true "Job ID"
// @Success 204 {object} nil
// @Failure 404 {object} error.HTTPError
// @Failure 500 {object} error.HTTPError
// @Router /v1/jobs/{id} [delete]
func (api *V1JobAPI) DeleteJob(ctx *gin.Context) {
	id := ctx.Param("id")

	err := api.JobService.Delete(id)
	if err != nil {
		error.NewError(ctx, 400, err)
		return
	}

	ctx.JSON(204, nil)
}

// RunJob godoc
// @Summary Run Job
// @Schemes
// @Tags Jobs
// @Accept json
// @Produce json
// @Param id path int true "Job ID"
// @Success 200 {object} models.Job
// @Failure 404 {object} error.HTTPError
// @Failure 500 {object} error.HTTPError
// @Router /v1/jobs/{id}/run [post]
func (api *V1JobAPI) RunJob(ctx *gin.Context) {
	id := ctx.Param("id")
	job, err := api.JobService.Get(id)
	if err != nil {
		error.NewError(ctx, 404, err)
		return
	}

	job.Status = models.Pending
	job.Result = []byte("{}")
	job, err = api.JobService.Update(id, job)
	if err != nil {
		error.NewError(ctx, 500, err)
		return
	}

	err = api.WorkerClient.SendTask(job)
	// err = api.WorkerClient.SendTask(ctx.Request.Context(), job)
	if err != nil {
		error.NewError(ctx, 500, err)
	}

	ctx.JSON(200, job)
}

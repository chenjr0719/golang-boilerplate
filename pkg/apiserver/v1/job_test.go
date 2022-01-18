package v1_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	v1 "github.com/chenjr0719/golang-boilerplate/pkg/apiserver/v1"
	"github.com/chenjr0719/golang-boilerplate/pkg/models"
	"github.com/chenjr0719/golang-boilerplate/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fastjson"
)

func setupBroker() {
	os.Setenv("BROKER", "amqp://rabbitmq:rabbitmq-password@localhost:5672/boilerplate")
	os.Setenv("RESULT_BACKEND", "amqp://rabbitmq:rabbitmq-password@localhost:5672/boilerplate")
	os.Setenv("DEFAULT_QUEUE", "boilerplate")
	os.Setenv("RESULTS_EXPIRE_IN", "3600")
	os.Setenv("AMQP_EXCHANGE", "boilerplate")
	os.Setenv("AMQP_EXCHANGE_TYPE", "direct")
	os.Setenv("AMQP_BINDING_KEY", "boilerplate")
	os.Setenv("AMQP_PREFETCH_COUNT", "1")
}

func setupTestV1JobAPIRouter() *gin.Engine {
	setupBroker()

	router := gin.New()
	group := router.Group("/")

	v1.NewV1JobGroup(group)

	return router
}

func setupTestUser() models.User {
	userService := services.NewUserService()

	user := models.User{
		Name:  "job-test",
		Email: "job-test@example.com",
	}

	user, _ = userService.Create(user)
	defer userService.Delete(fmt.Sprint(user.ID))

	return user
}

func TestV1JobAPIList(t *testing.T) {
	router := setupTestV1JobAPIRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/jobs/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var jsonParser fastjson.Parser
	jsonFields, err := jsonParser.Parse(w.Body.String())
	assert.Nil(t, err)
	assert.True(t, jsonFields.Exists("items"))
	items := jsonFields.GetArray("items")
	assert.Equal(t, 0, len(items))

}

func TestV1JobAPICreate(t *testing.T) {

	// Setup
	user := setupTestUser()

	//  Test
	router := setupTestV1JobAPIRouter()
	job := models.Job{
		Name:   "sleep",
		Args:   []byte(`{"seconds":3}`),
		UserID: user.ID,
	}

	w := httptest.NewRecorder()
	jsonBody, _ := json.Marshal(job)
	requestBody := bytes.NewBuffer(jsonBody)
	req, _ := http.NewRequest("POST", "/jobs/", requestBody)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var jsonParser fastjson.Parser
	jsonFields, err := jsonParser.Parse(w.Body.String())
	assert.Nil(t, err)
	assert.True(t, jsonFields.Exists("name"))
	assert.True(t, jsonFields.Exists("args"))
	assert.True(t, jsonFields.Exists("status"))

	name := string(jsonFields.GetStringBytes("name"))
	assert.Equal(t, job.Name, name)
	args := string(jsonFields.Get("args").MarshalTo(nil))
	assert.Equal(t, string(job.Args), args)
	status := string(jsonFields.GetStringBytes("status"))
	assert.Equal(t, "Pending", status)

}

func TestV1JobAPIGet(t *testing.T) {

	// Setup
	user := setupTestUser()

	jobService := services.NewJobService()
	job := models.Job{
		Name:   "sleep",
		Args:   []byte(`{"seconds":3}`),
		UserID: user.ID,
	}
	job, _ = jobService.Create(job)

	// Test
	router := setupTestV1JobAPIRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/jobs/%v", job.ID), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var jsonParser fastjson.Parser
	jsonFields, err := jsonParser.Parse(w.Body.String())
	assert.Nil(t, err)
	assert.True(t, jsonFields.Exists("name"))
	assert.True(t, jsonFields.Exists("args"))

	name := string(jsonFields.GetStringBytes("name"))
	assert.Equal(t, job.Name, name)
	args := string(jsonFields.Get("args").MarshalTo(nil))
	assert.Equal(t, string(job.Args), args)
}

func TestV1JobAPIUpdate(t *testing.T) {

	// Setup
	user := setupTestUser()

	jobService := services.NewJobService()
	job := models.Job{
		Name:   "sleep",
		Args:   []byte(`{"seconds":3}`),
		UserID: user.ID,
	}
	job, _ = jobService.Create(job)

	// Test
	router := setupTestV1JobAPIRouter()

	updateJob := models.Job{
		Name:   "sleep",
		Args:   []byte(`{"seconds":5}`),
		UserID: user.ID,
	}

	w := httptest.NewRecorder()
	jsonBody, _ := json.Marshal(updateJob)
	requestBody := bytes.NewBuffer(jsonBody)
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/jobs/%v", job.ID), requestBody)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var jsonParser fastjson.Parser
	jsonFields, err := jsonParser.Parse(w.Body.String())
	assert.Nil(t, err)
	assert.True(t, jsonFields.Exists("name"))
	assert.True(t, jsonFields.Exists("args"))
	assert.True(t, jsonFields.Exists("status"))

	name := string(jsonFields.GetStringBytes("name"))
	assert.Equal(t, updateJob.Name, name)
	args := string(jsonFields.Get("args").MarshalTo(nil))
	assert.Equal(t, string(updateJob.Args), args)

}

func TestV1JobAPIDelete(t *testing.T) {

	// Setup
	user := setupTestUser()

	jobService := services.NewJobService()
	job := models.Job{
		Name:   "sleep",
		Args:   []byte(`{"seconds":3}`),
		UserID: user.ID,
	}
	job, _ = jobService.Create(job)

	// Test
	router := setupTestV1JobAPIRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/jobs/%v", job.ID), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 204, w.Code)
	assert.Empty(t, w.Body.String())

}

func TestV1JobAPIRunJob(t *testing.T) {

	// Setup
	user := setupTestUser()

	jobService := services.NewJobService()
	job := models.Job{
		Name:   "sleep",
		Args:   []byte(`{"seconds":3}`),
		UserID: user.ID,
	}
	job, _ = jobService.Create(job)

	// Test
	router := setupTestV1JobAPIRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", fmt.Sprintf("/jobs/%v/run", job.ID), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var jsonParser fastjson.Parser
	jsonFields, err := jsonParser.Parse(w.Body.String())
	assert.Nil(t, err)
	assert.True(t, jsonFields.Exists("name"))
	assert.True(t, jsonFields.Exists("args"))
	assert.True(t, jsonFields.Exists("status"))

	name := string(jsonFields.GetStringBytes("name"))
	assert.Equal(t, job.Name, name)
	args := string(jsonFields.Get("args").MarshalTo(nil))
	assert.Equal(t, string(job.Args), args)
	status := string(jsonFields.GetStringBytes("status"))
	assert.Equal(t, "Pending", status)

}

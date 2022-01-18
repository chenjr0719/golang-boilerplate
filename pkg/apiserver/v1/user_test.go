package v1_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	v1 "github.com/chenjr0719/golang-boilerplate/pkg/apiserver/v1"
	"github.com/chenjr0719/golang-boilerplate/pkg/models"
	"github.com/chenjr0719/golang-boilerplate/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fastjson"
)

func setupTestV1UserAPIRouter() *gin.Engine {
	router := gin.New()
	group := router.Group("/")

	v1.NewV1UserGroup(group)

	return router
}

func TestV1UserAPIList(t *testing.T) {
	router := setupTestV1UserAPIRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var jsonParser fastjson.Parser
	jsonFields, err := jsonParser.Parse(w.Body.String())
	assert.Nil(t, err)
	assert.True(t, jsonFields.Exists("items"))
	items := jsonFields.GetArray("items")
	assert.Equal(t, 0, len(items))

}

func TestV1UserAPICreate(t *testing.T) {
	router := setupTestV1UserAPIRouter()
	user := models.User{
		Name:  "user-api-test-create",
		Email: "user-api-test-create@example.com",
	}

	w := httptest.NewRecorder()
	jsonBody, _ := json.Marshal(user)
	requestBody := bytes.NewBuffer(jsonBody)
	req, _ := http.NewRequest("POST", "/users/", requestBody)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var jsonParser fastjson.Parser
	jsonFields, err := jsonParser.Parse(w.Body.String())
	assert.Nil(t, err)
	assert.True(t, jsonFields.Exists("name"))
	assert.True(t, jsonFields.Exists("email"))

	name := string(jsonFields.GetStringBytes("name"))
	assert.Equal(t, user.Name, name)
	email := string(jsonFields.GetStringBytes("email"))
	assert.Equal(t, user.Email, email)

}

func TestV1UserAPIGet(t *testing.T) {
	router := setupTestV1UserAPIRouter()

	userService := services.NewUserService()
	user := models.User{
		Name:  "user-api-test-get",
		Email: "user-api-test-get@example.com",
	}
	user, _ = userService.Create(user)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/users/%v", user.ID), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var jsonParser fastjson.Parser
	jsonFields, err := jsonParser.Parse(w.Body.String())
	assert.Nil(t, err)
	assert.True(t, jsonFields.Exists("name"))
	assert.True(t, jsonFields.Exists("email"))

	name := string(jsonFields.GetStringBytes("name"))
	assert.Equal(t, user.Name, name)
	email := string(jsonFields.GetStringBytes("email"))
	assert.Equal(t, user.Email, email)
}

func TestV1UserAPIUpdate(t *testing.T) {
	router := setupTestV1UserAPIRouter()

	userService := services.NewUserService()
	user := models.User{
		Name:  "user-api-test-update",
		Email: "user-api-test-update@example.com",
	}
	user, _ = userService.Create(user)

	updateUser := models.User{
		Name:  "user-api-test-updated",
		Email: "user-api-test-updated@example.com",
	}

	w := httptest.NewRecorder()
	jsonBody, _ := json.Marshal(updateUser)
	requestBody := bytes.NewBuffer(jsonBody)
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/users/%v", user.ID), requestBody)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var jsonParser fastjson.Parser
	jsonFields, err := jsonParser.Parse(w.Body.String())
	assert.Nil(t, err)
	assert.True(t, jsonFields.Exists("name"))
	assert.True(t, jsonFields.Exists("email"))

	name := string(jsonFields.GetStringBytes("name"))
	assert.Equal(t, updateUser.Name, name)
	email := string(jsonFields.GetStringBytes("email"))
	assert.Equal(t, updateUser.Email, email)
}

func TestV1UserAPIDelete(t *testing.T) {
	router := setupTestV1UserAPIRouter()

	userService := services.NewUserService()
	user := models.User{
		Name:  "user-api-test-delete",
		Email: "user-api-test-delete@example.com",
	}
	user, _ = userService.Create(user)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/users/%v", user.ID), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 204, w.Code)
	assert.Empty(t, w.Body.String())

}

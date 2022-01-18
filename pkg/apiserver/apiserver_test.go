package apiserver_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/chenjr0719/golang-boilerplate/pkg/apiserver"
	"github.com/stretchr/testify/assert"
)

func TestLiveness(t *testing.T) {
	apiServer := apiserver.NewAPIServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/healthz", nil)
	apiServer.Router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestReadiness(t *testing.T) {
	apiServer := apiserver.NewAPIServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/healthz/readiness", nil)
	apiServer.Router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestRedirectDocs(t *testing.T) {
	apiServer := apiserver.NewAPIServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/docs/", nil)
	apiServer.Router.ServeHTTP(w, req)

	assert.Equal(t, 301, w.Code)
}

func TestAPIServerRun(t *testing.T) {
	apiServer := apiserver.NewAPIServer()
	go func() {
		apiServer.Run("0.0.0.0", 8080)
	}()
	time.Sleep(5 * time.Millisecond)

	assert.Error(t, apiServer.Run("0.0.0.0", 8080))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/healthz", nil)
	apiServer.Router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

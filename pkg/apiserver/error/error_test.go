package error

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNewError(t *testing.T) {
	w := httptest.NewRecorder()
	ginCtx, _ := gin.CreateTestContext(w)

	type args struct {
		ctx    *gin.Context
		status int
		err    error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"newError", args{ginCtx, 500, errors.New("Error")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewError(tt.args.ctx, tt.args.status, tt.args.err)
		})
	}
}

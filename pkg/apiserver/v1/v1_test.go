package v1

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNewV1Group(t *testing.T) {
	router := gin.Default()
	type args struct {
		apiGroup *gin.RouterGroup
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"newV1Group", args{router.Group("/")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewV1Group(tt.args.apiGroup)
		})
	}
}

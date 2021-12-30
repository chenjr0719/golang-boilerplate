package v1

import (
	"github.com/gin-gonic/gin"
)

func NewV1Group(apiGroup *gin.RouterGroup) *gin.RouterGroup {
	v1Group := apiGroup.Group("/v1")
	NewV1UserGroup(v1Group)
	NewV1JobGroup(v1Group)

	return v1Group
}

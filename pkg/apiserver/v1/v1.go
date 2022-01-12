package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewV1Group(apiGroup *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {
	v1Group := apiGroup.Group("/v1")
	NewV1UserGroup(v1Group, db)
	NewV1JobGroup(v1Group, db)

	return v1Group
}

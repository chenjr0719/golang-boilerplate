package v1

import (
	"github.com/chenjr0719/golang-boilerplate/pkg/apiserver/error"
	"github.com/chenjr0719/golang-boilerplate/pkg/models"
	"github.com/chenjr0719/golang-boilerplate/pkg/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type V1UserAPI struct {
	UserService services.UserService
}

type UserInput struct {
	Name  string `json:"name"  binding:"required"`
	Email string `json:"email"  binding:"required"`
} //@name UserInput

func NewV1UserGroup(apiGroup *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {
	v1UserAPI := &V1UserAPI{
		UserService: *services.NewUserService(db),
	}

	v1UserGroup := apiGroup.Group("/users")
	v1UserGroup.GET("/", v1UserAPI.ListUsers)
	v1UserGroup.POST("/", v1UserAPI.CreateUser)
	v1UserGroup.GET("/:id", v1UserAPI.GetUser)
	v1UserGroup.PUT("/:id", v1UserAPI.UpdateUser)
	v1UserGroup.DELETE("/:id", v1UserAPI.DeleteUser)

	return v1UserGroup
}

// ListUser godoc
// @Summary List User
// @Schemes
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} error.HTTPError
// @Router /v1/users [get]
func (api *V1UserAPI) ListUsers(ctx *gin.Context) {
	users, err := api.UserService.List()
	if err != nil {
		error.NewError(ctx, 500, err)
		return
	}

	ctx.JSON(200, gin.H{
		"items": users,
	})
}

// CreateUser godoc
// @Summary Create User
// @Schemes
// @Tags Users
// @Accept json
// @Produce json
// @Param user body UserInput true "User Input"
// @Success 200 {object} models.User
// @Failure 400 {object} error.HTTPError
// @Failure 500 {object} error.HTTPError
// @Router /v1/users [post]
func (api *V1UserAPI) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		error.NewError(ctx, 400, err)
		return
	}

	user, err := api.UserService.Create(user)
	if err != nil {
		error.NewError(ctx, 400, err)
		return
	}

	ctx.JSON(200, user)
}

// GetUser godoc
// @Summary Get User
// @Schemes
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 404 {object} error.HTTPError
// @Failure 500 {object} error.HTTPError
// @Router /v1/users/{id} [get]
func (api *V1UserAPI) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := api.UserService.Get(id)
	if err != nil {
		error.NewError(ctx, 404, err)
		return
	}

	ctx.JSON(200, user)
}

// UpdateUser godoc
// @Summary Update User
// @Schemes
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body UserInput true "User Input"
// @Success 200 {object} models.User
// @Failure 400 {object} error.HTTPError
// @Failure 404 {object} error.HTTPError
// @Failure 500 {object} error.HTTPError
// @Router /v1/users/{id} [put]
func (api *V1UserAPI) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		error.NewError(ctx, 400, err)
		return
	}

	user, err := api.UserService.Update(id, user)
	if err != nil {
		error.NewError(ctx, 500, err)
		return
	}

	ctx.JSON(200, user)
}

// DeleteUser godoc
// @Summary Delete User
// @Schemes
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 204 {object} nil
// @Failure 404 {object} error.HTTPError
// @Failure 500 {object} error.HTTPError
// @Router /v1/users/{id} [delete]
func (api *V1UserAPI) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	err := api.UserService.Delete(id)
	if err != nil {
		error.NewError(ctx, 400, err)
		return
	}

	ctx.JSON(204, nil)
}

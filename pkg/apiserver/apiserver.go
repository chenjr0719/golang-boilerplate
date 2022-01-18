package apiserver

import (
	"strconv"

	"github.com/chenjr0719/golang-boilerplate/docs"
	v1 "github.com/chenjr0719/golang-boilerplate/pkg/apiserver/v1"
	"github.com/chenjr0719/golang-boilerplate/pkg/config"
	"github.com/chenjr0719/golang-boilerplate/pkg/db"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type APIServer struct {
	Router *gin.Engine
}

func NewAPIServer() APIServer {
	if config.Config.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	server := APIServer{
		Router: gin.Default(),
	}

	// Default group
	apiGroup := server.Router.Group("/")
	apiGroup.GET("/healthz", liveness)
	apiGroup.GET("/healthz/readiness", readiness)

	// Add v1 APIs
	v1.NewV1Group(apiGroup)

	// Swagger
	docs.SwaggerInfo.BasePath = "/"
	server.Router.GET("/docs/*any", redirectDocs, ginSwagger.WrapHandler(swaggerfiles.Handler))

	return server
}

func liveness(ctx *gin.Context) {
	ctx.String(200, "")
}

func readiness(ctx *gin.Context) {
	sqlDB, err := db.DB.DB()
	if err != nil {
		ctx.String(500, "")
		return
	}
	err = sqlDB.Ping()
	if err != nil {
		ctx.String(500, "")
		return
	}
	ctx.String(200, "")
}

func redirectDocs(ctx *gin.Context) {
	if ctx.Request.URL.Path == "/docs/" {
		ctx.Redirect(301, "/docs/index.html")
		return
	}
}

func (server APIServer) Run(host string, port int) error {
	address := host + ":" + strconv.Itoa(port)
	err := endless.ListenAndServe(address, server.Router)
	return err
}

package apiserver

import (
	"strconv"

	"github.com/chenjr0719/golang-boilerplate/docs"
	v1 "github.com/chenjr0719/golang-boilerplate/pkg/apiserver/v1"
	"github.com/chenjr0719/golang-boilerplate/pkg/config"
	"github.com/chenjr0719/golang-boilerplate/pkg/db"
	"github.com/chenjr0719/golang-boilerplate/pkg/log"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type APIServer struct {
	config *config.Config
	router *gin.Engine
	db     *gorm.DB
}

func NewAPIServer(conf *config.Config) (*APIServer, error) {
	if conf.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	dbConnection, err := db.ConnectDatabase(conf)
	if err != nil {
		log.Fatal().Err(err).Msg("Create API Server failed")
		return nil, err
	}

	server := &APIServer{
		config: conf,
		router: gin.Default(),
		db:     dbConnection,
	}

	// Default group
	apiGroup := server.router.Group("/")
	apiGroup.GET("/healthz", server.liveness)
	apiGroup.GET("/healthz/readiness", server.readiness)

	// Add v1 APIs
	v1.NewV1Group(apiGroup, dbConnection)

	// Swagger
	docs.SwaggerInfo.BasePath = "/"
	server.router.GET("/docs/*any", server.redirectDocs, ginSwagger.WrapHandler(swaggerfiles.Handler))

	return server, nil
}

func (api *APIServer) liveness(ctx *gin.Context) {
	ctx.String(200, "")
}

func (api *APIServer) readiness(ctx *gin.Context) {
	sqlDB, err := api.db.DB()
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

func (api *APIServer) redirectDocs(ctx *gin.Context) {
	if ctx.Request.RequestURI == "/docs/" {
		ctx.Redirect(301, "/docs/index.html")
		return
	}
}

func (server APIServer) Run(host string, port int) error {
	address := host + ":" + strconv.Itoa(port)
	err := endless.ListenAndServe(address, server.router)
	return err
}

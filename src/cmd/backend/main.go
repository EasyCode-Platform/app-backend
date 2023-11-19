package main

import (
	"os"

	"github.com/EasyCode-Platform/app-backend/src/controller"
	"github.com/EasyCode-Platform/app-backend/src/driver/mongodb"
	"github.com/EasyCode-Platform/app-backend/src/driver/postgres"
	"github.com/EasyCode-Platform/app-backend/src/router"
	"github.com/EasyCode-Platform/app-backend/src/storage"
	"github.com/EasyCode-Platform/app-backend/src/utils/accesscontrol"
	"github.com/EasyCode-Platform/app-backend/src/utils/config"
	"github.com/EasyCode-Platform/app-backend/src/utils/cors"
	"github.com/EasyCode-Platform/app-backend/src/utils/logger"
	"github.com/EasyCode-Platform/app-backend/src/utils/recovery"
	"github.com/EasyCode-Platform/app-backend/src/utils/tokenvalidator"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	engine *gin.Engine
	router *router.Router
	logger *zap.SugaredLogger
	config *config.Config
}

func NewServer(config *config.Config, engine *gin.Engine, router *router.Router, logger *zap.SugaredLogger) *Server {
	return &Server{
		engine: engine,
		config: config,
		router: router,
		logger: logger,
	}
}

func initStorage(globalConfig *config.Config, logger *zap.SugaredLogger) *storage.Storage {
	postgresDriver, err := postgres.NewPostgresConnectionByGlobalConfig(globalConfig, logger)
	if err != nil {
		logger.Errorw("Error in startup, postgres init failed.")
	}
	mongodbDriver, err := mongodb.NewMongodbConnectionByGlobalConfig(globalConfig, logger)
	if err != nil {
		logger.Errorw("Error in startup, mongodb init failed")
	}
	return storage.NewStorage(postgresDriver, mongodbDriver, logger)
}

func initServer() (*Server, error) {
	globalConfig := config.GetInstance()
	engine := gin.New()
	sugaredLogger := logger.NewSugardLogger()

	// init validator
	validator := tokenvalidator.NewRequestTokenValidator()

	// init driver
	storage := initStorage(globalConfig, sugaredLogger)

	// init attribute group
	attrg, errInNewAttributeGroup := accesscontrol.NewRawAttributeGroup()
	if errInNewAttributeGroup != nil {
		return nil, errInNewAttributeGroup
	}

	// init controller
	c := controller.NewControllerForBackend(storage, validator, attrg)
	router := router.NewRouter(c)
	server := NewServer(globalConfig, engine, router, sugaredLogger)
	return server, nil

}

func (server *Server) Start() {
	server.logger.Infow("Starting ec-builder-backend...")

	// init
	gin.SetMode(gin.ReleaseMode)

	// init cors
	server.engine.Use(gin.CustomRecovery(recovery.CorsHandleRecovery))
	server.engine.Use(cors.Cors())
	server.router.RegisterRouters(server.engine)
	// run
	server.logger.Infow("ec-builder-backend started")
	err := server.engine.Run(server.config.ServerHost + ":" + server.config.ServerPort)
	if err != nil {
		server.logger.Errorw("Error in startup", "err", err)
		os.Exit(2)
	}

}

func main() {
	server, err := initServer()

	if err != nil {
		server.logger.Errorw("Error in startup, failed to initServer", "err", err)
	}

	server.Start()
}

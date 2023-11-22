package storage

import (
	"testing"

	"github.com/EasyCode-Platform/app-backend/src/driver/mongodb"
	"github.com/EasyCode-Platform/app-backend/src/driver/postgres"
	"github.com/EasyCode-Platform/app-backend/src/request"
	"github.com/EasyCode-Platform/app-backend/src/utils/config"
	"github.com/EasyCode-Platform/app-backend/src/utils/logger"
)

func Test_Storage(t *testing.T) {
	globalConfig := config.GetInstance()
	logger := logger.NewSugardLogger()

	// init validator

	// init driver
	postgresDriver, err := postgres.NewPostgresConnectionByGlobalConfig(globalConfig, logger)
	if err != nil {
		t.Errorf("Error in startup, postgres init failed.")
	}
	mongodbDriver, err := mongodb.NewMongodbConnectionByGlobalConfig(globalConfig, logger)
	if err != nil {
		t.Errorf("Error in startup, mongodb init failed")
	}
	storage := NewStorage(postgresDriver, mongodbDriver, logger)
	ans, err := storage.PostgresStorage.ExecutePostgresSql(request.NewExecuteSqlRequest("Select * From test"))
	println("ans of test %+v", ans)
}

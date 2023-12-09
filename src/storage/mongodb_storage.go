package storage

import (
	"github.com/EasyCode-Platform/app-backend/src/request"
	"github.com/EasyCode-Platform/app-backend/src/response"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type MongoDbStorage struct {
	logger *zap.SugaredLogger
	db     *mongo.Database
}

// NewMongoDb
// @param db
// @param collection
// @return *MongoDbStorage
func NewMongoDb(logger *zap.SugaredLogger, db *mongo.Database) *MongoDbStorage {
	return &MongoDbStorage{logger: logger, db: db}
}

func (impl *MongoDbStorage) ExecuteMongodbSql(sql *request.ExecuteSqlRequest) (*response.SqlResponse, error) {
	return nil, nil
}

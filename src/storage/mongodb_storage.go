package storage

import (
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

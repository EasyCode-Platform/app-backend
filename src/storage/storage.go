package storage

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Storage struct {
	MongodbStorage  *MongoDbStorage
	PostgresStorage *PostgresStorage
	// ActionStorage      *ActionStorage
	// AppSnapshotStorage *AppSnapshotStorage
	// KVStateStorage     *KVStateStorage
	// ResourceStorage    *ResourceStorage
	// SetStateStorage    *SetStateStorage
	// TreeStateStorage   *TreeStateStorage
}

// NewStorage
// @param postgresDriver
// @param mongodb
// @param logger
// @return *Storage
func NewStorage(postgresDriver *gorm.DB, mongodb *mongo.Database, logger *zap.SugaredLogger) *Storage {
	return &Storage{
		MongodbStorage:  NewMongoDb(logger, mongodb),
		PostgresStorage: NewPostgresStorage(logger, postgresDriver),
		// ActionStorage:      NewActionStorage(logger, postgresDriver),
		// AppSnapshotStorage: NewAppSnapshotStorage(logger, postgresDriver),
		// KVStateStorage:     NewKVStateStorage(logger, postgresDriver),
		// ResourceStorage:    NewResourceStorage(logger, postgresDriver),
		// SetStateStorage:    NewSetStateStorage(logger, postgresDriver),
		// TreeStateStorage:   NewTreeStateStorage(logger, postgresDriver),
	}
}

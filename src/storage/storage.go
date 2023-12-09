package storage

import (
	"github.com/EasyCode-Platform/app-backend/src/driver/minio"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Storage struct {
	MongodbStorage  *MongoDbStorage
	PostgresStorage *PostgresStorage
	ImageStorage    *ImageStorage
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
func NewStorage(postgresDriver *gorm.DB, mongodb *mongo.Database, imageS3Drive *minio.S3Drive, logger *zap.SugaredLogger) *Storage {
	return &Storage{
		MongodbStorage:  NewMongoDb(logger, mongodb),
		PostgresStorage: NewPostgresStorage(logger, postgresDriver),
		ImageStorage:    NewImageStorage(logger, imageS3Drive),
		// ActionStorage:      NewActionStorage(logger, postgresDriver),
		// AppSnapshotStorage: NewAppSnapshotStorage(logger, postgresDriver),
		// KVStateStorage:     NewKVStateStorage(logger, postgresDriver),
		// ResourceStorage:    NewResourceStorage(logger, postgresDriver),
		// SetStateStorage:    NewSetStateStorage(logger, postgresDriver),
		// TreeStateStorage:   NewTreeStateStorage(logger, postgresDriver),
	}
}

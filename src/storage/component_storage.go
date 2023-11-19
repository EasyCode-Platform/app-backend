package storage

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type ComponentStorage struct {
	logger  *zap.SugaredLogger
	storage *MongoDbStorage
}

// NewComponentStorage
// @param logger
// @param db
// @return *ComponentStorage
func NewComponentStorage(logger *zap.SugaredLogger, db *mongo.Database) *ComponentStorage {
	return &ComponentStorage{
		logger:  logger,
		storage: NewMongoDb(db, "components"),
	}
}

// @receiver impl
// @param componentJson the json string of the component
// @return string the id of the inserted component inside mongodb collection
// @return error
func (impl *ComponentStorage) CreateNewComponent(componentJson string) (string, error) {
	str, err := impl.storage.InsertToDb(componentJson)
	if err != nil {
		impl.logger.Errorw("failed to insert %s into mongodb", componentJson)
		return "", err
	}
	return str, nil
}

// RetrieveComponent
// @receiver impl
// @param ComponentId
// @return string
// @return error
func (impl *ComponentStorage) RetrieveComponent(ComponentId string) (string, error) {
	str, err := impl.storage.FindInfoById(ComponentId)
	if err != nil {
		impl.logger.Errorw("Can not retrieve component %", ComponentId)
		return "", err
	}
	return str, nil
}

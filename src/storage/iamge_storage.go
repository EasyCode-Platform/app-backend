package storage

import (
	"context"
	"mime/multipart"

	"github.com/EasyCode-Platform/app-backend/src/driver/minio"
	"go.uber.org/zap"
)

const IMAGE_STORAGE_BUCKET_NAME = "image"

type ImageStorage struct {
	logger *zap.SugaredLogger
	driver *minio.S3Drive
}

func NewImageStorage(logger *zap.SugaredLogger,
	driver *minio.S3Drive) *ImageStorage {
	return &ImageStorage{
		logger: logger,
		driver: driver,
	}
}

func (impl *ImageStorage) UploadImage(file *multipart.FileHeader) error {
	content, err := file.Open()
	if err != nil {
		return err
	}
	info, err := impl.driver.Instance.PutObject(context.TODO(), impl.driver.Config.BucketName, file.Filename, content, file.Size, impl.driver.Config.PutObjectOption)
	if err != nil {
		impl.logger.Errorf("Failed to insert file: %s with error: %s", file.Filename, err)
		return err
	}
	impl.logger.Infof("successfully store object with info : %s", info)
	return nil
}

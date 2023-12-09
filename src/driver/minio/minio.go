package minio

import (
	"context"
	"time"

	"github.com/EasyCode-Platform/app-backend/src/utils/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
)

type MINIOConfig struct {
	AccessKeyID      string
	AccessKeySecret  string
	Endpoint         string
	BucketName       string
	SSLEnabled       bool
	UploadTimeout    time.Duration
	MakeBucketOption minio.MakeBucketOptions
	PutObjectOption  minio.PutObjectOptions
}

func NewImageMINIOConfigByGlobalConfig(config *config.Config) *MINIOConfig {
	return &MINIOConfig{
		AccessKeyID:      config.GetMINIOAccessKeyID(),
		AccessKeySecret:  config.GetMINIOAccessKeySecret(),
		Endpoint:         config.GetMINIOEndpoint(),
		BucketName:       config.GetMINIOImageBucketName(),
		UploadTimeout:    config.GetMINIOTimeout(),
		MakeBucketOption: minio.MakeBucketOptions{},
		PutObjectOption:  minio.PutObjectOptions{},
	}
}

func CreateMINIOInstance(minioConfig *MINIOConfig) (*minio.Client, error) {
	minioInstance, err := minio.New(minioConfig.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioConfig.AccessKeyID, minioConfig.AccessKeySecret, ""),
		Secure: minioConfig.SSLEnabled,
	})
	if err != nil {
		return nil, err
	}
	return minioInstance, nil
}

type S3Drive struct {
	Instance *minio.Client
	Config   *MINIOConfig
	logger   *zap.SugaredLogger
}

func NewS3Drive(logger *zap.SugaredLogger, minioConfig *MINIOConfig) (*S3Drive, error) {
	s3Drive := &S3Drive{
		Config: minioConfig,
		logger: logger,
	}
	var err error
	s3Drive.Instance, err = CreateMINIOInstance(minioConfig)
	if err != nil {
		logger.Errorf("Failed to Create Minio Instance with error : %s", err)
		return nil, err
	}
	err = s3Drive.initDefaultBucket()
	if err != nil {
		logger.Errorf("Failed to init Default Bucket with error : %s", err)
		return nil, err
	}
	return s3Drive, nil
}

func (s3Drive *S3Drive) initDefaultBucket() error {
	ctx := context.Background()
	bucketName := s3Drive.Config.BucketName
	err := s3Drive.Instance.MakeBucket(ctx, bucketName, s3Drive.Config.MakeBucketOption)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := s3Drive.Instance.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			s3Drive.logger.Infof("We already own bucket \"%s\"\n", bucketName)
		} else {
			s3Drive.logger.Errorf("Failed to make bucket : %s with error: %s", bucketName, err)
			return err
		}
	} else {
		s3Drive.logger.Infof("Successfully created bucket \"%s\"\n", bucketName)
	}
	return nil
}

func (s3Drive *S3Drive) GetPreSignedPutURL(fileName string) (string, error) {
	ctx := context.Background()
	// get put request
	presignedURL, err := s3Drive.Instance.PresignedPutObject(ctx, s3Drive.Config.BucketName, fileName, s3Drive.Config.UploadTimeout)
	if err != nil {
		return "", err
	}
	return presignedURL.String(), nil
}

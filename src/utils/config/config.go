package config

import (
	"fmt"
	"sync"
	"time"

	"github.com/caarlos0/env/v10"
)

const DEPLOY_MODE_SELF_HOST = "self-host"
const DEPLOY_MODE_CLOUD = "cloud"
const DEPLOY_MODE_CLOUD_TEST = "cloud-test"
const DEPLOY_MODE_CLOUD_BETA = "cloud-beta"
const DEPLOY_MODE_CLOUD_PRODUCTION = "cloud-production"
const DRIVE_TYPE_AWS = "aws"
const DRIVE_TYPE_DO = "do"
const DRIVE_TYPE_MINIO = "minio"
const PROTOCOL_WEBSOCKET = "ws"
const PROTOCOL_WEBSOCKET_OVER_TLS = "wss"

var instance *Config
var once sync.Once

func GetInstance() *Config {
	once.Do(func() {
		var err error
		if instance == nil {
			instance, err = getConfig() // not thread safe
			if err != nil {
				panic(err)
			}
		}
	})
	return instance
}

type Config struct {
	// server config
	ServerHost         string `env:"SERVER_HOST"                    envDefault:"0.0.0.0"`
	ServerPort         string `env:"SERVER_PORT"                    envDefault:"8003"`
	InternalServerPort string `env:"SERVER_INTERNAL_PORT"           envDefault:"9005"`
	ServerMode         string `env:"SERVER_MODE"                    envDefault:"release"`
	DeployMode         string `env:"DEPLOY_MODE"                    envDefault:"self-host"`
	SecretKey          string `env:"SECRET_KEY"                     envDefault:"8xEMrWkBARcDDYQ"`
	WSSEnabled         string `env:"WSS_ENABLED"                    envDefault:"false"`

	// key for idconvertor
	RandomKey string `env:"RANDOM_KEY"  envDefault:"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"`
	// storage config
	PostgresAddr         string `env:"PG_ADDR" envDefault:"localhost"`
	PostgresPort         string `env:"PG_PORT" envDefault:"5432"`
	PostgresUser         string `env:"PG_USER" envDefault:"app_backend"`
	PostgresPassword     string `env:"PG_PASSWORD" envDefault:"scut2023"`
	PostgresDatabase     string `env:"PG_DATABASE" envDefault:"app_backend"`
	MongodbAddr          string `env:"MONGODB_ADDR" envDefault:"localhost"`
	MongodbPort          string `env:"MONGODB_PORT" envDefault:"27017"`
	MongodbDatabase      string `env:"MONGODB_DATABASE" envDefault:"app_backend"`
	MongodbMaxCollection int64  `env:"MONGODB_MAXCOLLECTION" envDefault:"10"`
	// cache config
	RedisAddr     string `env:"REDIS_ADDR" envDefault:"localhost"`
	RedisPort     string `env:"REDIS_PORT" envDefault:"6379"`
	RedisPassword string `env:"REDIS_PASSWORD" envDefault:"ec2022"`
	RedisDatabase int    `env:"REDIS_DATABASE" envDefault:"0"`
	// drive config
	DriveType             string `env:"DRIVE_TYPE"               envDefault:""`
	DriveAccessKeyID      string `env:"DRIVE_ACCESS_KEY_ID"      envDefault:"root"`
	DriveAccessKeySecret  string `env:"DRIVE_ACCESS_KEY_SECRET"  envDefault:"scut2023"`
	DriveRegion           string `env:"DRIVE_REGION"             envDefault:""`
	DriveEndpoint         string `env:"DRIVE_ENDPOINT"           envDefault:"localhost:9000"`
	DriveImageBucketName  string `env:"DRIVE_IMAGE_BUCKET_NAME"  envDefault:"image"`
	DriveUploadTimeoutRaw string `env:"DRIVE_UPLOAD_TIMEOUT"     envDefault:"30s"`
	DriveUploadTimeout    time.Duration
	// supervisor API
	ecSupervisorInternalRestAPI string `env:"SUPERVISOR_INTERNAL_API"     envDefault:"http://127.0.0.1:9001/api/v1"`

	// peripheral API
	ecPeripheralAPI string `env:"PERIPHERAL_API" envDefault:"https://peripheral-api.ecsoft.com/v1/"`
	// resource manager API
	ecResourceManagerRestAPI         string `env:"RESOURCE_MANAGER_API"     envDefault:"http://app-resource-manager-backend:8006"`
	ecResourceManagerInternalRestAPI string `env:"RESOURCE_MANAGER_INTERNAL_API"     envDefault:"http://app-resource-manager-backend-internal:9004"`
	// ec marketplace config
	ecMarketplaceInternalRestAPI string `env:"MARKETPLACE_INTERNAL_API"     envDefault:"http://app-marketplace-backend-internal:9003/api/v1"`
	// token for internal api
	ControlToken string `env:"CONTROL_TOKEN"     envDefault:""`
	// google config
	ecGoogleSheetsClientID     string `env:"GS_CLIENT_ID"           envDefault:""`
	ecGoogleSheetsClientSecret string `env:"GS_CLIENT_SECRET"       envDefault:""`
	ecGoogleSheetsRedirectURI  string `env:"GS_REDIRECT_URI"        envDefault:""`
}

func getConfig() (*Config, error) {
	// fetch
	cfg := &Config{}
	err := env.Parse(cfg)
	// process data
	var errInParseDuration error
	cfg.DriveUploadTimeout, errInParseDuration = time.ParseDuration(cfg.DriveUploadTimeoutRaw)
	if errInParseDuration != nil {
		return nil, errInParseDuration
	}
	// ok
	fmt.Printf("----------------\n")
	fmt.Printf("%+v\n", cfg)
	fmt.Printf("%+v\n", err)

	return cfg, err
}

func (c *Config) IsSelfHostMode() bool {
	return c.DeployMode == DEPLOY_MODE_SELF_HOST
}

func (c *Config) IsCloudMode() bool {
	if c.DeployMode == DEPLOY_MODE_CLOUD || c.DeployMode == DEPLOY_MODE_CLOUD_TEST || c.DeployMode == DEPLOY_MODE_CLOUD_BETA || c.DeployMode == DEPLOY_MODE_CLOUD_PRODUCTION {
		return true
	}
	return false
}

func (c *Config) IsCloudTestMode() bool {
	return c.DeployMode == DEPLOY_MODE_CLOUD_TEST
}

func (c *Config) IsCloudBetaMode() bool {
	return c.DeployMode == DEPLOY_MODE_CLOUD_BETA
}

func (c *Config) IsCloudProductionMode() bool {
	return c.DeployMode == DEPLOY_MODE_CLOUD_PRODUCTION
}

func (c *Config) GetWebsocketProtocol() string {
	if c.WSSEnabled == "true" {
		return PROTOCOL_WEBSOCKET_OVER_TLS
	}
	return PROTOCOL_WEBSOCKET
}

func (c *Config) GetRuntimeEnv() string {
	if c.IsCloudBetaMode() {
		return DEPLOY_MODE_CLOUD_BETA
	} else if c.IsCloudProductionMode() {
		return DEPLOY_MODE_CLOUD_PRODUCTION
	} else {
		return DEPLOY_MODE_CLOUD_TEST
	}
}

func (c *Config) GetSecretKey() string {
	return c.SecretKey
}

func (c *Config) GetRandomKey() string {
	return c.RandomKey
}

func (c *Config) GetPostgresAddr() string {
	return c.PostgresAddr
}

func (c *Config) GetPostgresPort() string {
	return c.PostgresPort
}

func (c *Config) GetPostgresUser() string {
	return c.PostgresUser
}

func (c *Config) GetPostgresPassword() string {
	return c.PostgresPassword
}

func (c *Config) GetPostgresDatabase() string {
	return c.PostgresDatabase
}

func (c *Config) GetRedisAddr() string {
	return c.RedisAddr
}

func (c *Config) GetRedisPort() string {
	return c.RedisPort
}

func (c *Config) GetRedisPassword() string {
	return c.RedisPassword
}

func (c *Config) GetRedisDatabase() int {
	return c.RedisDatabase
}

func (c *Config) GetDriveType() string {
	return c.DriveType
}

func (c *Config) IsAWSTypeDrive() bool {
	if c.DriveType == DRIVE_TYPE_AWS || c.DriveType == DRIVE_TYPE_DO {
		return true
	}
	return false
}

func (c *Config) IsMINIODrive() bool {
	return c.DriveType == DRIVE_TYPE_MINIO
}

func (c *Config) GetAWSS3Endpoint() string {
	return c.DriveEndpoint
}

func (c *Config) GetAWSS3AccessKeyID() string {
	return c.DriveAccessKeyID
}

func (c *Config) GetAWSS3AccessKeySecret() string {
	return c.DriveAccessKeySecret
}

func (c *Config) GetAWSS3Region() string {
	return c.DriveRegion
}

func (c *Config) GetAWSS3ImageBucketName() string {
	return c.DriveImageBucketName
}

func (c *Config) GetAWSS3Timeout() time.Duration {
	return c.DriveUploadTimeout
}

func (c *Config) GetMINIOAccessKeyID() string {
	return c.DriveAccessKeyID
}

func (c *Config) GetMINIOAccessKeySecret() string {
	return c.DriveAccessKeySecret
}

func (c *Config) GetMINIOEndpoint() string {
	return c.DriveEndpoint
}

func (c *Config) GetMINIOImageBucketName() string {
	return c.DriveImageBucketName
}

func (c *Config) GetMINIOTimeout() time.Duration {
	return c.DriveUploadTimeout
}

func (c *Config) GetControlToken() string {
	return c.ControlToken
}

func (c *Config) GetecSupervisorInternalRestAPI() string {
	return c.ecSupervisorInternalRestAPI
}

func (c *Config) GetecPeripheralAPI() string {
	return c.ecPeripheralAPI
}

func (c *Config) GetecResourceManagerRestAPI() string {
	return c.ecResourceManagerRestAPI
}

func (c *Config) GetecResourceManagerInternalRestAPI() string {
	return c.ecResourceManagerInternalRestAPI
}

func (c *Config) GetecMarketplaceInternalRestAPI() string {
	return c.ecMarketplaceInternalRestAPI
}

func (c *Config) GetecGoogleSheetsClientID() string {
	return c.ecGoogleSheetsClientID
}

func (c *Config) GetecGoogleSheetsClientSecret() string {
	return c.ecGoogleSheetsClientSecret
}

func (c *Config) GetecGoogleSheetsRedirectURI() string {
	return c.ecGoogleSheetsRedirectURI
}
func (c *Config) GetMongodbAddr() string {
	return c.MongodbAddr
}
func (c *Config) GetMongodbPort() string {
	return c.MongodbPort
}
func (c *Config) GetMongodbDatabase() string {
	return c.MongodbDatabase
}
func (c *Config) GetMongodbMaxCollection() int64 {
	return c.MongodbMaxCollection
}

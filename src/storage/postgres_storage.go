package storage

import (
	"github.com/EasyCode-Platform/app-backend/src/request"
	"github.com/EasyCode-Platform/app-backend/src/response"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PostgresStorage struct {
	logger *zap.SugaredLogger
	db     *gorm.DB
}

func NewPostgresStorage(logger *zap.SugaredLogger, db *gorm.DB) *PostgresStorage {
	return &PostgresStorage{
		logger: logger,
		db:     db,
	}
}

func (storage *Storage) ExecutePostgresSql(sql *request.ExecuteSqlRequest) *response.SqlResponse {
	var ans interface{}
	storage.PostgresStorage.db.Raw(sql.SqlStatement, sql.SqlValues).Scan(&ans)
	storage.PostgresStorage.logger.Info("return from storage.ExecutePostgresSql:{:?}", ans)
	print("return from storage.ExecutePostgresSql:{:?}", ans)
	return response.NewSqlResponse(ans)
}

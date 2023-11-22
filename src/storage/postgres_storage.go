package storage

import (
	"fmt"

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

func (impl *PostgresStorage) ExecutePostgresSql(sql *request.ExecuteSqlRequest) (*response.SqlResponse, error) {
	var ans interface{}
	err := impl.db.Raw(sql.SqlStatement, sql.SqlValues).Scan(&ans).Error
	if err != nil {
		impl.logger.Errorf("Failed to execute sql %+v with error %+v", sql, err)
	}
	impl.logger.Infof("return from storage.ExecutePostgresSql: %+v", ans)
	fmt.Printf("return from storage.ExecutePostgresSql %+v", ans)
	return response.NewSqlResponse(ans), nil
}

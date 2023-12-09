package storage

import (
	"fmt"
	"strings"

	"github.com/EasyCode-Platform/app-backend/src/model"
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

func (impl *PostgresStorage) CreateTable(table *model.Table) error {
	var ans interface{}
	if err := impl.db.Raw(fmt.Sprintf("create table if not exists %s(    id bigserial not null primary key,%s);", table.TableName, table.ExportDefs())).Scan(&ans).Error; err != nil {
		impl.logger.Errorf("Failed to create table %+v with error %+v", table, err)
		return err
	}
	return nil
}

func (impl *PostgresStorage) InsertRecord(table *model.Table, record *model.Record) error {
	valueString := ""
	nameString := ""
	var ans map[string]interface{}
	impl.logger.Infof("Inserting %+v into table %+v", record, table)
	for columnName, columnType := range table.Columns {
		switch columnType {
		case "text":
			{
				nameString += columnName
				nameString += ","
				record.Record[columnName] = record.Record[columnName].(string)
				valueString += fmt.Sprintf("'%s'", record.Record[columnName])
				valueString += ","
			}
		case "integer":
			{
				nameString += columnName
				nameString += ","
				record.Record[columnName] = record.Record[columnName].(string)
				valueString += fmt.Sprintf("%s", record.Record[columnName])
				valueString += ","
			}
		default:
			{
				return fmt.Errorf("Unsupported type of column")
			}
		}
	}
	valueString = strings.TrimRight(valueString, ",")
	nameString = strings.TrimRight(nameString, ",")
	if err := impl.db.Raw(fmt.Sprintf("INSERT INTO %s(%s)VALUES(%s);", table.TableName, nameString, valueString)).Scan(&ans).Error; err != nil {
		impl.logger.Errorf("Failed to insert record into table: %s with error %+v", table.TableName, err)
		return err
	}
	return nil
}

func (impl *PostgresStorage) DisplayTable(tableName string) (*response.SqlResponse, error) {
	var ans map[string]interface{}
	if err := impl.db.Raw(fmt.Sprintf("SELECT * FROM  %s ;", tableName)).Scan(&ans).Error; err != nil {
		impl.logger.Errorf("Failed to retrieve all records in table: %s with error %+v", tableName, err)
		return nil, err
	}
	return response.NewSqlResponse(ans), nil
}

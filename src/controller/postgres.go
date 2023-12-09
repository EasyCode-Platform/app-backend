package controller

import (
	"github.com/EasyCode-Platform/app-backend/src/model"
	"github.com/EasyCode-Platform/app-backend/src/request"
	"github.com/gin-gonic/gin"
	"github.com/xwb1989/sqlparser"
)

func (controller *Controller) ExecutePostgresSql(c *gin.Context) {
	rawStatement := c.Request.Context().Value("sqlStatement").(string)
	rawStatements, err := sqlparser.SplitStatementToPieces(rawStatement)
	if err != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_SQL_INVALID, "validate sql sentence  error: "+err.Error())
		return
	}
	for _, str := range rawStatements {
		_, err := sqlparser.Parse(str)
		if err != nil {
			controller.FeedbackBadRequest(c, ERROR_FLAG_SQL_INVALID, "validate sql sentence error: "+err.Error())
			return
		}
	}
	if err != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_SQL_INVALID, "validate sql sentence error: "+err.Error())
		return
	}
	var executeSqlRequest *request.ExecuteSqlRequest
	if c.Request.Context().Value("sqlValues").([]interface{}) != nil {
		executeSqlRequest = request.NewExecuteSqlRequestWithValues(rawStatement, c.Request.Context().Value("sqlValues").([]interface{}))
	}
	{
		executeSqlRequest = request.NewExecuteSqlRequest(rawStatement)
	}
	sqlResponse, err := controller.Storage.PostgresStorage.ExecutePostgresSql(executeSqlRequest)
	if err != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_SQL_INVALID, "validate sql sentence error: "+err.Error())
	}
	controller.FeedbackOK(c, sqlResponse)
}

func (controller *Controller) ValidatePostgresSql(c *gin.Context) {
	rawStatement := c.Request.Context().Value("sqlStatement").(string)
	rawStatements, err := sqlparser.SplitStatementToPieces(rawStatement)
	if err != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_SQL_INVALID, "validate sql sentence  error: "+err.Error())
		return
	}
	for _, str := range rawStatements {
		_, err := sqlparser.Parse(str)
		if err != nil {
			controller.FeedbackBadRequest(c, ERROR_FLAG_SQL_INVALID, "validate sql sentence  error: "+err.Error())
			return
		}
	}
	if err != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_SQL_INVALID, "validate sql sentence  error: "+err.Error())
		return
	}
}
func (controller *Controller) CreateTable(c *gin.Context) {
	table := model.Table{}
	if err := c.ShouldBind(&table); err != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_PARSE_TABLE_LIST_FAILED, "Failed to retrieve table colums from request error:"+err.Error())
		return
	}
	if err := controller.Storage.PostgresStorage.CreateTable(&table); err != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_CREATE_TABLE_FAILED, "Failed to create table error:"+err.Error())
		return
	}
	controller.FeedbackOK(c, nil)
}

func (controller *Controller) InsertRecord(c *gin.Context) {
	recordRequest := request.RecordRequest{}
	if err := c.ShouldBind(&recordRequest); err != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_PARSE_TABLE_LIST_FAILED, "Failed to retrieve recordRequest from request error:"+err.Error())
		return
	}
	if err := controller.Storage.PostgresStorage.InsertRecord(recordRequest.Table, recordRequest.Record); err != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_INSERT_RECORD_FAILED, "Failed to insert record error:"+err.Error())
		return
	}
	controller.FeedbackOK(c, nil)
}

func (controller *Controller) RemoveRecord(c *gin.Context) {
	return
}

func (controller *Controller) DisplayTable(c *gin.Context) {
	table := request.DisplayRecordRequest{}
	if err := c.ShouldBind(&table); err != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_PARSE_TABLE_LIST_FAILED, "Failed to retrieve table colums from request error:"+err.Error())
		return
	}
	ans, err := controller.Storage.PostgresStorage.DisplayTable(table.TableName)
	if err != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_DISPLAY_RECORD_FAILED, "Failed to retrieve all records error:"+err.Error())
		return
	}
	controller.FeedbackOK(c, ans)
}

package controller

import (
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

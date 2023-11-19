package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xwb1989/sqlparser"
)

func (controller *Controller) ExecuteSql(c *gin.Context) {
	sqlSentence := c.Request.Context().Value("sqlSentence").(string)
	_, err := sqlparser.Parse(sqlSentence)
	if err != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_SQL_INVALID, "validate sql sentence  error: "+err.Error())
	}

}

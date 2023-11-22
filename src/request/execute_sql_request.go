package request

type ExecuteSqlRequest struct {
	SqlStatement string        `json:"sqlStatement" validate:"require"`
	SqlValues    []interface{} `json:"sqlValues" `
}

func NewExecuteSqlRequest(sqlStatement string) *ExecuteSqlRequest {
	return &ExecuteSqlRequest{
		SqlStatement: sqlStatement,
		SqlValues:    nil,
	}
}
func NewExecuteSqlRequestWithValues(sqlStatement string, sqlValues []interface{}) *ExecuteSqlRequest {
	return &ExecuteSqlRequest{
		SqlStatement: sqlStatement,
		SqlValues:    sqlValues,
	}
}

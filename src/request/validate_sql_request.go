package request

type ValidateSqlRequest struct {
	SqlStatement string `json:"sqlStatement" validate:"require"`
}

func NewValidateSqlRequest(sqlStatement string) *ValidateSqlRequest {
	return &ValidateSqlRequest{
		SqlStatement: sqlStatement,
	}
}

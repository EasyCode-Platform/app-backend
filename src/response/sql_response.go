package response

type SqlResponse struct {
	result interface{}
}

func NewSqlResponse(result interface{}) *SqlResponse {
	return &SqlResponse{
		result: result,
	}
}

func (response *SqlResponse) ExportForFeedback() interface{} {
	return response.result
}

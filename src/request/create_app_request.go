package request

type CreateAppRequest struct {
	Name       string `json:"appName" validate:"required"`
	InitScheme string `json:"initScheme"`
}

func NewCreateAppRequest() *CreateAppRequest {
	return &CreateAppRequest{}
}

func (req *CreateAppRequest) ExportAppName() string {
	return req.Name
}

func (req *CreateAppRequest) ExportInitScheme() string {
	return req.InitScheme
}

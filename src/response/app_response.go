package response

import "github.com/EasyCode-Platform/app-backend/src/model"

type AppListResponse struct {
	AppList []*model.App
}

func (impl *AppListResponse) ExportForFeedback() interface{} {
	return impl.AppList
}

func NewAppListResponse(appList []*model.App) *AppListResponse {
	return &AppListResponse{
		AppList: appList,
	}
}

type AppResponse struct {
	App *model.App
}

func (impl *AppResponse) ExportForFeedback() interface{} {
	return impl.App
}

func NewAppResponse(app *model.App) *AppResponse {
	return &AppResponse{
		App: app,
	}
}

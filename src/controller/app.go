package controller

import (
	"encoding/json"

	"github.com/EasyCode-Platform/app-backend/src/model"
	"github.com/EasyCode-Platform/app-backend/src/request"
	"github.com/EasyCode-Platform/app-backend/src/response"
	"github.com/EasyCode-Platform/app-backend/src/utils/accesscontrol"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// CreateApp
// @receiver controller
// @param c
func (controller *Controller) CreateApp(c *gin.Context) {
	// fetch needed param
	teamID, errInGetTeamID := controller.GetMagicIntParamFromRequest(c, PARAM_TEAM_ID)
	userID, errInGetUserID := controller.GetUserIDFromAuth(c)
	userAuthToken, errInGetAuthToken := controller.GetUserAuthTokenFromHeader(c)
	if errInGetTeamID != nil || errInGetUserID != nil || errInGetAuthToken != nil {
		return
	}

	// Parse request body
	req := request.NewCreateAppRequest()
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_PARSE_REQUEST_BODY_FAILED, "parse request body error: "+err.Error())
		return
	}

	// Validate request body
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_VALIDATE_REQUEST_BODY_FAILED, "validate request body error: "+err.Error())
		return
	}

	// validate
	canManage, errInCheckAttr := controller.AttributeGroup.CanManage(
		teamID,
		userAuthToken,
		accesscontrol.UNIT_TYPE_APP,
		accesscontrol.DEFAULT_UNIT_ID,
		accesscontrol.ACTION_MANAGE_CREATE_APP,
	)
	if errInCheckAttr != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_ACCESS_DENIED, "error in check attribute: "+errInCheckAttr.Error())
		return
	}
	if !canManage {
		controller.FeedbackBadRequest(c, ERROR_FLAG_ACCESS_DENIED, "you can not access this attribute due to access control policy.")
		return
	}

	ComponentId, err := controller.Storage.ComponentStorage.CreateNewComponent(req.InitScheme)
	if err != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_CAN_NOT_CREATE_APP, "Failed to initialize the components of the app")
		return
	}

	newApp := model.NewApp(req.Name, teamID, userID, ComponentId)

	// create app
	_, errInCreateApp := controller.Storage.AppStorage.CreateApp(newApp)
	if errInCreateApp != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_CAN_NOT_CREATE_APP, "error in create app: "+errInCreateApp.Error())
		return
	}
	appResponse := response.NewAppResponse(newApp)
	controller.FeedbackOK(c, appResponse)
}

// RetrieveApp
// @receiver controller
// @param c
func (controller *Controller) RetrieveApp(c *gin.Context) {
	// fetch needed param
	teamID, errInGetTeamID := controller.GetMagicIntParamFromRequest(c, PARAM_TEAM_ID)
	_, errInGetUserID := controller.GetUserIDFromAuth(c)
	userAuthToken, errInGetAuthToken := controller.GetUserAuthTokenFromHeader(c)
	if errInGetTeamID != nil || errInGetUserID != nil || errInGetAuthToken != nil {
		return
	}

	// Parse request body
	req := request.NewCreateAppRequest()
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_PARSE_REQUEST_BODY_FAILED, "parse request body error: "+err.Error())
		return
	}

	// Validate request body
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_VALIDATE_REQUEST_BODY_FAILED, "validate request body error: "+err.Error())
		return
	}

	// validate
	canManage, errInCheckAttr := controller.AttributeGroup.CanAccess(
		teamID,
		userAuthToken,
		accesscontrol.UNIT_TYPE_APP,
		accesscontrol.DEFAULT_UNIT_ID,
		accesscontrol.ACTION_MANAGE_CREATE_APP,
	)
	if errInCheckAttr != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_ACCESS_DENIED, "error in check attribute: "+errInCheckAttr.Error())
		return
	}
	if !canManage {
		controller.FeedbackBadRequest(c, ERROR_FLAG_ACCESS_DENIED, "you can not access this attribute due to access control policy.")
		return
	}
	// get request
}

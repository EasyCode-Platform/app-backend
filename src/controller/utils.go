package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/EasyCode-Platform/app-backend/src/response"
	"github.com/EasyCode-Platform/app-backend/src/utils/idconvertor"
	"github.com/gin-gonic/gin"
)

const (
	PARAM_SQL_STATEMENT = "sql_statement"
)

const (
	SUCCESS_FLAG = "SUCCESSFULLY REQUESTED"
	// request error
	ERROR_FLAG_PARSE_REQUEST_FILE_FAILED  = "ERROR_FLAG_PARSE_REQUEST_FILE_FAILED"
	ERROR_FLAG_PARSE_TABLE_LIST_FAILED    = "ERROR_FLAG_PARSE_TABLE_LIST_FAILED"
	ERROR_FLAG_PARSE_RECORDREQUEST_FAILED = "ERROR_FLAG_PARSE_RECORDREQUEST_FAILED"
	// validate failed
	ERROR_FLAG_VALIDATE_ACCOUNT_FAILED                  = "ERROR_FLAG_VALIDATE_ACCOUNT_FAILED"
	ERROR_FLAG_VALIDATE_REQUEST_BODY_FAILED             = "ERROR_FLAG_VALIDATE_REQUEST_BODY_FAILED"
	ERROR_FLAG_VALIDATE_REQUEST_TOKEN_FAILED            = "ERROR_FLAG_VALIDATE_REQUEST_TOKEN_FAILED"
	ERROR_FLAG_VALIDATE_REQUEST_PARAM_FAILED            = "ERROR_FLAG_VALIDATE_REQUEST_PARAM_FAILED"
	ERROR_FLAG_VALIDATE_VERIFICATION_CODE_FAILED        = "ERROR_FLAG_VALIDATE_VERIFICATION_CODE_FAILED"
	ERROR_FLAG_VALIDATE_RESOURCE_FAILED                 = "ERROR_FLAG_VALIDATE_RESOURCE_FAILED"
	ERROR_FLAG_PARSE_REQUEST_BODY_FAILED                = "ERROR_FLAG_PARSE_REQUEST_BODY_FAILED"
	ERROR_FLAG_PARSE_REQUEST_URI_FAILED                 = "ERROR_FLAG_PARSE_REQUEST_URI_FAILED"
	ERROR_FLAG_PARSE_INVITE_LINK_HASH_FAILED            = "ERROR_FLAG_PARSE_INVITE_LINK_HASH_FAILED"
	ERROR_FLAG_CAN_NOT_TRANSFER_OWNER_TO_PENDING_USER   = "ERROR_FLAG_CAN_NOT_TRANSFER_OWNER_TO_PENDING_USER"
	ERROR_FLAG_CAN_NOT_REMOVE_OWNER_FROM_TEAM           = "ERROR_FLAG_CAN_NOT_REMOVE_OWNER_FROM_TEAM"
	ERROR_FLAG_SIGN_UP_EMAIL_MISMATCH                   = "ERROR_FLAG_SIGN_UP_EMAIL_MISMATCH"
	ERROR_FLAG_OWNER_ROLE_MUST_BE_TRANSFERED            = "ERROR_FLAG_OWNER_ROLE_MUST_BE_TRANSFERED"
	ERROR_FLAG_PASSWORD_INVALIED                        = "ERROR_FLAG_PASSWORD_INVALIED"
	ERROR_FLAG_TEAM_MUST_TRANSFERED_BEFORE_USER_SUSPEND = "ERROR_FLAG_TEAM_MUST_TRANSFERED_BEFORE_USER_SUSPEND"

	// can note create
	ERROR_FLAG_CAN_NOT_CREATE_USER            = "ERROR_FLAG_CAN_NOT_CREATE_USER"
	ERROR_FLAG_CAN_NOT_CREATE_TEAM            = "ERROR_FLAG_CAN_NOT_CREATE_TEAM"
	ERROR_FLAG_CAN_NOT_CREATE_TEAM_MEMBER     = "ERROR_FLAG_CAN_NOT_CREATE_TEAM_MEMBER"
	ERROR_FLAG_CAN_NOT_CREATE_INVITE          = "ERROR_FLAG_CAN_NOT_CREATE_INVITE"
	ERROR_FLAG_CAN_NOT_CREATE_INVITATION_CODE = "ERROR_FLAG_CAN_NOT_CREATE_INVITATION_CODE"
	ERROR_FLAG_CAN_NOT_CREATE_DOMAIN          = "ERROR_FLAG_CAN_NOT_CREATE_DOMAIN"
	ERROR_FLAG_CAN_NOT_CREATE_ACTION          = "ERROR_FLAG_CAN_NOT_CREATE_ACTION"
	ERROR_FLAG_CAN_NOT_CREATE_RESOURCE        = "ERROR_FLAG_CAN_NOT_CREATE_RESOURCE"
	ERROR_FLAG_CAN_NOT_CREATE_APP             = "ERROR_FLAG_CAN_NOT_CREATE_APP"
	ERROR_FLAG_CAN_NOT_CREATE_STATE           = "ERROR_FLAG_CAN_NOT_CREATE_STATE"
	ERROR_FLAG_CAN_NOT_CREATE_SNAPSHOT        = "ERROR_FLAG_CAN_NOT_CREATE_SNAPSHOT"
	ERROR_FLAG_CAN_NOT_CREATE_COMPONENT_TREE  = "ERROR_FLAG_CAN_NOT_CREATE_COMPONENT_TREE"

	// can not get resource
	ERROR_FLAG_CAN_NOT_GET_USER                = "ERROR_FLAG_CAN_NOT_GET_USER"
	ERROR_FLAG_CAN_NOT_GET_TEAM                = "ERROR_FLAG_CAN_NOT_GET_TEAM"
	ERROR_FLAG_CAN_NOT_GET_TEAM_MEMBER         = "ERROR_FLAG_CAN_NOT_GET_TEAM_MEMBER"
	ERROR_FLAG_CAN_NOT_GET_INVITE              = "ERROR_FLAG_CAN_NOT_GET_INVITE"
	ERROR_FLAG_CAN_NOT_GET_INVITATION_CODE     = "ERROR_FLAG_CAN_NOT_GET_INVITATION_CODE"
	ERROR_FLAG_CAN_NOT_GET_DOMAIN              = "ERROR_FLAG_CAN_NOT_GET_DOMAIN"
	ERROR_FLAG_CAN_NOT_GET_ACTION              = "ERROR_FLAG_CAN_NOT_GET_ACTION"
	ERROR_FLAG_CAN_NOT_GET_RESOURCE            = "ERROR_FLAG_CAN_NOT_GET_RESOURCE"
	ERROR_FLAG_CAN_NOT_GET_RESOURCE_META_INFO  = "ERROR_FLAG_CAN_NOT_GET_RESOURCE_META_INFO"
	ERROR_FLAG_CAN_NOT_GET_APP                 = "ERROR_FLAG_CAN_NOT_GET_APP"
	ERROR_FLAG_CAN_NOT_GET_BUILDER_DESCRIPTION = "ERROR_FLAG_CAN_NOT_GET_BUILDER_DESCRIPTION"
	ERROR_FLAG_CAN_NOT_GET_STATE               = "ERROR_FLAG_CAN_NOT_GET_STATE"
	ERROR_FLAG_CAN_NOT_GET_SNAPSHOT            = "ERROR_FLAG_CAN_NOT_GET_SNAPSHOT"

	// can not update resource
	ERROR_FLAG_CAN_NOT_UPDATE_USER            = "ERROR_FLAG_CAN_NOT_UPDATE_USER"
	ERROR_FLAG_CAN_NOT_UPDATE_TEAM            = "ERROR_FLAG_CAN_NOT_UPDATE_TEAM"
	ERROR_FLAG_CAN_NOT_UPDATE_TEAM_MEMBER     = "ERROR_FLAG_CAN_NOT_UPDATE_TEAM_MEMBER"
	ERROR_FLAG_CAN_NOT_UPDATE_INVITE          = "ERROR_FLAG_CAN_NOT_UPDATE_INVITE"
	ERROR_FLAG_CAN_NOT_UPDATE_INVITATION_CODE = "ERROR_FLAG_CAN_NOT_UPDATE_INVITATION_CODE"
	ERROR_FLAG_CAN_NOT_UPDATE_DOMAIN          = "ERROR_FLAG_CAN_NOT_UPDATE_DOMAIN"
	ERROR_FLAG_CAN_NOT_UPDATE_ACTION          = "ERROR_FLAG_CAN_NOT_UPDATE_ACTION"
	ERROR_FLAG_CAN_NOT_UPDATE_RESOURCE        = "ERROR_FLAG_CAN_NOT_UPDATE_RESOURCE"
	ERROR_FLAG_CAN_NOT_UPDATE_APP             = "ERROR_FLAG_CAN_NOT_UPDATE_APP"
	ERROR_FLAG_CAN_NOT_UPDATE_TREE_STATE      = "ERROR_FLAG_CAN_NOT_UPDATE_TREE_STATE"
	ERROR_FLAG_CAN_NOT_UPDATE_SNAPSHOT        = "ERROR_FLAG_CAN_NOT_UPDATE_SNAPSHOT"

	// can not delete
	ERROR_FLAG_CAN_NOT_DELETE_USER            = "ERROR_FLAG_CAN_NOT_DELETE_USER"
	ERROR_FLAG_CAN_NOT_DELETE_TEAM            = "ERROR_FLAG_CAN_NOT_DELETE_TEAM"
	ERROR_FLAG_CAN_NOT_DELETE_TEAM_MEMBER     = "ERROR_FLAG_CAN_NOT_DELETE_TEAM_MEMBER"
	ERROR_FLAG_CAN_NOT_DELETE_INVITE          = "ERROR_FLAG_CAN_NOT_DELETE_INVITE"
	ERROR_FLAG_CAN_NOT_DELETE_INVITATION_CODE = "ERROR_FLAG_CAN_NOT_DELETE_INVITATION_CODE"
	ERROR_FLAG_CAN_NOT_DELETE_DOMAIN          = "ERROR_FLAG_CAN_NOT_DELETE_DOMAIN"
	ERROR_FLAG_CAN_NOT_DELETE_ACTION          = "ERROR_FLAG_CAN_NOT_DELETE_ACTION"
	ERROR_FLAG_CAN_NOT_DELETE_RESOURCE        = "ERROR_FLAG_CAN_NOT_DELETE_RESOURCE"
	ERROR_FLAG_CAN_NOT_DELETE_APP             = "ERROR_FLAG_CAN_NOT_DELETE_APP"

	// can not other operation
	ERROR_FLAG_CAN_NOT_CHECK_TEAM_MEMBER        = "ERROR_FLAG_CAN_NOT_CHECK_TEAM_MEMBER"
	ERROR_FLAG_CAN_NOT_DUPLICATE_APP            = "ERROR_FLAG_CAN_NOT_DUPLICATE_APP"
	ERROR_FLAG_CAN_NOT_RELEASE_APP              = "ERROR_FLAG_CAN_NOT_RELEASE_APP"
	ERROR_FLAG_CAN_NOT_TEST_RESOURCE_CONNECTION = "ERROR_FLAG_CAN_NOT_TEST_RESOURCE_CONNECTION"

	// permission failed
	ERROR_FLAG_ACCESS_DENIED                  = "ERROR_FLAG_ACCESS_DENIED"
	ERROR_FLAG_TEAM_CLOSED_THE_PERMISSION     = "ERROR_FLAG_TEAM_CLOSED_THE_PERMISSION"
	ERROR_FLAG_EMAIL_ALREADY_USED             = "ERROR_FLAG_EMAIL_ALREADY_USED"
	ERROR_FLAG_EMAIL_HAS_BEEN_TAKEN           = "ERROR_FLAG_EMAIL_HAS_BEEN_TAKEN"
	ERROR_FLAG_INVITATION_CODE_ALREADY_USED   = "ERROR_FLAG_INVITATION_CODE_ALREADY_USED"
	ERROR_FLAG_INVITATION_LINK_UNAVALIABLE    = "ERROR_FLAG_INVITATION_LINK_UNAVALIABLE"
	ERROR_FLAG_TEAM_IDENTIFIER_HAS_BEEN_TAKEN = "ERROR_FLAG_TEAM_IDENTIFIER_HAS_BEEN_TAKEN"
	ERROR_FLAG_USER_ALREADY_JOINED_TEAM       = "ERROR_FLAG_USER_ALREADY_JOINED_TEAM"
	ERROR_FLAG_SIGN_IN_FAILED                 = "ERROR_FLAG_SIGN_IN_FAILED"
	ERROR_FLAG_NO_SUCH_USER                   = "ERROR_FLAG_NO_SUCH_USER"

	// call resource failed
	ERROR_FLAG_SEND_EMAIL_FAILED             = "ERROR_FLAG_SEND_EMAIL_FAILED"
	ERROR_FLAG_SEND_VERIFICATION_CODE_FAILED = "ERROR_FLAG_SEND_VERIFICATION_CODE_FAILED"
	ERROR_FLAG_CREATE_LINK_FAILED            = "ERROR_FLAG_CREATE_LINK_FAILED"
	ERROR_FLAG_CREATE_UPLOAD_URL_FAILED      = "ERROR_FLAG_CREATE_UPLOAD_URL_FAILED"
	ERROR_FLAG_EXECUTE_ACTION_FAILED         = "ERROR_FLAG_EXECUTE_ACTION_FAILED"
	ERROR_FLAG_GENERATE_SQL_FAILED           = "ERROR_FLAG_GENERATE_SQL_FAILED"
	ERROR_FLAG_UPLOAD_IMAGE_FAILED           = "ERROR_FLAG_UPLOAD_IMAGE_FAILED"
	ERROR_FLAG_CREATE_TABLE_FAILED           = "ERROR_FLAG_CREATE_TABLE_FAILED"
	ERROR_FLAG_RECORD_TYPE_ERROR             = "ERROR_FLAG_RECORD_TYPE_ERROR"
	ERROR_FLAG_INSERT_RECORD_FAILED          = "ERROR_FLAG_INSERT_RECORD_FAILED"
	ERROR_FLAG_DISPLAY_RECORD_FAILED         = "ERROR_FLAG_DISPLAY_RECORD_FAILED"
	// internal failed
	ERROR_FLAG_BUILD_TEAM_MEMBER_LIST_FAILED = "ERROR_FLAG_BUILD_TEAM_MEMBER_LIST_FAILED"
	ERROR_FLAG_BUILD_TEAM_CONFIG_FAILED      = "ERROR_FLAG_BUILD_TEAM_CONFIG_FAILED"
	ERROR_FLAG_BUILD_TEAM_PERMISSION_FAILED  = "ERROR_FLAG_BUILD_TEAM_PERMISSION_FAILED"
	ERROR_FLAG_BUILD_USER_INFO_FAILED        = "ERROR_FLAG_BUILD_USER_INFO_FAILED"
	ERROR_FLAG_BUILD_APP_CONFIG_FAILED       = "ERROR_FLAG_BUILD_APP_CONFIG_FAILED"
	ERROR_FLAG_GENERATE_PASSWORD_FAILED      = "ERROR_FLAG_GENERATE_PASSWORD_FAILED"

	// google sheets oauth2 failed
	ERROR_FLAG_CAN_NOT_CREATE_TOKEN               = "ERROR_FLAG_CAN_NOT_CREATE_TOKEN"
	ERROR_FLAG_CAN_NOT_AUTHORIZE_GOOGLE_SHEETS    = "ERROR_FLAG_CAN_NOT_AUTHORIZE_GOOGLE_SHEETS"
	ERROR_FLAG_CAN_NOT_REFRESH_GOOGLE_SHEETS      = "ERROR_FLAG_CAN_NOT_REFRESH_GOOGLE_SHEETS"
	ERROR_FLAG_CAN_NOT_FORK_APP                   = "ERROR_FLAG_CAN_NOT_FORK_APP"
	ERROR_FLAG_CAN_NOT_PUBLISH_APP_TO_MARKETPLACE = "ERROR_FLAG_CAN_NOT_PUBLISH_APP_TO_MARKETPLACE"
	ERROR_FLAG_CAN_NOT_GET_TOKEN                  = "ERROR_FLAG_CAN_NOT_GET_TOKEN"
	ERROR_FLAG_CAN_NOT_REFRESH_TOKEN              = "ERROR_FLAG_CAN_NOT_REFRESH_TOKEN"

	ERROR_FLAG_SQL_INVALID = "ERROR_FLAG_SQL_INVALID"
)

var SKIPPING_MAGIC_ID = map[string]int{
	"0":  0,
	"-1": -1,
	"-2": -2,
	"-3": -3,
}

func (controller *Controller) GetStringFromFormData(c *gin.Context, paramName string) (string, error) {
	// get request param
	paramValue := c.PostFormArray(paramName)

	// ho hit, convert
	if len(paramValue) == 0 {
		controller.FeedbackBadRequest(c, ERROR_FLAG_VALIDATE_REQUEST_PARAM_FAILED, "please input param for request.")
		return "", errors.New("input missing " + paramName + " field.")
	}
	return paramValue[0], nil
}

func (controller *Controller) GetOptionalStringFromFormData(c *gin.Context, paramName string) string {
	// get request param
	paramValue := c.PostFormArray(paramName)

	// ho hit, convert
	if len(paramValue) == 0 {
		return ""
	}
	return paramValue[0]
}

// GetMagicIntParamFromRequest
// @receiver controller
// @param c
// @param paramName
// @return int
// @return error
func (controller *Controller) GetMagicIntParamFromRequest(c *gin.Context, paramName string) (int, error) {
	// get request param
	paramValue := c.Param(paramName)
	// check skipping id
	if intID, hitSkippingID := SKIPPING_MAGIC_ID[paramValue]; hitSkippingID {
		return intID, nil
	}
	// ho hit, convert
	if len(paramValue) == 0 {
		controller.FeedbackBadRequest(c, ERROR_FLAG_VALIDATE_REQUEST_PARAM_FAILED, "please input param for request.")
		return 0, errors.New("input missing " + paramName + " field.")
	}
	paramValueInt := idconvertor.ConvertStringToInt(paramValue)
	return paramValueInt, nil
}

// TestMagicIntParamFromRequest test if Magic int exists in param, if not ,return 0 and an error.
// @receiver controller
// @param c
// @param paramName
// @return int
// @return error
func (controller *Controller) TestMagicIntParamFromRequest(c *gin.Context, paramName string) (int, error) {
	// get request param
	paramValue := c.Param(paramName)
	if len(paramValue) == 0 {
		return 0, errors.New("input missing " + paramName + " field.")
	}
	paramValueInt := idconvertor.ConvertStringToInt(paramValue)
	return paramValueInt, nil
}

// GetIntParamFromRequest
// @receiver controller
// @param c
// @param paramName
// @return int
// @return error
func (controller *Controller) GetIntParamFromRequest(c *gin.Context, paramName string) (int, error) {
	// get request param
	paramValue := c.Param(paramName)
	if len(paramValue) == 0 {
		controller.FeedbackBadRequest(c, ERROR_FLAG_VALIDATE_REQUEST_PARAM_FAILED, "please input param for request.")
		return 0, errors.New("input missing " + paramName + " field.")
	}
	paramValueInt, okAssert := strconv.Atoi(paramValue)
	if okAssert != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_VALIDATE_REQUEST_PARAM_FAILED, "please input param in int format.")
		return 0, errors.New("input teamID in wrong format.")
	}
	return paramValueInt, nil
}

// GetStringParamFromRequest
// @receiver controller
// @param c
// @param paramName
// @return string
// @return error
func (controller *Controller) GetStringParamFromRequest(c *gin.Context, paramName string) (string, error) {
	// get request param
	paramValue := c.Param(paramName)
	if len(paramValue) == 0 {
		controller.FeedbackBadRequest(c, ERROR_FLAG_VALIDATE_REQUEST_PARAM_FAILED, "please input param for request.")
		return "", errors.New("input missing " + paramName + " field.")
	}
	return paramValue, nil
}

// TestStringParamFromRequest
// @receiver controller
// @param c
// @param paramName
// @return string
// @return error
func (controller *Controller) TestStringParamFromRequest(c *gin.Context, paramName string) (string, error) {
	// get request param
	paramValue := c.Param(paramName)
	if len(paramValue) == 0 {
		return "", errors.New("input missing " + paramName + " field.")
	}
	return paramValue, nil
}

// TestFirstStringParamValueFromURI
// @receiver controller
// @param c
// @param paramName
// @return string
// @return error
func (controller *Controller) TestFirstStringParamValueFromURI(c *gin.Context, paramName string) (string, error) {
	valueMaps := c.Request.URL.Query()
	paramValues, hit := valueMaps[paramName]
	// get request param
	if !hit {
		return "", errors.New("input missing " + paramName + " field.")
	}
	return paramValues[0], nil
}

// GetFirstStringParamValueFromURI
// @receiver controller
// @param c
// @param paramName
// @return string
// @return error
func (controller *Controller) GetFirstStringParamValueFromURI(c *gin.Context, paramName string) (string, error) {
	valueMaps := c.Request.URL.Query()
	paramValues, hit := valueMaps[paramName]
	// get request param
	if !hit {
		controller.FeedbackBadRequest(c, ERROR_FLAG_VALIDATE_REQUEST_PARAM_FAILED, "please input param for request.")
		return "", errors.New("input missing " + paramName + " field.")
	}
	return paramValues[0], nil
}

// GetStringParamValuesFromURI
// @receiver controller
// @param c
// @param paramName
// @return []string
// @return error
func (controller *Controller) GetStringParamValuesFromURI(c *gin.Context, paramName string) ([]string, error) {
	valueMaps := c.Request.URL.Query()
	paramValues, hit := valueMaps[paramName]
	// get request param
	if !hit {
		controller.FeedbackBadRequest(c, ERROR_FLAG_VALIDATE_REQUEST_PARAM_FAILED, "please input param for request.")
		return nil, errors.New("input missing " + paramName + " field.")
	}
	return paramValues, nil
}

// GetStringParamFromHeader
// @receiver controller
// @param c
// @param paramName
// @return string
// @return error
func (controller *Controller) GetStringParamFromHeader(c *gin.Context, paramName string) (string, error) {
	paramValue := c.Request.Header[paramName]
	var ret string
	if len(paramValue) != 1 {
		controller.FeedbackBadRequest(c, ERROR_FLAG_VALIDATE_REQUEST_PARAM_FAILED, "can not fetch param from header.")
		return "", errors.New("can not fetch param from header.")
	} else {
		ret = paramValue[0]
	}
	return ret, nil
}

// GetUserIDFromAuth @note: this param was setted by authenticator.JWTAuth() method
// @receiver controller
// @param c
// @return int
// @return error
func (controller *Controller) GetUserIDFromAuth(c *gin.Context) (int, error) {
	// get request param
	userID, ok := c.Get("userID")
	if !ok {
		controller.FeedbackBadRequest(c, ERROR_FLAG_VALIDATE_REQUEST_TOKEN_FAILED, "auth token invalied, can not fetch user ID in it.")
		return 0, errors.New("input missing userID field.")
	}
	userIDInt, okAssert := userID.(int)
	if !okAssert {
		controller.FeedbackBadRequest(c, ERROR_FLAG_VALIDATE_REQUEST_TOKEN_FAILED, "auth token invalied,user ID is not int type in it.")
		return 0, errors.New("input userID in wrong format.")
	}
	return userIDInt, nil
}

// FeedbackOK
// @receiver controller
// @param c
// @param resp
func (controller *Controller) FeedbackOK(c *gin.Context, resp response.Response) {
	if resp != nil {
		c.JSON(http.StatusOK, gin.H{
			"Code":    200,
			"Flag":    SUCCESS_FLAG,
			"Message": "success",
			"data":    resp.ExportForFeedback(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Code":    200,
		"Flag":    SUCCESS_FLAG,
		"Message": "success",
		"data":    "",
	})

}

// FeedbackCreated
// @receiver controller
// @param c
// @param resp
func (controller *Controller) FeedbackCreated(c *gin.Context, resp response.Response) {
	if resp != nil {
		c.JSON(http.StatusCreated, gin.H{
			"Code":    201,
			"Flag":    SUCCESS_FLAG,
			"Message": "success",
			"data":    resp.ExportForFeedback(),
		})
		return
	}
	// HTTP 201 with empty response
	c.JSON(http.StatusCreated, nil)
}

// FeedbackBadRequest
// @receiver controller
// @param c
// @param Flag
// @param Message
func (controller *Controller) FeedbackBadRequest(c *gin.Context, Flag string, Message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"Code":    400,
		"Flag":    Flag,
		"Message": Message,
		"data":    "",
	})
	return
}

// FeedbackBadRequest
// @receiver controller
// @param c
// @param Flag
// @param Message
func (controller *Controller) FeedbackBadRequestWithResponse(c *gin.Context, Flag string, Message string, res response.Response) {
	c.JSON(http.StatusBadRequest, gin.H{
		"Code":    400,
		"Flag":    Flag,
		"Message": Message,
		"data":    res.ExportForFeedback(),
	})
	return
}

// FeedbackRedirect
// @receiver controller
// @param c
// @param uri
func (controller *Controller) FeedbackRedirect(c *gin.Context, uri string) {
	c.Redirect(302, uri)
	return
}

// FeedbackInternalServerError
// @receiver controller
// @param c
// @param Flag
// @param Message
func (controller *Controller) FeedbackInternalServerError(c *gin.Context, Flag string, Message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"Code":    500,
		"Flag":    Flag,
		"Message": Message,
		"data":    "",
	})
	return
}

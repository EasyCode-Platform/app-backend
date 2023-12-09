package controller

import (
	"github.com/EasyCode-Platform/app-backend/src/response"
	"github.com/gin-gonic/gin"
)

func (controller *Controller) UploadImage(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_PARSE_REQUEST_FILE_FAILED, "Failed to get file from request")
		return
	}
	// 获取所有图片
	files := form.File["image-file"]
	// 因为没有事务操作，所以这里要统计成功插入的数量
	count := 0
	// 遍历所有图片
	for _, file := range files {
		//TODO重复的同名文件提示
		if err != nil {
			controller.FeedbackBadRequestWithResponse(c, ERROR_FLAG_PARSE_REQUEST_FILE_FAILED, "Failed to get file from request", response.NewFileResponse(count))
		}
		if err := controller.Storage.ImageStorage.UploadImage(file); err != nil {
			controller.FeedbackBadRequestWithResponse(c, ERROR_FLAG_UPLOAD_IMAGE_FAILED, "Failed to upload image", response.NewFileResponse(count))
		}
		count++
	}
	controller.FeedbackOK(c, response.NewFileResponse(count))

}

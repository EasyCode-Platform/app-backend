package response

type FileResponse struct {
	FileUploaded int
}

func NewFileResponse(num int) *FileResponse {
	return &FileResponse{
		FileUploaded: num,
	}
}

func (res *FileResponse) ExportForFeedback() interface{} {
	return res.FileUploaded
}

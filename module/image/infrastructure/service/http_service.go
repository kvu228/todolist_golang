package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"to_do_list/module/image/usecase"
)

type HttpImageService interface {
	handleUploadImage() gin.HandlerFunc
	Routes(group *gin.RouterGroup)
}

type httpService struct {
	uploader     usecase.Uploader
	imageUseCase usecase.ImageUseCase
}

func NewHttpImageService(uploader usecase.Uploader, imageUseCase usecase.ImageUseCase) HttpImageService {
	return &httpService{uploader: uploader, imageUseCase: imageUseCase}
}

func (s *httpService) handleUploadImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		f, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		file, err := f.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		defer file.Close()

		fileData := make([]byte, f.Size)
		if _, err := file.Read(fileData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		dto := usecase.UploadDTO{
			Name:     c.PostForm("name"),
			FileName: f.Filename,
			// FileType: f.Header.Get("Content-Type"),  --> this method will get value from http Header
			FileType: http.DetectContentType(fileData),
			FileSize: int(f.Size),
			FileData: fileData,
		}

		image, err := s.imageUseCase.UploadImage(c.Request.Context(), dto)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		image.SetCDNDomain(s.uploader.GetDomain())

		c.JSON(http.StatusOK, gin.H{"data": image})
	}
}

func (s *httpService) Routes(group *gin.RouterGroup) {
	group.POST("/upload-image", s.handleUploadImage())
}

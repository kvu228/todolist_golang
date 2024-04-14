package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"to_do_list/module/image/infrastructure/repository/mysql"
	"to_do_list/module/image/usecase"
)

type httpService struct {
	uploader usecase.Uploader
	repo     mysql.ImageRepository
}

func NewHttpService(uploader usecase.Uploader, repo mysql.ImageRepository) *httpService {
	return &httpService{uploader: uploader, repo: repo}
}

func (s *httpService) handleUploadIamge() gin.HandlerFunc {
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
		uc := usecase.NewUseCase(s.uploader, s.repo)

		image, err := uc.UploadImage(c.Request.Context(), dto)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		image.SetCDNDomain(s.uploader.GetDomain())

		c.JSON(http.StatusOK, gin.H{"data": image})
	}
}

func (s httpService) Routes(group *gin.RouterGroup) {
	group.POST("/upload-image", s.handleUploadIamge())
}

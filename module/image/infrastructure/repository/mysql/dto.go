package mysql

import (
	"github.com/google/uuid"
	"time"
	"to_do_list/module/image/domain"
)

type ImageDTO struct {
	Id              uuid.UUID `json:"id"`
	Title           string    `json:"title"`
	FileName        string    `json:"file_name"`
	FileURL         string    `json:"file_url" gorm:"-"`
	FileSize        int       `json:"file_size"`
	FileType        string    `json:"file_type"`
	StorageProvider string    `json:"storage_provider"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UploadedAt      time.Time `json:"uploaded_at"`
}

func (dto *ImageDTO) ToEntity() (image *domain.Image, err error) {
	return domain.NewImage(
		dto.Id,
		dto.Title,
		dto.FileName,
		dto.FileSize,
		dto.FileType,
		dto.StorageProvider,
		dto.Status,
		dto.CreatedAt,
		dto.UploadedAt,
	), nil
}

package domain

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Image struct {
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

func NewImage(id uuid.UUID, title string, fileName string, fileSize int, fileType string, storageProvider string, status string, createdAt time.Time, uploadedAt time.Time) *Image {
	return &Image{Id: id, Title: title, FileName: fileName, FileSize: fileSize, FileType: fileType, StorageProvider: storageProvider, Status: status, CreatedAt: createdAt, UploadedAt: uploadedAt}
}

func (i *Image) SetCDNDomain(domain string) {
	i.FileURL = fmt.Sprintf("%s/%s", domain, i.FileName)
}

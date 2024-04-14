package mysql

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"to_do_list/common"
	"to_do_list/module/image/domain"
)

type imageMySQLRepo struct {
	db *gorm.DB
}

func (i imageMySQLRepo) Find(ctx context.Context, id uuid.UUID) (*common.Image, error) {
	var image common.Image
	if err := i.db.Table(common.TbNameImages).Where("id =?", id).First(&image).Error; err != nil {
		return nil, err
	}
	return &common.Image{
		Id:              image.Id,
		Title:           image.Title,
		FileName:        image.FileName,
		FileURL:         image.FileURL,
		FileSize:        image.FileSize,
		FileType:        image.FileType,
		StorageProvider: image.StorageProvider,
		Status:          image.Status,
	}, nil
}

func NewImageMySQLRepo(db *gorm.DB) ImageRepository {
	return &imageMySQLRepo{db: db}
}

func (i imageMySQLRepo) Create(ctx context.Context, entity *domain.Image) error {
	return i.db.Table(common.TbNameImages).Create(&entity).Error
}

func (i imageMySQLRepo) SetImageStatusActivated(ctx context.Context, id uuid.UUID) error {
	return i.db.Table(common.TbNameImages).Where("id = ?", id).Updates(domain.Image{Status: domain.StatusActivated}).Error
}

type ImageRepository interface {
	ImageQueryRepository
	ImageCmdRepository
}

type ImageCmdRepository interface {
	Create(ctx context.Context, entity *domain.Image) error
	SetImageStatusActivated(ctx context.Context, id uuid.UUID) error
}

type ImageQueryRepository interface {
	Find(ctx context.Context, id uuid.UUID) (*common.Image, error)
}

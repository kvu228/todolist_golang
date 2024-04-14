package usecase

import (
	"context"
	"fmt"
	"time"
	"to_do_list/common"
	"to_do_list/module/image/domain"
)

type UseCase interface {
	UploadImage(ctx context.Context, dto UploadDTO) (*domain.Image, error)
}

type useCase struct {
	uploader Uploader
	repo     CmdRepository
}

func NewUseCase(uploader Uploader, repo CmdRepository) UseCase {
	return useCase{uploader: uploader, repo: repo}
}

func (uc useCase) UploadImage(ctx context.Context, dto UploadDTO) (*domain.Image, error) {
	dstFileName := fmt.Sprintf("%d_%s", time.Now().UTC().UnixNano(), dto.FileName)
	if err := uc.uploader.SaveFileUploaded(ctx, dto.FileData, dstFileName); err != nil {
		return nil, domain.ErrCannotUploadImage
	}

	imageId, _ := common.GenUUID()
	image := domain.NewImage(
		imageId,
		dto.Name,
		dstFileName,
		dto.FileSize,
		dto.FileType,
		uc.uploader.GetName(),
		domain.StatusUploaded,
		time.Now().UTC(),
		time.Now().UTC(),
	)
	if err := uc.repo.Create(ctx, image); err != nil {

		return nil, domain.ErrCannotUploadImage
	}

	return image, nil
}

type Uploader interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) error
	GetName() string
	GetDomain() string
}

type CmdRepository interface {
	Create(ctx context.Context, entity *domain.Image) error
}

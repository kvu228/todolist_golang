package usecase

import (
	"context"
	"fmt"
	"time"
	"to_do_list/common"
	"to_do_list/module/image/domain"
)

type ImageUseCase interface {
	UploadImage(ctx context.Context, dto UploadDTO) (*domain.Image, error)
}

type imageUseCase struct {
	uploader           Uploader
	imageCmdRepository ImageCmdRepository
}

func NewImageUseCase(uploader Uploader, imageCmdRepository ImageCmdRepository) ImageUseCase {
	return &imageUseCase{uploader: uploader, imageCmdRepository: imageCmdRepository}
}

func (uc *imageUseCase) UploadImage(ctx context.Context, dto UploadDTO) (*domain.Image, error) {
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
	if err := uc.imageCmdRepository.Create(ctx, image); err != nil {

		return nil, domain.ErrCannotUploadImage
	}

	return image, nil
}

type Uploader interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) error
	GetName() string
	GetDomain() string
}

type ImageCmdRepository interface {
	Create(ctx context.Context, entity *domain.Image) error
}

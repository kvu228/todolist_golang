package command

import (
	"context"
	"to_do_list/common"
	"to_do_list/module/post/domain"
	"to_do_list/module/post/usecase"
)

type CreatePostUseCase interface {
	CreatePost(ctx context.Context, dto usecase.NewPostDTO) error
}

type createPostUseCase struct {
	postCommandRepository PostCmdRepository
}

func NewCreatePostUseCase(postCmdRepository PostCmdRepository) CreatePostUseCase {
	return &createPostUseCase{postCommandRepository: postCmdRepository}
}
func (uc *createPostUseCase) CreatePost(ctx context.Context, dto usecase.NewPostDTO) error {
	id, _ := common.GenUUID()
	postEntity := domain.NewPost(
		id,
		dto.Title,
		dto.Body,
		"doing",
		dto.OwnerId,
	)
	return uc.postCommandRepository.CreatePost(ctx, postEntity)
}

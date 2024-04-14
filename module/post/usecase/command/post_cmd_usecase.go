package command

import (
	"context"
	"to_do_list/module/post/domain"
)

type PostCmdUseCase interface {
	CreatePostUseCase
}

type postCmdUseCase struct {
	CreatePostUseCase
}

func NewPostCmdUseCase(postCmdRepository PostCmdRepository) PostCmdUseCase {
	return &postCmdUseCase{
		CreatePostUseCase: NewCreatePostUseCase(postCmdRepository),
	}
}

type PostCmdRepository interface {
	CreatePost(ctx context.Context, post *domain.Post) error
}

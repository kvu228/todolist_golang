package query

import (
	"context"
	"github.com/google/uuid"
	"to_do_list/module/post/usecase"
)

type PostQueryUseCase interface {
	ListPostsUseCase
}

type postQueryUseCase struct {
	ListPostsUseCase
}

func NewPostQueryUseCase(postQueryRepository PostQueryRepository, userQueryRepository UserQueryRepository) PostQueryUseCase {
	return &postQueryUseCase{ListPostsUseCase: NewListPostsQueryUseCase(postQueryRepository, userQueryRepository)}
}

type PostQueryRepository interface {
	FindWithIds(ctx context.Context, ids []uuid.UUID) (posts []usecase.PostDTO, err error)
	FindWithParams(ctx context.Context, params *usecase.ListPostsParams) (posts []usecase.PostDTO, err error)
}

type UserQueryRepository interface {
	FindWithIds(ctx context.Context, ids []uuid.UUID) (owners []usecase.OwnerDTO, err error)
}

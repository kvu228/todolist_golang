package query

import (
	"context"
	"github.com/google/uuid"
	"to_do_list/module/users/domain"
)

type GetUserUseCase interface {
	GetUser(ctx context.Context, id uuid.UUID) (*domain.UserDTO, error)
}

type getUserUseCase struct {
	userQueryRepo UserQueryRepository
}

func NewGetUserUseCase(userQueryRepo UserQueryRepository) *getUserUseCase {
	return &getUserUseCase{userQueryRepo: userQueryRepo}
}

func (uc *getUserUseCase) GetUser(ctx context.Context, id uuid.UUID) (user *domain.UserDTO, err error) {
	user, err = uc.userQueryRepo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, err
}

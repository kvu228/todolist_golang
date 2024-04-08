package query

import (
	"context"
	"github.com/google/uuid"
	"to_do_list/module/users/infrastructure/repositories/mysql"
)

type GetUserUseCase interface {
	GetUser(ctx context.Context, id uuid.UUID) (*mysql.UserDTO, error)
}

type getUserUseCase struct {
	userQueryRepo UserQueryRepository
}

func NewGetUserUseCase(userQueryRepo UserQueryRepository) GetUserUseCase {
	return &getUserUseCase{userQueryRepo: userQueryRepo}
}

func (uc *getUserUseCase) GetUser(ctx context.Context, id uuid.UUID) (user *mysql.UserDTO, err error) {
	user, err = uc.userQueryRepo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, err
}

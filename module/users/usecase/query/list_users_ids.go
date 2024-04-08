package query

import (
	"context"
	"github.com/google/uuid"
	"to_do_list/module/users/infrastructure/repositories/mysql"
)

type ListUserByIdsUseCase interface {
	ListUsersByIds(ctx context.Context, ids []uuid.UUID) ([]*mysql.UserDTO, error)
}

type listUserByIdsUseCase struct {
	userQueryRepo UserQueryRepository
}

func NewListUserByIdsUseCase(userQueryRepo UserQueryRepository) ListUserByIdsUseCase {
	return &listUserByIdsUseCase{userQueryRepo: userQueryRepo}
}

func (uc *listUserByIdsUseCase) ListUsersByIds(ctx context.Context, ids []uuid.UUID) (users []*mysql.UserDTO, err error) {

	users, err = uc.userQueryRepo.FindByIds(ctx, ids)
	if err != nil {
		return nil, err
	}

	return users, nil

}

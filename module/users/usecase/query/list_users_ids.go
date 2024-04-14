package query

import (
	"context"
	"github.com/google/uuid"
	"to_do_list/module/users/usecase"
)

type ListUserByIdsUseCase interface {
	ListUsersByIds(ctx context.Context, ids []uuid.UUID) ([]*usecase.UserDTO, error)
}

type listUserByIdsUseCase struct {
	userQueryRepo UserQueryRepository
}

func NewListUserByIdsUseCase(userQueryRepo UserQueryRepository) ListUserByIdsUseCase {
	return &listUserByIdsUseCase{userQueryRepo: userQueryRepo}
}

func (uc *listUserByIdsUseCase) ListUsersByIds(ctx context.Context, ids []uuid.UUID) (users []*usecase.UserDTO, err error) {

	userEntities, err := uc.userQueryRepo.FindByIds(ctx, ids)
	if err != nil {
		return nil, err
	}
	users = make([]*usecase.UserDTO, len(userEntities))
	for i, userEntity := range userEntities {
		users[i] = &usecase.UserDTO{
			Id:        userEntity.Id(),
			FirstName: userEntity.FirstName(),
			LastName:  userEntity.LastName(),
			Email:     userEntity.Email(),
			Status:    userEntity.Status(),
			Role:      userEntity.Role(),
			Avatar:    userEntity.Avatar(),
			CreatedAt: userEntity.CreatedAt(),
			UpdatedAt: userEntity.UpdatedAt(),
		}
	}
	return users, nil

}

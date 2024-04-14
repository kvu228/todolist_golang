package query

import (
	"context"
	"github.com/google/uuid"
	"to_do_list/module/users/usecase"
)

type GetUserUseCase interface {
	GetUser(ctx context.Context, id uuid.UUID) (*usecase.UserDTO, error)
}

type getUserUseCase struct {
	userQueryRepo UserQueryRepository
}

func NewGetUserUseCase(userQueryRepo UserQueryRepository) GetUserUseCase {
	return &getUserUseCase{userQueryRepo: userQueryRepo}
}

func (uc *getUserUseCase) GetUser(ctx context.Context, id uuid.UUID) (user *usecase.UserDTO, err error) {
	userEntity, err := uc.userQueryRepo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	user = &usecase.UserDTO{
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
	return user, err
}

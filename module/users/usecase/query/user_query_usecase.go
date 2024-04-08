package query

import (
	"context"
	"github.com/google/uuid"
	"to_do_list/module/users/domain"
)

type UserQueryUseCase interface {
	GetUser(ctx context.Context, id uuid.UUID) (*domain.UserDTO, error)
	ListUsersByIds(ctx context.Context, ids []uuid.UUID) ([]*domain.UserDTO, error)
}

type userQueryUseCase struct {
	*listUserByIdsUseCase
	*getUserUseCase
}

func NewUserQueryUseCase(userQueryRepository UserQueryRepository) *userQueryUseCase {
	return &userQueryUseCase{
		listUserByIdsUseCase: NewListUserByIdsUseCase(userQueryRepository),
		getUserUseCase:       NewGetUserUseCase(userQueryRepository)}
}

// Repositories interfaces
type UserQueryRepository interface {
	FindById(ctx context.Context, id uuid.UUID) (user *domain.UserDTO, err error)
	FindByEmail(ctx context.Context, email string) (user *domain.UserDTO, err error)
	FindByIds(ctx context.Context, ids []uuid.UUID) (uses []*domain.UserDTO, err error)
}

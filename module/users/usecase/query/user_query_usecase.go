package query

import (
	"context"
	"github.com/google/uuid"
	"to_do_list/module/users/infrastructure/repositories/mysql"
)

type UserQueryUseCase interface {
	GetUser(ctx context.Context, id uuid.UUID) (*mysql.UserDTO, error)
	ListUsersByIds(ctx context.Context, ids []uuid.UUID) ([]*mysql.UserDTO, error)
}

type userQueryUseCase struct {
	ListUserByIdsUseCase
	GetUserUseCase
}

func NewUserQueryUseCase(userQueryRepository UserQueryRepository) *userQueryUseCase {
	return &userQueryUseCase{
		ListUserByIdsUseCase: NewListUserByIdsUseCase(userQueryRepository),
		GetUserUseCase:       NewGetUserUseCase(userQueryRepository)}
}

// Repositories interfaces
type UserQueryRepository interface {
	FindById(ctx context.Context, id uuid.UUID) (user *mysql.UserDTO, err error)
	FindByEmail(ctx context.Context, email string) (user *mysql.UserDTO, err error)
	FindByIds(ctx context.Context, ids []uuid.UUID) (uses []*mysql.UserDTO, err error)
}

package query

import (
	"context"
	"github.com/google/uuid"
	"to_do_list/module/users/domain"
	"to_do_list/module/users/usecase"
)

type UserQueryUseCase interface {
	GetUser(ctx context.Context, id uuid.UUID) (*usecase.UserDTO, error)
	ListUsersByIds(ctx context.Context, ids []uuid.UUID) ([]*usecase.UserDTO, error)
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
	FindById(ctx context.Context, id uuid.UUID) (user *domain.User, err error)
	FindByEmail(ctx context.Context, email string) (user *domain.User, err error)
	FindByIds(ctx context.Context, ids []uuid.UUID) (uses []*domain.User, err error)
}

type SessionRepository interface {
	SessionQueryRepository
	SessionCmdRepository
}

type SessionQueryRepository interface {
	Find(ctx context.Context, id uuid.UUID) (session *domain.Session, err error)
	FindByRefreshToken(ctx context.Context, refreshToken string) (session *domain.Session, err error)
}

type SessionCmdRepository interface {
	Create(ctx context.Context, session *domain.Session) error
	Delete(ctx context.Context, id uuid.UUID) error
}

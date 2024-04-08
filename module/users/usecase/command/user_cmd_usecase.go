package command

import (
	"context"
	"github.com/google/uuid"
	"to_do_list/module/users/domain"
	"to_do_list/module/users/usecase/query"
)

type UserCmdUseCase interface {
	Register(ctx context.Context, dto domain.EmailPasswordRegistrationDTO) error
}

type userCmdUseCase struct {
	*registerUseCase
}

func NewUserCmdUseCase(userRepository UserRepository, hasher Hasher) *userCmdUseCase {
	return &userCmdUseCase{
		registerUseCase: NewRegisterUseCase(userRepository, userRepository, hasher),
	}
}

type UserRepository interface {
	query.UserQueryRepository
	UserCmdRepository
}

type UserCmdRepository interface {
	Create(ctx context.Context, user *domain.User) error
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}

package command

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"to_do_list/common"
	"to_do_list/module/users/domain"
	"to_do_list/module/users/usecase"
	"to_do_list/module/users/usecase/query"
)

type TokenProvider interface {
	IssueToken(ctx context.Context, id, sub string) (token string, err error)
	ParseToken(ctx context.Context, tokenString string) (claims *jwt.RegisteredClaims, err error)
	ExpireTokenInSeconds() int
	ExpireRefreshTokenInSeconds() int
}

type Hasher interface {
	RandomString(length int) (string, error)
	HashPassword(salt string, password string) (string, error)
	CheckPassword(salt string, password string, hashPassword string) bool
}

type UserCmdUseCase interface {
	Register(ctx context.Context, dto usecase.EmailPasswordRegistrationDTO) error
	LoginEmailPassword(ctx context.Context, dto usecase.EmailPasswordLoginDTO) (*usecase.TokenResponseDTO, error)
	ChangeAvatar(ctx context.Context, dto usecase.SetSingleImageDTO) error
	RefreshToken(ctx context.Context, refreshToken string) (*usecase.TokenResponseDTO, error)
}

type userCmdUseCase struct {
	RegisterUseCase
	LoginEmailPasswordUseCase
	ChangeAvatarUseCase
	RefreshTokenUseCase
}

func NewUserCmdUseCase(userRepository UserRepository, sessionCmdRepository SessionCmdRepository, imageRepository ImageRepository, tokenProvider TokenProvider, hasher Hasher) UserCmdUseCase {
	return &userCmdUseCase{
		RegisterUseCase:           NewRegisterUseCase(userRepository, userRepository, hasher),
		LoginEmailPasswordUseCase: NewLoginEmailPasswordUseCase(userRepository, sessionCmdRepository, tokenProvider, hasher),
		ChangeAvatarUseCase:       NewChangeAvatarUseCase(userRepository, imageRepository),
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

type SessionCmdRepository interface {
	Create(ctx context.Context, session *domain.Session) error
}

type ImageRepository interface {
	ImageCmdRepository
	ImageQueryRepository
}

type ImageCmdRepository interface {
	SetImageStatusActivated(ctx context.Context, id uuid.UUID) error
}

type ImageQueryRepository interface {
	Find(ctx context.Context, id uuid.UUID) (*common.Image, error)
}

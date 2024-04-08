package command

import (
	"context"
	"errors"
	"to_do_list/common"
	"to_do_list/module/users/domain"
	"to_do_list/module/users/usecase/query"
)

type Hasher interface {
	RandomString(length int) (string, error)
	HashPassword(salt string, password string) (string, error)
	CheckPassword(salt string, password string, hashPassword string) bool
}

type registerUseCase struct {
	userCommandRepository UserCmdRepository
	userQueryRepository   query.UserQueryRepository
	hasher                Hasher
}

func NewRegisterUseCase(userCommandRepository UserCmdRepository, userQueryRepository query.UserQueryRepository, hasher Hasher) *registerUseCase {
	return &registerUseCase{userCommandRepository: userCommandRepository, userQueryRepository: userQueryRepository, hasher: hasher}
}
func (uc *registerUseCase) Register(ctx context.Context, dto domain.EmailPasswordRegistrationDTO) error {

	user, err := uc.userQueryRepository.FindByEmail(ctx, dto.Email)
	if user != nil {
		return domain.ErrEmailExisted
	}
	if err != nil && !errors.Is(err, common.ErrRecordNotFound) {
		return domain.ErrCannotRegister
	}

	salt, _ := uc.hasher.RandomString(common.LenSalt)
	hashedPassword, err := uc.hasher.HashPassword(salt, dto.Password)
	if err != nil {
		return err
	}

	id, _ := common.GenUUID()
	userEntity := domain.NewUser(
		id,
		dto.FirstName,
		dto.LastName,
		dto.Email,
		hashedPassword,
		salt,
		"",
		"activated",
		"user",
	)

	if err := uc.userCommandRepository.Create(ctx, userEntity); err != nil {
		return err
	}
	return nil
}

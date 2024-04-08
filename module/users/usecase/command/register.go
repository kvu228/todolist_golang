package command

import (
	"context"
	"errors"
	"to_do_list/common"
	"to_do_list/module/users/domain"
	"to_do_list/module/users/usecase"
	"to_do_list/module/users/usecase/query"
)

type RegisterUseCase interface {
	Register(ctx context.Context, dto usecase.EmailPasswordRegistrationDTO) error
}

type registerUseCase struct {
	userCommandRepository UserCmdRepository
	userQueryRepository   query.UserQueryRepository
	hasher                Hasher
}

func NewRegisterUseCase(userCommandRepository UserCmdRepository, userQueryRepository query.UserQueryRepository, hasher Hasher) RegisterUseCase {
	return &registerUseCase{userCommandRepository: userCommandRepository, userQueryRepository: userQueryRepository, hasher: hasher}
}
func (uc *registerUseCase) Register(ctx context.Context, dto usecase.EmailPasswordRegistrationDTO) error {

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

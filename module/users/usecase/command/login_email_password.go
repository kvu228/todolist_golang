package command

import (
	"context"
	"time"
	"to_do_list/common"
	"to_do_list/module/users/domain"
	"to_do_list/module/users/usecase"
	"to_do_list/module/users/usecase/query"
)

type LoginEmailPasswordUseCase interface {
	LoginEmailPassword(ctx context.Context, dto usecase.EmailPasswordLoginDTO) (*usecase.TokenResponseDTO, error)
}

type loginEmailPasswordUseCase struct {
	userQueryRepository  query.UserQueryRepository
	sessionCmdRepository SessionCmdRepository
	tokenProvider        TokenProvider
	hasher               Hasher
}

func NewLoginEmailPasswordUseCase(userQueryRepository query.UserQueryRepository, sessionCmdRepository SessionCmdRepository, tokenProvider TokenProvider, hasher Hasher) LoginEmailPasswordUseCase {
	return &loginEmailPasswordUseCase{userQueryRepository: userQueryRepository, sessionCmdRepository: sessionCmdRepository, tokenProvider: tokenProvider, hasher: hasher}
}

func (uc *loginEmailPasswordUseCase) LoginEmailPassword(ctx context.Context, dto usecase.EmailPasswordLoginDTO) (*usecase.TokenResponseDTO, error) {
	// 1. Find user by email
	user, err := uc.userQueryRepository.FindByEmail(ctx, dto.Email)
	if err != nil {
		return nil, err
	}

	// 2. Hash and compare password
	if ok := uc.hasher.CheckPassword(user.Salt(), dto.Password, user.Password()); !ok {
		return nil, domain.ErrInvalidEmailPassword
	}
	userId := user.Id()
	sessionId, _ := common.GenUUID()

	// 3. Gen JWT Token
	accessToken, err := uc.tokenProvider.IssueToken(ctx, sessionId.String(), userId.String())
	if err != nil {
		return nil, err
	}
	// 4. Gen RefreshToken
	refreshToken, _ := uc.hasher.RandomString(common.LenRefreshToken)
	accessTokenExpAt := time.Now().UTC().Add(time.Duration(uc.tokenProvider.ExpireTokenInSeconds()) * time.Second)
	refreshTokenExpAt := time.Now().UTC().Add(time.Duration(uc.tokenProvider.ExpireRefreshTokenInSeconds()) * time.Second)

	// 5. Create and insert session to DB
	session := domain.NewSession(
		sessionId,
		userId,
		refreshToken,
		refreshTokenExpAt,
		accessTokenExpAt,
	)
	if err := uc.sessionCmdRepository.Create(ctx, session); err != nil {
		return nil, err
	}

	// 6. Return tokenDTO
	return &usecase.TokenResponseDTO{
		AccessToken:       accessToken,
		AccessTokenExpIn:  uc.tokenProvider.ExpireTokenInSeconds(),
		RefreshToken:      refreshToken,
		RefreshTokenExpIn: uc.tokenProvider.ExpireRefreshTokenInSeconds(),
	}, nil
}

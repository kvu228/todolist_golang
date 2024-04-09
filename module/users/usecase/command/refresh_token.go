package command

import (
	"context"
	"errors"
	"time"
	"to_do_list/common"
	"to_do_list/module/users/domain"
	"to_do_list/module/users/infrastructure/repositories/mysql"
	"to_do_list/module/users/usecase"
	"to_do_list/module/users/usecase/query"
)

type RefreshTokenUseCase interface {
	RefreshToken(ctx context.Context, refreshToken string) (*usecase.TokenResponseDTO, error)
}

type refreshTokenUseCase struct {
	userQueryRepository query.UserQueryRepository
	sessionRepo         mysql.SessionRepository
	tokenProvider       TokenProvider
	hasher              Hasher
}

func NewRefreshTokenUC(userQueryRepository query.UserQueryRepository, sessionRepo mysql.SessionRepository, tokenProvider TokenProvider, hasher Hasher) RefreshTokenUseCase {
	return &refreshTokenUseCase{userQueryRepository: userQueryRepository, sessionRepo: sessionRepo, tokenProvider: tokenProvider, hasher: hasher}
}

func (uc *refreshTokenUseCase) RefreshToken(ctx context.Context, refreshToken string) (*usecase.TokenResponseDTO, error) {
	// 1. Find Session with refresh token
	session, err := uc.sessionRepo.FindByRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}
	// 1.1 Check if refreshtoken is still valid
	if session.RefreshTokenExpiresAt().UnixNano() < time.Now().UnixNano() {
		return nil, errors.New("refresh token expired")
	}

	// 2. Find user
	user, err := uc.userQueryRepository.FindById(ctx, session.UserId())
	if err != nil {
		return nil, err
	}
	// 2.1 check if user banned
	if user.Status == "banned" {
		return nil, errors.New("user has been banned")
	}

	// 3. gen jwt
	userId := user.Id
	// 3.1 Generate new session ID
	sessionId, _ := common.GenUUID()
	accessToken, err := uc.tokenProvider.IssueToken(ctx, sessionId.String(), userId.String())
	if err != nil {
		return nil, err
	}

	//4. Insert new session to DB
	newRefreshToken, _ := uc.hasher.RandomString(common.LenRefreshToken)
	tokenExpAt := time.Now().UTC().Add(time.Second * time.Duration(uc.tokenProvider.ExpireTokenInSeconds()))
	refreshExpAt := time.Now().UTC().Add(time.Second * time.Duration(uc.tokenProvider.ExpireRefreshTokenInSeconds()))
	newSession := domain.NewSession(
		sessionId,
		userId,
		newRefreshToken,
		tokenExpAt,
		refreshExpAt,
	)
	if err := uc.sessionRepo.Create(ctx, newSession); err != nil {
		return nil, err
	}

	// 4.1 Delete previous session
	go func() {
		_ = uc.sessionRepo.Delete(ctx, session.Id())
	}()

	// 5. Return TokenResponseDTO
	return &usecase.TokenResponseDTO{
		AccessToken:       accessToken,
		AccessTokenExpIn:  uc.tokenProvider.ExpireTokenInSeconds(),
		RefreshToken:      newRefreshToken,
		RefreshTokenExpIn: uc.tokenProvider.ExpireRefreshTokenInSeconds(),
	}, nil
}

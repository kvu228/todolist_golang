package query

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"to_do_list/common"
)

type TokenParser interface {
	ParseToken(ctx context.Context, tokenString string) (claims *jwt.RegisteredClaims, err error)
}

type IntrospectUseCase interface {
	IntrospectToken(ctx context.Context, accessToken string) (common.Requester, error)
}

type introspectUseCase struct {
	userQueryRepo    UserQueryRepository
	sessionQueryRepo SessionQueryRepository
	tokenParser      TokenParser
}

func NewIntrospectUC(userQueryRepo UserQueryRepository, sessionQueryRepo SessionQueryRepository, tokenParser TokenParser) IntrospectUseCase {
	return &introspectUseCase{userQueryRepo: userQueryRepo, sessionQueryRepo: sessionQueryRepo, tokenParser: tokenParser}
}

func (uc *introspectUseCase) IntrospectToken(ctx context.Context, accessToken string) (common.Requester, error) {
	// Get claims from token
	claims, err := uc.tokenParser.ParseToken(ctx, accessToken)
	if err != nil {
		return nil, err
	}

	// get userId from claims
	userID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return nil, err
	}

	// get sessionId from claims
	sessionID, err := uuid.Parse(claims.ID)
	if err != nil {
		return nil, err
	}

	// Check if session exist
	if _, err := uc.sessionQueryRepo.Find(ctx, sessionID); err != nil {
		return nil, err
	}

	// Find user
	user, err := uc.userQueryRepo.FindById(ctx, userID)
	if err != nil {
		return nil, err
	}
	// check user status
	if user.Status() == "banned" {
		return nil, errors.New("user has been banned")
	}
	return common.NewRequesterData(
		userID,
		sessionID,
		user.FirstName(),
		user.LastName(),
		user.Email(),
		user.Role(),
		user.Status(),
	), nil
}

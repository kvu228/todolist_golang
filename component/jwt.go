package components

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	//sctx "github.com/viettranx/service-context"
	"time"
)

const (
	defaultSecret                      = "very-important-please-change-it!"
	defaultExpireTokenInSeconds        = 60 * 60 * 24 * 7  //one week
	defaultExpireRefreshTokenInSeconds = 60 * 60 * 24 * 14 // two weeks
)

var (
	ErrSecretKeyNotValid = errors.New("secret key is not valid")
	ErrTokenLifeTooShort = errors.New("token life is too short")
)

type jwtx struct {
	id                          string
	secret                      string
	expireTokenInSeconds        int
	expireRefreshTokenInSeconds int
}

func NewJWTProvider(secret string, expireTokenInSeconds, expireRefreshTokenInSeconds int) *jwtx {
	return &jwtx{
		secret:                      secret,
		expireTokenInSeconds:        expireTokenInSeconds,
		expireRefreshTokenInSeconds: expireRefreshTokenInSeconds,
	}
}

func NewJWT(id string) *jwtx {
	return &jwtx{
		id: id,
	}
}

func (j *jwtx) ID() string {
	return j.id
}

func (j *jwtx) ExpireTokenInSeconds() int {
	return j.expireTokenInSeconds
}

func (j *jwtx) ExpireRefreshTokenInSeconds() int {
	return j.expireRefreshTokenInSeconds
}

/*
func (j *jwtx) InitFlags() {
	flag.StringVar(
		&j.secret,
		"jwt-secret",
		defaultSecret,
		"Secret key to sign JWT",
	)

	flag.IntVar(
		&j.expireTokenInSeconds,
		"jwt-exp-secs",
		defaultExpireTokenInSeconds,
		"Number of seconds access token will expired",
	)

	flag.IntVar(
		&j.expireRefreshTokenInSeconds,
		"refresh-exp-secs",
		defaultExpireRefreshTokenInSeconds,
		"Number of seconds refresh token will expired",
	)
}

func (j *jwtx) Activate(_ sctx.ServiceContext) error {
	if len(j.secret) < 32 {
		return ErrSecretKeyNotValid
	}
	if j.expireRefreshTokenInSeconds <= 60 {
		return ErrTokenLifeTooShort
	}
	return nil
}

func (j *jwtx) Stop() error {
	return nil
}

*/

func (j *jwtx) IssueToken(ctx context.Context, id, subject string) (token string, err error) {
	now := time.Now().UTC()

	// create a claim
	claims := jwt.RegisteredClaims{
		Subject:   subject,
		ExpiresAt: jwt.NewNumericDate(now.Add(time.Second * time.Duration(j.expireRefreshTokenInSeconds))),
		NotBefore: jwt.NewNumericDate(now),
		IssuedAt:  jwt.NewNumericDate(now),
		ID:        id,
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// issue token
	tokenSignedStr, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return "", err
	}

	return tokenSignedStr, nil
}

func (j *jwtx) ParseToken(ctx context.Context, tokenString string) (claims *jwt.RegisteredClaims, err error) {
	var registerClaims jwt.RegisteredClaims
	token, err := jwt.ParseWithClaims(tokenString, &registerClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}

	return &registerClaims, nil
}

// TokenProvider : Define token provider interface
type TokenProvider interface {
	IssueToken(ctx context.Context, id, sub string) (token string, err error)
	ExpireTokenInSeconds() int
	ExpireRefreshTokenInSeconds() int
	ParseToken(ctx context.Context, tokenString string) (claims *jwt.RegisteredClaims, err error)
}

package domain

import (
	"github.com/google/uuid"
	"time"
)

type Session struct {
	id                    uuid.UUID
	userId                uuid.UUID
	refreshToken          string
	refreshTokenExpiresAt time.Time
	accessTokenExpiredAt  time.Time
}

func NewSession(id uuid.UUID, userId uuid.UUID, refreshToken string, refreshTokenExpiresAt time.Time, accessTokenExpiredAt time.Time) *Session {
	return &Session{id: id, userId: userId, refreshToken: refreshToken, refreshTokenExpiresAt: refreshTokenExpiresAt, accessTokenExpiredAt: accessTokenExpiredAt}
}

func (s *Session) RefreshTokenExpiresAt() time.Time {
	return s.refreshTokenExpiresAt
}

func (s *Session) Id() uuid.UUID {
	return s.id
}

func (s *Session) UserId() uuid.UUID {
	return s.userId
}

func (s *Session) RefreshToken() string {
	return s.refreshToken
}

func (s *Session) AccessTokenExpiredAt() time.Time {
	return s.accessTokenExpiredAt
}

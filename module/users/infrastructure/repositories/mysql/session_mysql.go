package mysql

import (
	"context"
	"gorm.io/gorm"
	"to_do_list/common"
	"to_do_list/module/users/domain"
)

type sessionMySQLRepo struct {
	db *gorm.DB
}

func NewSessionMySQLRepo(db *gorm.DB) *sessionMySQLRepo {
	return &sessionMySQLRepo{db: db}
}

func (s *sessionMySQLRepo) Create(ctx context.Context, session *domain.Session) error {
	dto := SessionDTO{
		Id:                session.Id(),
		UserId:            session.UserId(),
		RefreshToken:      session.RefreshToken(),
		RefreshTokenExpAt: session.RefreshTokenExpiresAt(),
		AccessTokenExpAt:  session.AccessTokenExpiredAt(),
	}

	if err := s.db.Table(common.TbNameSessions).Create(&dto).Error; err != nil {
		return err
	}
	return nil
}

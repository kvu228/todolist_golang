package mysql

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"to_do_list/common"
	"to_do_list/module/users/domain"
)

type sessionMySQLRepo struct {
	db *gorm.DB
}

func NewSessionMySQLRepo(db *gorm.DB) SessionRepository {
	return &sessionMySQLRepo{db: db}
}

func (s *sessionMySQLRepo) Find(ctx context.Context, id uuid.UUID) (session *domain.Session, err error) {
	var dto SessionDTO
	err = s.db.Table(common.TbNameSessions).Where("id = ?", id).First(&dto).Error
	if err != nil {
		return nil, err
	}
	return dto.ToEntity()
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

type SessionRepository interface {
	SessionQueryRepository
	SessionCmdRepository
}

type SessionQueryRepository interface {
	Find(ctx context.Context, id uuid.UUID) (session *domain.Session, err error)
}

type SessionCmdRepository interface {
	Create(ctx context.Context, session *domain.Session) error
}

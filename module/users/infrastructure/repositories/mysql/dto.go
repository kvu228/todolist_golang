package mysql

import (
	"github.com/google/uuid"
	"time"
	"to_do_list/module/users/domain"
)

type UserDTO struct {
	Id        uuid.UUID `json:"id" gorm:"column:id"`
	FirstName string    `json:"first_name" gorm:"column:first_name"`
	LastName  string    `json:"last_name" gorm:"column:last_name"`
	Email     string    `json:"email" gorm:"column:email"`
	Password  string    `json:"password" gorm:"column:password"`
	Salt      string    `json:"salt" gorm:"column:salt"`
	Avatar    string    `json:"avatar" gorm:"column:avatar"`
	Status    string    `json:"status" gorm:"column:status"`
	Role      string    `json:"role" gorm:"column:role"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (dto *UserDTO) ToEntity() (*domain.User, error) {
	return domain.NewUser(
		dto.Id,
		dto.FirstName,
		dto.LastName,
		dto.Email,
		dto.Password,
		dto.Salt,
		dto.Avatar,
		dto.Status,
		dto.Role,
		dto.CreatedAt,
		dto.UpdatedAt,
	), nil
}

type SessionDTO struct {
	Id                uuid.UUID `gorm:"column:id"`
	UserId            uuid.UUID `gorm:"column:user_id"`
	RefreshToken      string    `gorm:"column:refresh_token"`
	RefreshTokenExpAt time.Time `gorm:"column:refresh_token_exp_at"`
	AccessTokenExpAt  time.Time `gorm:"column:access_token_exp_at"`
}

func (s *SessionDTO) ToEntity() (*domain.Session, error) {
	return domain.NewSession(
		s.Id,
		s.UserId,
		s.RefreshToken,
		s.RefreshTokenExpAt,
		s.AccessTokenExpAt,
	), nil
}

type OwnerDTO struct {
	Id        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}

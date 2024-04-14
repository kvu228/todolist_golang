package mysql

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"to_do_list/common"
	"to_do_list/module/users/domain"
)

type userMySQLRepo struct {
	db *gorm.DB
}

func NewUserMySQLRepo(db *gorm.DB) UserRepository {
	return &userMySQLRepo{db: db}
}

func (u *userMySQLRepo) Create(ctx context.Context, user *domain.User) error {
	dto := UserDTO{
		Id:        user.Id(),
		FirstName: user.FirstName(),
		LastName:  user.LastName(),
		Email:     user.Email(),
		Password:  user.Password(),
		Salt:      user.Salt(),
		Avatar:    user.Avatar(),
		Status:    user.Status(),
		Role:      user.Role(),
	}
	if err := u.db.Table(common.TbNameUsers).Create(&dto).Error; err != nil {
		return err
	}
	return nil
}

func (u *userMySQLRepo) Update(ctx context.Context, user *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userMySQLRepo) Delete(ctx context.Context, id uuid.UUID) error {
	if err := u.db.Table(common.TbNameUsers).Where("id = ?", id).Delete(&UserDTO{}).Error; err != nil {
		return err
	}
	return nil
}

func (u *userMySQLRepo) FindById(ctx context.Context, id uuid.UUID) (user *domain.User, err error) {
	userDTO := &UserDTO{}
	if err := u.db.Table(common.TbNameUsers).Where("id = ?", id).First(&userDTO).Error; err != nil {
		return nil, err
	}
	return userDTO.ToEntity()
}

func (u *userMySQLRepo) FindByIds(ctx context.Context, ids []uuid.UUID) (users []*domain.User, err error) {
	usersDTO := make([]*UserDTO, len(ids))
	if err := u.db.Table(common.TbNameUsers).Where("id IN (?)", ids).Find(&usersDTO).Error; err != nil {
		return nil, err
	}
	users = make([]*domain.User, len(usersDTO))
	for index, userDTO := range usersDTO {
		users[index], _ = userDTO.ToEntity()
	}
	return users, nil
}

func (u *userMySQLRepo) FindByEmail(ctx context.Context, email string) (user *domain.User, err error) {
	userDTO := &UserDTO{}
	if err := u.db.Table(common.TbNameUsers).Where("email = ?", email).First(userDTO).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound
		}
		return nil, err
	}
	return userDTO.ToEntity()
}

//func (u *userMySQLRepo) FindWithIds(ctx context.Context, ids []uuid.UUID) (owners []OwnerDTO, err error) {
//	if err := u.db.Table(common.TbNameUsers).Where("id IN (?)", ids).Find(&owners).Error; err != nil {
//		return nil, err
//	}
//	return owners, nil
//}

type UserRepository interface {
	UserCmdRepository
	UserQueryRepository
}

type UserCmdRepository interface {
	Create(ctx context.Context, user *domain.User) error
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type UserQueryRepository interface {
	FindById(ctx context.Context, id uuid.UUID) (user *domain.User, err error)
	FindByIds(ctx context.Context, ids []uuid.UUID) (uses []*domain.User, err error)
	FindByEmail(ctx context.Context, email string) (user *domain.User, err error)
	//FindWithIds(ctx context.Context, ids []uuid.UUID) (owners []OwnerDTO, err error)
}

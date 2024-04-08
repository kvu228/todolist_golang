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

func NewUserMySQLRepo(db *gorm.DB) *userMySQLRepo {
	return &userMySQLRepo{db: db}
}

func (u *userMySQLRepo) Create(ctx context.Context, user *domain.User) error {
	dto := domain.UserDTO{
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
	if err := u.db.Table(common.TbNameUsers).Where("id = ?", id).Delete(&domain.UserDTO{}).Error; err != nil {
		return err
	}
	return nil
}

func (u *userMySQLRepo) FindById(ctx context.Context, id uuid.UUID) (user *domain.UserDTO, err error) {
	user = &domain.UserDTO{}
	if err := u.db.Table(common.TbNameUsers).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userMySQLRepo) FindByIds(ctx context.Context, ids []uuid.UUID) (uses []*domain.UserDTO, err error) {
	if err := u.db.Table(common.TbNameUsers).Where("id IN (?)", ids).Find(&uses).Error; err != nil {
		return nil, err
	}
	return uses, nil
}

func (u *userMySQLRepo) FindByEmail(ctx context.Context, email string) (user *domain.UserDTO, err error) {
	if err := u.db.Table(common.TbNameUsers).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound
		}
		return nil, err
	}
	return user, nil
}

func (u *userMySQLRepo) FindWithIds(ctx context.Context, ids []uuid.UUID) (owners []domain.OwnerDTO, err error) {

	if err := u.db.Table(common.TbNameUsers).Where("id IN (?)", ids).Find(&owners).Error; err != nil {
		return nil, err
	}
	return owners, nil
}

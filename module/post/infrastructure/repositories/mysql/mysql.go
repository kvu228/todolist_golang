package mysql

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"to_do_list/common"
	"to_do_list/module/post/domain"
)

type postMySQLRepo struct {
	db *gorm.DB
}

func NewPostsMySQLRepo(db *gorm.DB) *postMySQLRepo {
	return &postMySQLRepo{db: db}
}

func (repo *postMySQLRepo) FindWithIds(ctx context.Context, ids []uuid.UUID) (posts []domain.PostDTO, err error) {
	if err := repo.db.Table(common.TbNamePosts).
		Where("id IN (?)", ids).
		Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (repo *postMySQLRepo) FindWithParams(ctx context.Context, params *domain.ListPostsParams) (posts []domain.PostDTO, err error) {
	db := repo.db.Table(common.TbNamePosts)
	if params.OwnerId != "" {
		db = db.Where("owner_id = ?", params.OwnerId)
	}

	// Get number of total post
	var count int64
	if err := db.Count(&count).Error; err != nil {
		return nil, err
	}
	params.Total = count

	// Validate Page, Limit params
	params.Process()

	offset := params.Limit * (params.Page - 1)
	if err := db.Offset(offset).
		Limit(params.Limit).
		Order("id desc").
		Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

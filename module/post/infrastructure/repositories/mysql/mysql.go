package mysql

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"to_do_list/common"
	"to_do_list/module/post/domain"
	"to_do_list/module/post/usecase"
)

type PostRepository interface {
	PostCmdRepository
	PostQueryRepository
}

type postMySQLRepo struct {
	db *gorm.DB
}

func NewPostsMySQLRepo(db *gorm.DB) PostRepository {
	return &postMySQLRepo{db: db}
}

func (repo *postMySQLRepo) CreatePost(ctx context.Context, post *domain.Post) error {
	postDTO := PostDTO{
		BaseModel: common.BaseModel{
			Id: post.Id(),
		},
		Title:   post.Title(),
		Body:    post.Body(),
		Status:  post.Status(),
		OwnerId: post.OwnerId(),
	}
	return repo.db.Table(common.TbNamePosts).Create(&postDTO).Error
}

func (repo *postMySQLRepo) FindWithIds(ctx context.Context, ids []uuid.UUID) (posts []usecase.PostDTO, err error) {
	if err := repo.db.Table(common.TbNamePosts).
		Where("id IN (?)", ids).
		Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (repo *postMySQLRepo) FindWithParams(ctx context.Context, params *usecase.ListPostsParams) (posts []usecase.PostDTO, err error) {
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

type PostCmdRepository interface {
	CreatePost(ctx context.Context, post *domain.Post) error
}

type PostQueryRepository interface {
	FindWithIds(ctx context.Context, ids []uuid.UUID) (posts []usecase.PostDTO, err error)
	FindWithParams(ctx context.Context, params *usecase.ListPostsParams) (posts []usecase.PostDTO, err error)
}

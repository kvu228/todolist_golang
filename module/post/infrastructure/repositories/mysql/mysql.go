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
		Title:   post.Title(),
		Body:    post.Body(),
		Status:  "doing",
		OwnerId: post.OwnerId(),
	}
	postDTO.BaseModel.Id = post.Id()
	postDTO.CreatedAt = post.CreatedAt()
	postDTO.UpdatedAt = post.UpdatedAt()

	return repo.db.Table(common.TbNamePosts).Create(&postDTO).Error
}

func (repo *postMySQLRepo) FindWithIds(ctx context.Context, ids []uuid.UUID) (posts []*domain.Post, err error) {
	postDTOs := make([]*PostDTO, len(ids))
	if err := repo.db.Table(common.TbNamePosts).
		Where("id IN (?)", ids).
		Find(&postDTOs).Error; err != nil {
		return nil, err
	}
	posts = make([]*domain.Post, len(postDTOs))
	for i, postDTO := range postDTOs {
		posts[i], _ = postDTO.ToEntity()
	}
	return posts, nil
}

func (repo *postMySQLRepo) FindWithParams(ctx context.Context, params *usecase.ListPostsParams) (posts []*domain.Post, err error) {
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
	postDTOs := make([]*PostDTO, offset)
	if err := db.Offset(offset).
		Limit(params.Limit).
		Order("id desc").
		Find(&postDTOs).Error; err != nil {
		return nil, err
	}
	posts = make([]*domain.Post, len(postDTOs))
	for i, postDTO := range postDTOs {
		posts[i], _ = postDTO.ToEntity()
	}
	return posts, nil
}

type PostCmdRepository interface {
	CreatePost(ctx context.Context, post *domain.Post) error
}

type PostQueryRepository interface {
	FindWithIds(ctx context.Context, ids []uuid.UUID) (posts []*domain.Post, err error)
	FindWithParams(ctx context.Context, params *usecase.ListPostsParams) (posts []*domain.Post, err error)
}

package query

import (
	"context"
	"github.com/google/uuid"
	"to_do_list/module/post/domain"
)

// useCases
type listPostsQueryUseCase struct {
	postRepository PostRepository
	userRepository UserRepository
}

func NewListPostsQueryUseCase(postsRepository PostRepository, userRepository UserRepository) *listPostsQueryUseCase {
	return &listPostsQueryUseCase{
		postRepository: postsRepository,
		userRepository: userRepository}
}

func (uc *listPostsQueryUseCase) ListPosts(ctx context.Context, param *domain.ListPostsParams) ([]domain.PostDTO, error) {
	// Fetch all post with params
	posts, err := uc.postRepository.FindWithParams(ctx, param)
	if err != nil {
		return nil, err
	}
	// Get the owner_id from those post
	ownerIds := make([]uuid.UUID, len(posts))
	for i := range posts {
		ownerIds[i] = posts[i].OwnerId
	}

	// From the ownerIds above fetch data of those users
	owners, err := uc.userRepository.FindWithIds(ctx, ownerIds)
	if err != nil {
		return nil, err
	}

	// Create a map with (key:value) = (uuid:ownerDTO)
	ownersMap := make(map[uuid.UUID]*domain.OwnerDTO, len(owners))
	for i, owner := range owners {
		ownersMap[owner.Id] = &owners[i]
	}

	for i, post := range posts {
		posts[i].Owner = *ownersMap[post.OwnerId]
	}

	return posts, nil
}

// repositories Interfaces
type PostRepository interface {
	FindWithIds(ctx context.Context, ids []uuid.UUID) (posts []domain.PostDTO, err error)
	FindWithParams(ctx context.Context, params *domain.ListPostsParams) (posts []domain.PostDTO, err error)
}

type UserRepository interface {
	FindWithIds(ctx context.Context, ids []uuid.UUID) (owners []domain.OwnerDTO, err error)
}

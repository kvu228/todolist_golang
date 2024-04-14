package query

import (
	"context"
	"github.com/google/uuid"
	"to_do_list/module/post/usecase"
)

type ListPostsUseCase interface {
	ListPosts(ctx context.Context, param *usecase.ListPostsParams) ([]usecase.PostDTO, error)
}

// useCases
type listPostsQueryUseCase struct {
	postQueryRepository PostQueryRepository
	userQueryRepository UserQueryRepository
}

func NewListPostsQueryUseCase(postQueryRepository PostQueryRepository, userQueryRepository UserQueryRepository) ListPostsUseCase {
	return &listPostsQueryUseCase{
		postQueryRepository: postQueryRepository,
		userQueryRepository: userQueryRepository}
}

func (uc *listPostsQueryUseCase) ListPosts(ctx context.Context, param *usecase.ListPostsParams) ([]usecase.PostDTO, error) {
	// Fetch all post with params
	posts, err := uc.postQueryRepository.FindWithParams(ctx, param)
	if err != nil {
		return nil, err
	}
	// Get the owner_id from those post
	ownerIds := make([]uuid.UUID, len(posts))
	for i := range posts {
		ownerIds[i] = posts[i].OwnerId
	}

	// From the ownerIds above fetch data of those users
	owners, err := uc.userQueryRepository.FindWithIds(ctx, ownerIds)
	if err != nil {
		return nil, err
	}

	// Create a map with (key:value) = (uuid:ownerDTO)
	ownersMap := make(map[uuid.UUID]*usecase.OwnerDTO, len(owners))
	for i, owner := range owners {
		ownersMap[owner.Id] = &owners[i]
	}

	for i, post := range posts {
		posts[i].Owner = *ownersMap[post.OwnerId]
	}

	return posts, nil
}

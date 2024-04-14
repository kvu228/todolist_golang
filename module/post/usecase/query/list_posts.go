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
	postEntities, err := uc.postQueryRepository.FindWithParams(ctx, param)
	if err != nil {
		return nil, err
	}
	// Get the owner_id from those post
	ownerIds := make([]uuid.UUID, len(postEntities))
	for i := range postEntities {
		ownerIds[i] = postEntities[i].OwnerId()
	}

	// From the ownerIds above fetch data of those users
	ownerDTOs, err := uc.userQueryRepository.FindWithIds(ctx, ownerIds)
	if err != nil {
		return nil, err
	}

	// Create a map with (key:value) = (uuid:ownerDTO)
	ownersMap := make(map[uuid.UUID]*usecase.OwnerDTO, len(ownerDTOs))
	for i, owner := range ownerDTOs {
		ownersMap[owner.Id] = ownerDTOs[i]
	}

	postDTOs := make([]usecase.PostDTO, len(postEntities))
	for i, post := range postEntities {
		postDTOs[i].Id = post.Id()
		postDTOs[i].Title = post.Title()
		postDTOs[i].Body = post.Body()
		postDTOs[i].CreatedAt = post.CreatedAt()
		postDTOs[i].UpdatedAt = post.UpdatedAt()
		postDTOs[i].OwnerId = post.OwnerId()
		postDTOs[i].Owner = *ownersMap[post.OwnerId()]
	}

	return postDTOs, nil
}

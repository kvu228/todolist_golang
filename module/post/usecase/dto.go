package usecase

import (
	"github.com/google/uuid"
	"time"
	"to_do_list/common"
	"to_do_list/module/post/domain"
)

type OwnerDTO struct {
	Id        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}

type PostDTO struct {
	common.BaseModel
	Title   string    `json:"title"`
	Body    string    `json:"body"`
	Status  string    `json:"status"`
	OwnerId uuid.UUID `json:"owner_id"`
	Owner   OwnerDTO  `json:"owner"`
}

type NewPostDTO struct {
	Id      uuid.UUID
	Title   string
	Body    string
	Status  string
	OwnerId uuid.UUID
}

func (n *NewPostDTO) ToEntity() *domain.Post {
	return domain.NewPost(
		n.Id,
		n.Title,
		n.Body,
		n.Status,
		n.OwnerId,
		time.Now().UTC(),
		time.Now().UTC(),
	)
}

type ListPostsParams struct {
	common.Paging
	ListPostsFilter
}

type ListPostsFilter struct {
	OwnerId string `json:"owner_id,omitempty" form:"owner_id"`
}

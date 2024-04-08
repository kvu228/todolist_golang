package domain

import (
	"github.com/google/uuid"
	"to_do_list/common"
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
	Owner   OwnerDTO
}

type ListPostsParams struct {
	common.Paging
	ListPostsFilter
}

type ListPostsFilter struct {
	OwnerId string `json:"owner_id"`
}

func (dto *PostDTO) ToEntity() (*Post, error) {
	return NewPost(
		dto.Id,
		dto.Title,
		dto.Body,
		dto.CreatedAt,
		dto.UpdatedAt,
		dto.Status,
		dto.OwnerId,
	), nil
}

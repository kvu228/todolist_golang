package mysql

import (
	"github.com/google/uuid"
	"to_do_list/common"
	"to_do_list/module/post/domain"
)

type PostDTO struct {
	common.BaseModel
	Title   string    `json:"title"`
	Body    string    `json:"body"`
	Status  string    `json:"status"`
	OwnerId uuid.UUID `json:"owner_id"`
}

func (p *PostDTO) ToEntity() (*domain.Post, error) {
	return domain.NewPost(
		p.Id,
		p.Title,
		p.Body,
		p.Status,
		p.OwnerId,
		p.CreatedAt,
		p.UpdatedAt,
	), nil
}

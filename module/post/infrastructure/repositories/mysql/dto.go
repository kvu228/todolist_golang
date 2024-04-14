package mysql

import (
	"github.com/google/uuid"
	"to_do_list/common"
)

type PostDTO struct {
	common.BaseModel
	Title   string    `json:"title"`
	Body    string    `json:"body"`
	Status  string    `json:"status"`
	OwnerId uuid.UUID `json:"owner_id"`
}

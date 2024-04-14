package domain

import (
	"github.com/google/uuid"
)

type Post struct {
	id      uuid.UUID
	title   string
	body    string
	status  string
	ownerId uuid.UUID
}

func NewPost(id uuid.UUID, title string, body string, status string, ownerId uuid.UUID) *Post {
	return &Post{id: id, title: title, body: body, status: status, ownerId: ownerId}
}

func (p Post) Id() uuid.UUID {
	return p.id
}

func (p Post) Title() string {
	return p.title
}

func (p Post) Body() string {
	return p.body
}

func (p Post) Status() string {
	return p.status
}

func (p Post) OwnerId() uuid.UUID {
	return p.ownerId
}

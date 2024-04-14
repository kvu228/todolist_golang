package domain

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	id        uuid.UUID
	title     string
	body      string
	status    string
	ownerId   uuid.UUID
	createdAt time.Time
	updatedAt time.Time
}

func NewPost(id uuid.UUID, title string, body string, status string, ownerId uuid.UUID, createdAt time.Time, updatedAt time.Time) *Post {
	return &Post{id: id, title: title, body: body, status: status, ownerId: ownerId, createdAt: createdAt, updatedAt: updatedAt}
}

func (p *Post) Id() uuid.UUID {
	return p.id
}

func (p *Post) Title() string {
	return p.title
}

func (p *Post) Body() string {
	return p.body
}

func (p *Post) Status() string {
	return p.status
}

func (p *Post) OwnerId() uuid.UUID {
	return p.ownerId
}

func (p *Post) CreatedAt() time.Time {
	return p.createdAt
}
func (p *Post) UpdatedAt() time.Time {
	return p.updatedAt
}

package domain

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	id        uuid.UUID
	title     string
	body      string
	createdAt time.Time
	updatedAt time.Time
	status    string
	ownerId   uuid.UUID
}

func NewPost(id uuid.UUID, title string, body string, createdAt time.Time, updatedAt time.Time, status string, ownerId uuid.UUID) *Post {
	return &Post{id: id, title: title, body: body, createdAt: createdAt, updatedAt: updatedAt, status: status, ownerId: ownerId}
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

func (p Post) CreatedAt() time.Time {
	return p.createdAt
}

func (p Post) UpdatedAt() time.Time {
	return p.updatedAt
}

func (p Post) Status() string {
	return p.status
}

func (p Post) OwnerId() uuid.UUID {
	return p.ownerId
}

package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	id        uuid.UUID
	firstName string
	lastName  string
	email     string
	password  string
	salt      string
	avatar    string
	status    string
	role      string
	createdAt time.Time
	updatedAt time.Time
}

func NewUser(id uuid.UUID, firstName string, lastName string, email string, password string, salt string, avatar string, status string, role string, createdAt time.Time, updatedAt time.Time) *User {
	return &User{id: id, firstName: firstName, lastName: lastName, email: email, password: password, salt: salt, avatar: avatar, status: status, role: role, createdAt: createdAt, updatedAt: updatedAt}
}

func (u *User) Id() uuid.UUID {
	return u.id
}

func (u *User) FirstName() string {
	return u.firstName
}

func (u *User) LastName() string {
	return u.lastName
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Password() string {
	return u.password
}

func (u *User) Salt() string {
	return u.salt
}

func (u *User) Avatar() string {
	return u.avatar
}

func (u *User) Status() string {
	return u.status
}

func (u *User) Role() string {
	return u.role
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u *User) ChangeAvatar(ava string) error {
	u.avatar = ava
	return nil
}

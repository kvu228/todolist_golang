package common

import "github.com/google/uuid"

type Requester interface {
	Id() uuid.UUID
	SessionId() uuid.UUID
	FirstName() string
	LastName() string
	Email() string
	Role() string
	Status() string
}

type requesterData struct {
	id        uuid.UUID
	sessionId uuid.UUID
	firstName string
	lastName  string
	email     string
	role      string
	status    string
}

func NewRequesterData(id uuid.UUID, sessionId uuid.UUID, firstName string, lastName string, email string, role string, status string) Requester {
	return &requesterData{id: id, sessionId: sessionId, firstName: firstName, lastName: lastName, email: email, role: role, status: status}
}

func (r *requesterData) Id() uuid.UUID {
	return r.id
}

func (r *requesterData) SessionId() uuid.UUID {
	return r.sessionId
}

func (r *requesterData) FirstName() string {
	return r.firstName
}

func (r *requesterData) LastName() string {
	return r.lastName
}

func (r *requesterData) Email() string {
	return r.email
}

func (r *requesterData) Role() string {
	return r.role
}

func (r *requesterData) Status() string {
	return r.status
}

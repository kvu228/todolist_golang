package common

import (
	"github.com/google/uuid"
	"time"
)

type BaseModel struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GenNewBaseModel() *BaseModel {
	now := time.Now()
	return &BaseModel{
		Id:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func GenUUID() (uuid.UUID, error) {
	return uuid.NewV7()
}
func ParseUUIDFromString(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}

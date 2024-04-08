package domain

import "github.com/google/uuid"

type UserDTO struct {
	Id        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Salt      string    `json:"salt"`
	Avatar    string    `json:"avatar"`
	Status    string    `json:"status"`
	Role      string    `json:"role"`
}

func (dto *UserDTO) ToEntity() (*User, error) {
	return NewUser(
		dto.Id,
		dto.FirstName,
		dto.LastName,
		dto.Email,
		dto.Password,
		dto.Salt,
		dto.Avatar,
		dto.Status,
		dto.Role,
	), nil
}

type EmailPasswordRegistrationDTO struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type EmailPasswordLoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponseDTO struct {
	AccessToken       string `json:"access_token"`
	AccessTokenExpIn  int    `json:"access_token_exp_in"`
	RefreshToken      string `json:"refresh_token"`
	RefreshTokenExpIn int    `json:"refresh_token_exp_in"`
}

type OwnerDTO struct {
	Id        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}

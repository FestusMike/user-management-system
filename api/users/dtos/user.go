package dtos

import (
	"errors"
	"math/rand/v2"
	"strings"
	"unicode/utf8"
	user_models "github.com/FestusMike/user-management-system/api/users/models"
	helpers "github.com/FestusMike/user-management-system/includes"
)

type CreateUserRequestDTO struct{
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func (cr CreateUserRequestDTO) Create() (*user_models.User, error){
	if err := helpers.Validate.Struct(cr); err != nil{
		return nil, err
	}

	Name := strings.TrimSpace(cr.Name)
	Email := strings.TrimSpace(cr.Email)

	if Name == "" || utf8.RuneCountInString(Name) < 3 {
		return nil, errors.New("Name must be at least three charaacters")
	}

	if Email == ""{
		return nil, errors.New("Email addres is invalid")
	}

	return &user_models.User{
		ID: rand.Uint64(),
		Email: Email,
		Name: Name,
	}, nil
}

type CreateUserResponseDTO struct{
	ID uint64 `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

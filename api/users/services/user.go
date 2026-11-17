package services

import (
	user_models "github.com/FestusMike/user-management-system/api/users/models"
	user_dtos "github.com/FestusMike/user-management-system/api/users/dtos"
)

func ToCreateUserResponse(user user_models.User) user_dtos.CreateUserResponseDTO{
	return user_dtos.CreateUserResponseDTO{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
	}
}
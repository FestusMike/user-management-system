package dtos

import (
	"errors"
	"math/rand"
	"strings"
	student_models "github.com/FestusMike/user-management-system/api/learnhub/models"
	user_models "github.com/FestusMike/user-management-system/api/users/models"
	helpers "github.com/FestusMike/user-management-system/includes"
)

type CreateStudentRequestDTO struct{
	Department string `json:"department" validate:"required"`
}

func (cs CreateStudentRequestDTO) Create(user user_models.User) (*student_models.Student, error){
	if err := helpers.Validate.Struct(cs); err != nil {
		return nil, err
	}

	Department := strings.TrimSpace(cs.Department)

	if Department == "" {
		return nil, errors.New("Department Cannot be blank")
	}

	return &student_models.Student{
		ID: rand.Uint64(),
		User: user,
		Department: Department,
	}, nil
}

type CreateStudentResponseDTO struct{
	ID uint64 `json:"id"`
	User user_models.User `json:"user"`
	Department string `json:"department"`
}
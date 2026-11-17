package services


import (
	learnhub_dtos "github.com/FestusMike/user-management-system/api/learnhub/dtos"
	student_models "github.com/FestusMike/user-management-system/api/learnhub/models"
)


func ToStudentCreateResponse(student student_models.Student) learnhub_dtos.CreateStudentResponseDTO{
	return learnhub_dtos.CreateStudentResponseDTO{
		ID: student.ID,
		User: student.User,
		Department: student.Department,
	}
}
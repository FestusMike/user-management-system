package main

import (
	"encoding/json"
	"fmt"
	"log"
	//"os"

	learnhub_dtos "github.com/FestusMike/user-management-system/api/learnhub/dtos"
	learnhub_models "github.com/FestusMike/user-management-system/api/learnhub/models"
	learnhub_services "github.com/FestusMike/user-management-system/api/learnhub/services"
	user_dtos "github.com/FestusMike/user-management-system/api/users/dtos"
	user_services "github.com/FestusMike/user-management-system/api/users/services"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	userBody := []byte(`{
		"name": "Mike Festus",
		"email": "mike@example.com"
	}`)
	var userRequest user_dtos.CreateUserRequestDTO
	if err := json.Unmarshal(userBody, &userRequest); err != nil {
		return fmt.Errorf("decode user request: %w", err)
	}
	user, err := userRequest.Create()
	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}

	studentBody := []byte(`{
		"department": "Computer Science"
	}`)
	var studentRequest learnhub_dtos.CreateStudentRequestDTO
	if err := json.Unmarshal(studentBody, &studentRequest); err != nil {
		return fmt.Errorf("decode student request: %w", err)
	}
	student, err := studentRequest.Create(*user)
	if err != nil {
		return fmt.Errorf("create student: %w", err)
	}

	courseBody := []byte(`{
		"title": "Introduction to Go",
		"description": "Learn Go structs, validation, and relationships.",
		"unit": 3
	}`)
	var courseRequest learnhub_dtos.CreateCourseRequestDTO
	if err := json.Unmarshal(courseBody, &courseRequest); err != nil {
		return fmt.Errorf("decode course request: %w", err)
	}
	course, err := courseRequest.Create()
	if err != nil {
		return fmt.Errorf("create course: %w", err)
	}

	// Simulate a client sending the IDs returned when these records were created.
	enrollmentBody := []byte(fmt.Sprintf(`{
		"student_id": %d,
		"course_id": %d
	}`, student.ID, course.ID))
	var enrollmentRequest learnhub_dtos.EnrollmentRequestDTO
	if err := json.Unmarshal(enrollmentBody, &enrollmentRequest); err != nil {
		return fmt.Errorf("decode enrollment request: %w", err)
	}
	enrollment, err := enrollmentRequest.Create(*student, *course)
	if err != nil {
		return fmt.Errorf("create enrollment: %w", err)
	}

	response := struct {
		User       user_dtos.CreateUserResponseDTO        `json:"user"`
		Student    learnhub_dtos.CreateStudentResponseDTO `json:"student"`
		Course     *learnhub_models.Course                `json:"course"`
		Enrollment *learnhub_models.Enrollment            `json:"enrollment"`
	}{
		User:       user_services.ToCreateUserResponse(*user),
		Student:    learnhub_services.ToStudentCreateResponse(*student),
		Course:     course,
		Enrollment: enrollment,
	}

	respJson, marshalerr := json.MarshalIndent(response, "", "  ")

	if marshalerr != nil{
		return fmt.Errorf("Error converting struct to json: %w", marshalerr)
	}

	fmt.Println(string(respJson))

	return nil
}

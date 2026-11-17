package dtos

import (
	"errors"
	learnhub_models "github.com/FestusMike/user-management-system/api/learnhub/models"
)

type EnrollmentRequestDTO struct{
	StudentID uint64 `json:"student_id"`
	CourseID uint64 `json:"course_id"`
}

func (eq EnrollmentRequestDTO) Create(
    student learnhub_models.Student,
    course learnhub_models.Course,
) (*learnhub_models.Enrollment, error) {
    if eq.StudentID == 0 || eq.CourseID == 0 {
        return nil, errors.New("Student ID and course ID are required")
    }

    if student.ID != eq.StudentID {
        return nil, errors.New("Student does not match the request")
    }

    if course.ID != eq.CourseID {
        return nil, errors.New("Course does not match the request")
    }

    return &learnhub_models.Enrollment{
        Student: student,
        Course:  course,
    }, nil
}
package models

import(
	user_models "github.com/FestusMike/user-management-system/api/users/models" 
)

type Student struct{
	ID uint64
	user_models.User
	Department string
}

type Course struct{
	ID uint64
	Title string
	Description string
	Unit uint64
}

type Enrollment struct{
	Student Student
	Course Course
}
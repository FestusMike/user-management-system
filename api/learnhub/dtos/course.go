package dtos

import (
	"errors"
	"math/rand"
	"strings"
	learnhub_models "github.com/FestusMike/user-management-system/api/learnhub/models"
	helpers "github.com/FestusMike/user-management-system/includes"
)

type CreateCourseRequestDTO struct{
	Title string `json:"title"`
	Description string `json:"description"`
	Unit uint64 `json:"unit"`
}


func (cc CreateCourseRequestDTO) Create() (*learnhub_models.Course, error){
	if err := helpers.Validate.Struct(cc); err != nil{
		return nil, err
	}

	Title := strings.TrimSpace(cc.Title)

	Description := strings.TrimSpace(cc.Description)

	if Title == ""{
		return nil, errors.New("Title cannot be empty")
	}

	if Description == ""{
		return nil, errors.New("Description cannot be empty")
	}

	if cc.Unit <= 0{
		return nil, errors.New("A course must have at least 1 unit")
	}

	return &learnhub_models.Course{
		ID: rand.Uint64(),
		Title: Title,
		Description: Description,
		Unit: cc.Unit,
	}, nil
}
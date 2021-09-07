package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-student-api/graph/generated"
	"github.com/go-student-api/graph/model"
	"github.com/go-student-api/models"
	uuid "github.com/satori/go.uuid"
)

func (r *mutationResolver) RegisterStudent(ctx context.Context, input model.RegisterStudentInput) (*string, error) {
	res := "Inserted successfully!"
	u1 := uuid.NewV4()

	student := models.Student{
		ID:         u1,
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		MiddleName: input.MiddleName,
		Gender:     input.Gender,
		Email:      input.Email,
		Avatar:     input.Avatar,
		GradeLevel: input.GradeLevel,
		Section:    input.Section,
	}

	err := r.DB.Create(&student).Error

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *mutationResolver) UpdateStudent(ctx context.Context, studentID string, input model.RegisterStudentInput) (*string, error) {
	res := "Updated successfully!"
	u1, err := uuid.FromString(studentID)
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return nil, err
	}

	update := models.Student{
		ID: u1,
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		MiddleName: input.MiddleName,
		Gender:     input.Gender,
		Email:      input.Email,
		Avatar:     input.Avatar,
		GradeLevel: input.GradeLevel,
		Section:    input.Section,
	}

	err = r.DB.Save(&update).Error

	if err != nil {
		return nil, errors.New("Error updating...")
	}

	return &res, nil
}

func (r *mutationResolver) DeleteStudent(ctx context.Context, studentID string) (*string, error) {
	res := "Deleted successfully!"
	err := r.DB.Where("student_id = ?", studentID).Delete(&models.Student{}).Error

	if err != nil {
		return nil, errors.New("Error deleting...")
	}

	return &res, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

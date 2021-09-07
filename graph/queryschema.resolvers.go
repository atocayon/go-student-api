package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-student-api/graph/generated"
	"github.com/go-student-api/graph/model"
)

func (r *queryResolver) Student(ctx context.Context, studentID string) (*model.Student, error) {
	student := model.Student{}

	err := r.DB.First(&student, "student_id = ?", studentID).Error

	if err != nil {
		return nil, errors.New("Student not found")
	}

	return &student, nil
}

func (r *queryResolver) Students(ctx context.Context) ([]*model.Student, error) {
	students := []*model.Student{}
	err := r.DB.Find(&students).Error

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Empty")
	}
	return students, nil
}

func (r *queryResolver) Section(ctx context.Context, section string) ([]*model.Student, error) {
	students := []*model.Student{}
	r.DB.Where("section = ?", section).Find(&students)

	return students, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

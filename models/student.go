package models

import (
	uuid "github.com/satori/go.uuid"
)


type Student struct {
  ID 			uuid.UUID 		`json:"id" gorm:"column:student_id;primary_key;"`
  FirstName 	string			`json:"firstName" gorm:"column:first_name"`
  MiddleName	string 			`json:"middleName" gorm:"column:middle_name"`
  LastName		string 			`json:"lastName" gorm:"column:last_name"`
  Gender		string 			`json:"gender" gorm:"column:gender"`
  Email			string 			`json:"email" gorm:"column:email"`
  Avatar		string 			`json:"avatar" gorm:"column:avatar"`
  GradeLevel	string 			`json:"gradeLevel" gorm:"column:grade_level"`
  Section		string 			`json:"section" gorm:"column:section"`
}

func (Student) IsEntity()  {}
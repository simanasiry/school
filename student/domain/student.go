package domain

import (
	"time"
)

type Student struct {
	ID        uint64     `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	Student   string     `json:"student"`
	SchoolId  uint64     `json:"school"`
}

func (Student) TableName() string {
	return "students"
}

type AddStudentRequest struct {
	UserName string `json:"userName"binding:"required"`
	Role     string `json:"role" binding:"required"`

	Student string `json:"student" binding:"required"`
}

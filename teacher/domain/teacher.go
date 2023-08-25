package domain

import (
	"time"
)

type Teacher struct {
	ID        uint64     `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	Teacher   string     `json:"teacher"`
	SchoolId  uint64     `json:"school"`
}

func (Teacher) TableName() string {
	return "teachers"
}

type AddTeacherRequest struct {
	UserRequested string `json:"userRequested"binding:"required"`
	Teacher       string `json:"teacher" binding:"required"`
	SchoolId      uint64 `json:"schoolId" binding:"required"`
}

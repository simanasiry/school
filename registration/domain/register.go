package domain

import (
	"time"
)

type Register struct {
	ID        uint64     `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	TeacherId uint64     `json:"teacherId"`
	StudentId uint64     `json:"studentId"`
}

func (Register) TableName() string {
	return "registrations"
}

type AddRegisterRequest struct {
	UserName  string `json:"userName" binding:"required"`
	Role      string `json:"role" binding:"required"`
	TeacherId uint64 `json:"teacherId" binding:"required"`
	StudentId uint64 `json:"studentId" binding:"required"`
}

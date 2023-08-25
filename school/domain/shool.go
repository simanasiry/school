package domain

import (
	"time"
)

type School struct {
	ID         uint64     `json:"id"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
	HeadMaster string     `json:"headMaster"`
	School     string     `json:"school"`
}

func (School) TableName() string {
	return "schools"
}

type AddSchoolRequest struct {
	Role     string `json:"role" binding:"required"`
	UserName string `json:"userName" binding:"required"`

	HeadMaster string `json:"headMaster" binding:"required"`
	School     string `json:"school" binding:"required"`
}

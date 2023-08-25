package repository

import (
	"github.com/simanasiry/school/student/domain"
	"gorm.io/gorm"
)

type repository struct {
	*gorm.DB
}

func NewRepository(db *gorm.DB) domain.Reposiory {
	p := new(repository)
	p.DB = db
	return p
}

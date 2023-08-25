package domain

import (
	std "github.com/simanasiry/school/student/domain"
	"gorm.io/gorm"
)

type Reposiory interface {
	Create(interface{}) *gorm.DB
	Model(interface{}) *gorm.DB
	First(interface{}, ...interface{}) *gorm.DB
	GetStudents(teacherId uint64) (*[]std.Student, error)
}

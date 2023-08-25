package usecase

import (
	"github.com/simanasiry/school/registration/domain"
	"github.com/simanasiry/school/school"
	"github.com/simanasiry/school/student"
	"github.com/simanasiry/school/teacher"
)

type usecase struct {
	repo          domain.Reposiory
	schoolModule  *school.Module
	teacherModule *teacher.Module
	studentModule *student.Module
}

func NewUc(repo domain.Reposiory, schoolModule *school.Module,
	studentModule *student.Module, teacherModule *teacher.Module) domain.Usecase {

	uc := &usecase{
		repo:          repo,
		schoolModule:  schoolModule,
		studentModule: studentModule,
		teacherModule: teacherModule,
	}
	return uc
}

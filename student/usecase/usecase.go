package usecase

import (
	"github.com/simanasiry/school/school"
	"github.com/simanasiry/school/student/domain"
)

type usecase struct {
	repo         domain.Reposiory
	schoolModule *school.Module
}

func NewUc(repo domain.Reposiory, schoolModule *school.Module) domain.Usecase {

	uc := &usecase{
		repo:         repo,
		schoolModule: schoolModule,
	}
	return uc
}
